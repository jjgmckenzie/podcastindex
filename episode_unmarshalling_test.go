package podcastindex

import (
	"encoding/json"
	"os"
	"testing"
)

// TestEpisodeMarshalUnmarshalJSON verifies that the MarshalJSON and UnmarshalJSON methods
// work correctly and maintain symmetry (marshal -> unmarshal -> marshal produces the same result).
func TestEpisodeMarshalUnmarshalJSON(t *testing.T) {
	// Read the mock JSON from file
	originalJSON, err := os.ReadFile("testdata/episodes_by_feed_id.json")
	if err != nil {
		t.Fatalf("Failed to read mock JSON file: %v", err)
	}

	var response getEpisodeResponse
	// Step 1: Unmarshal the original JSON to a response containing Episodes
	err = json.Unmarshal(originalJSON, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal original JSON: %v", err)
	}

	// Ensure we have at least one episode to test with
	if len(response.Items) == 0 {
		t.Fatalf("No episodes found in test data")
	}

	// Extract all original items from JSON for comparison
	var originalObj map[string]interface{}
	if err := json.Unmarshal(originalJSON, &originalObj); err != nil {
		t.Fatalf("Failed to parse original JSON: %v", err)
	}

	originalItems, ok := originalObj["items"].([]interface{})
	if !ok || len(originalItems) == 0 {
		t.Fatalf("Original JSON doesn't contain expected 'items' array")
	}

	// Test each episode individually
	for i, originalEpisode := range response.Items {
		t.Logf("Testing episode #%d: %s", i, originalEpisode.Title)

		// Step 2: Marshal the Episode struct back to JSON
		remarshaledJSON, err := json.Marshal(&originalEpisode)
		if err != nil {
			t.Fatalf("Failed to marshal episode #%d: %v", i, err)
		}

		// Extract the corresponding episode from the original JSON
		originalEpisodeJSON, err := json.Marshal(originalItems[i])
		if err != nil {
			t.Fatalf("Failed to extract original episode JSON for episode #%d: %v", i, err)
		}

		// Unmarshal both JSONs to maps for a fair comparison
		var originalMap, remarshaledMap map[string]interface{}
		if err := json.Unmarshal(originalEpisodeJSON, &originalMap); err != nil {
			t.Fatalf("Failed to unmarshal original episode JSON to map for episode #%d: %v", i, err)
		}
		if err := json.Unmarshal(remarshaledJSON, &remarshaledMap); err != nil {
			t.Fatalf("Failed to unmarshal remarshaled JSON to map for episode #%d: %v", i, err)
		}

		// Step 3: Compare the original and remarshaled JSON
		if !jsonMapsEqual(originalMap, remarshaledMap) {
			t.Errorf("Original JSON and remarshaled JSON are not equal for episode #%d", i)
			// Pretty print the JSONs for easier comparison in the error output
			originalPretty, _ := json.MarshalIndent(originalMap, "", "  ")
			remarshaledPretty, _ := json.MarshalIndent(remarshaledMap, "", "  ")
			t.Errorf("Original JSON: %s", originalPretty)
			t.Errorf("Remarshaled JSON: %s", remarshaledPretty)
		}
	}
}
