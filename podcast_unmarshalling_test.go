package podcastindex

import (
	"encoding/json"
	"os"
	"sort"
	"strings"
	"testing"
)

// TestPodcastMarshalUnmarshalJSON verifies that the MarshalJSON and UnmarshalJSON methods
// work correctly and maintain symmetry (marshal -> unmarshal -> marshal produces the same result).
func TestPodcastMarshalUnmarshalJSON(t *testing.T) {
	// Read the mock JSON from file
	originalJSON, err := os.ReadFile("testdata/example_podcasts_by_title.json")
	if err != nil {
		t.Fatalf("Failed to read mock JSON file: %v", err)
	}

	var response searchResponse
	// Step 1: Unmarshal the original JSON to a Podcast struct
	err = json.Unmarshal(originalJSON, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal original JSON: %v", err)
	}

	// Ensure we have at least one podcast to test with
	if len(response.Feeds) == 0 {
		t.Fatalf("No podcasts found in test data")
	}

	// Extract all original feeds from JSON for comparison
	var originalObj map[string]interface{}
	if err := json.Unmarshal(originalJSON, &originalObj); err != nil {
		t.Fatalf("Failed to parse original JSON: %v", err)
	}

	originalFeeds, ok := originalObj["feeds"].([]interface{})
	if !ok || len(originalFeeds) == 0 {
		t.Fatalf("Original JSON doesn't contain expected 'feeds' array")
	}

	// Test each podcast individually
	for i, originalPodcast := range response.Feeds {
		t.Logf("Testing podcast #%d: %s", i, originalPodcast.Title)

		// Step 2: Marshal the Podcast struct back to JSON
		remarshaledJSON, err := json.Marshal(originalPodcast)
		if err != nil {
			t.Fatalf("Failed to marshal podcast #%d: %v", i, err)
		}

		// Extract the corresponding podcast from the original JSON
		originalPodcastJSON, err := json.Marshal(originalFeeds[i])
		if err != nil {
			t.Fatalf("Failed to extract original podcast JSON for podcast #%d: %v", i, err)
		}

		// Unmarshal both JSONs to maps for a fair comparison
		var originalMap, remarshaledMap map[string]interface{}
		if err := json.Unmarshal(originalPodcastJSON, &originalMap); err != nil {
			t.Fatalf("Failed to unmarshal original podcast JSON to map for podcast #%d: %v", i, err)
		}
		if err := json.Unmarshal(remarshaledJSON, &remarshaledMap); err != nil {
			t.Fatalf("Failed to unmarshal remarshaled JSON to map for podcast #%d: %v", i, err)
		}

		// Sort categories in maps for deterministic comparison
		sortCategoriesInMap(originalMap)
		sortCategoriesInMap(remarshaledMap)

		// Step 3: Compare the original and remarshaled JSON
		if !jsonMapsEqual(originalMap, remarshaledMap) {
			t.Errorf("Original JSON and remarshaled JSON are not equal for podcast #%d", i)
			// Pretty print the JSONs for easier comparison in the error output
			originalPretty, _ := json.MarshalIndent(originalMap, "", "  ")
			remarshaledPretty, _ := json.MarshalIndent(remarshaledMap, "", "  ")
			t.Errorf("Original JSON: %s", originalPretty)
			t.Errorf("Remarshaled JSON: %s", remarshaledPretty)
		}
	}
}

func TestPodcastUnmarshalJSONWithExtraVariableFails(t *testing.T) {
	// Read the mock JSON from file
	originalJSON, err := os.ReadFile("testdata/example_podcast_with_extra_var.json")
	if err != nil {
		t.Fatalf("Failed to read mock JSON file: %v", err)
	}

	// Unmarshal the original JSON to a Podcast struct
	var podcast Podcast
	if err := json.Unmarshal(originalJSON, &podcast); err != nil {
		t.Fatalf("Failed to unmarshal original JSON: %v", err)
	}

	// Marshal the Podcast struct back to JSON
	remarshaledJSON, err := json.Marshal(podcast)
	if err != nil {
		t.Fatalf("Failed to marshal podcast: %v", err)
	}

	// Unmarshal both JSONs to maps for comparison
	var originalMap, remarshaledMap map[string]interface{}
	if err := json.Unmarshal(originalJSON, &originalMap); err != nil {
		t.Fatalf("Failed to unmarshal original JSON to map: %v", err)
	}
	if err := json.Unmarshal(remarshaledJSON, &remarshaledMap); err != nil {
		t.Fatalf("Failed to unmarshal remarshaled JSON to map: %v", err)
	}

	// Sort categories in maps for deterministic comparison
	sortCategoriesInMap(originalMap)
	sortCategoriesInMap(remarshaledMap)

	// Assert that the original and remarshaled JSON are NOT equal due to the extra variable
	if jsonMapsEqual(originalMap, remarshaledMap) {
		t.Errorf("Expected original JSON and remarshaled JSON to be different due to extra variable")
		// Pretty print the JSONs for easier comparison in the error output
		originalPretty, _ := json.MarshalIndent(originalMap, "", "  ")
		remarshaledPretty, _ := json.MarshalIndent(remarshaledMap, "", "  ")
		t.Errorf("Original JSON: %s", originalPretty)
		t.Errorf("Remarshaled JSON: %s", remarshaledPretty)
	}
}

// sortCategoriesInMap sorts the categories in a map representation of a podcast
func sortCategoriesInMap(podcastMap map[string]interface{}) {
	if categories, ok := podcastMap["categories"].([]interface{}); ok {
		sort.Slice(categories, func(i, j int) bool {
			catI, okI := categories[i].(map[string]interface{})
			catJ, okJ := categories[j].(map[string]interface{})
			if !okI || !okJ {
				return false
			}
			idI, okI := catI["id"].(float64)
			idJ, okJ := catJ["id"].(float64)
			if !okI || !okJ {
				return false
			}
			return idI < idJ
		})
	}
}

// jsonMapsEqual compares two maps representing JSON objects
func jsonMapsEqual(a, b map[string]interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k, va := range a {
		vb, ok := b[k]
		if !ok {
			return false
		}

		// Special case for language field - compare as lowercase
		// https://github.com/Podcastindex-org/docs-api/issues/142
		if k == "language" {
			if vaStr, okA := va.(string); okA {
				if vbStr, okB := vb.(string); okB {
					// If both are valid strings, compare them case-insensitively
					if strings.EqualFold(vaStr, vbStr) {
						continue // Values match case-insensitively, move to next field
					}
					return false
				}
			}
		}

		switch va := va.(type) {
		case map[string]interface{}:
			if vb, ok := vb.(map[string]interface{}); ok {
				if !jsonMapsEqual(va, vb) {
					return false
				}
			} else {
				return false
			}
		case []interface{}:
			if vb, ok := vb.([]interface{}); ok {
				if len(va) != len(vb) {
					return false
				}
				for i := range va {
					if !jsonValuesEqual(va[i], vb[i]) {
						return false
					}
				}
			} else {
				return false
			}
		default:
			if va != vb {
				return false
			}
		}
	}
	return true
}

// jsonValuesEqual compares two JSON values
func jsonValuesEqual(a, b interface{}) bool {
	switch va := a.(type) {
	case map[string]interface{}:
		if vb, ok := b.(map[string]interface{}); ok {
			return jsonMapsEqual(va, vb)
		}
		return false
	case []interface{}:
		if vb, ok := b.([]interface{}); ok {
			if len(va) != len(vb) {
				return false
			}
			for i := range va {
				if !jsonValuesEqual(va[i], vb[i]) {
					return false
				}
			}
			return true
		}
		return false
	default:
		return a == b
	}
}
