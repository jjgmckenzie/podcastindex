package podcastindex

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type SearchServer struct {
	Server *httptest.Server
	path   *string
	values *url.Values
}

// GetSearchServer returns a new httptest.Server and the requested path and query parameters.
func GetSearchServer(t *testing.T) *SearchServer {
	t.Helper()
	var requestedPath string
	var requestedQuery url.Values

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestedPath = r.URL.Path
		requestedQuery = r.URL.Query()

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("{}"))
		if err != nil {
			t.Fatalf("Failed to write mock JSON response: %v", err)
		}
	}))
	return &SearchServer{server, &requestedPath, &requestedQuery}
}

func (s *SearchServer) expectPath(path string) error {
	if *s.path != path {
		return fmt.Errorf("expected path %s, got %s", *s.path, path)
	}
	return nil
}

func (s *SearchServer) expectQuery(query url.Values) error {
	// Check if all expected keys are present and match (with boolean handling)
	for key, expectedVals := range query {
		actualVals, ok := (*s.values)[key]
		if !ok {
			return fmt.Errorf("expected query key %s not found in actual query %s", key, s.values.Encode())
		}

		// Handle boolean flag case: expected "true" matches actual "" or "true"
		if len(expectedVals) == 1 && expectedVals[0] == "true" {
			if len(actualVals) == 1 && (actualVals[0] == "" || actualVals[0] == "true") {
				continue // Match found
			} else {
				return fmt.Errorf("expected boolean query param %s=true, got %s=%v", key, key, actualVals)
			}
		}

		// Default comparison for non-boolean or multi-value params
		if !slicesEqual(expectedVals, actualVals) {
			return fmt.Errorf("expected query param %s=%v, got %s=%v", key, expectedVals, key, actualVals)
		}
	}

	// Check if there are any extra keys in the actual query
	for key := range *s.values {
		if _, ok := query[key]; !ok {
			return fmt.Errorf("unexpected query key %s found in actual query %s", key, s.values.Encode())
		}
	}

	return nil
}

// slicesEqual compares two string slices for equality, ignoring order.
// It treats nil, empty, and single-empty-string slices as equivalent.
func slicesEqual(a, b []string) bool {
	lenA, lenB := len(a), len(b)

	// Treat nil/empty slice and slice with one empty string as equivalent
	isAeffectivelyEmpty := lenA == 0 || (lenA == 1 && a[0] == "")
	isBeffectivelyEmpty := lenB == 0 || (lenB == 1 && b[0] == "")

	if isAeffectivelyEmpty && isBeffectivelyEmpty {
		return true
	}
	// If one is effectively empty and the other isn't, they are not equal
	if isAeffectivelyEmpty != isBeffectivelyEmpty {
		return false
	}

	// If we reach here, neither is effectively empty, proceed with comparison
	if lenA != lenB { // This check might be redundant now but harmless
		return false
	}

	// Create maps for efficient comparison, order doesn't matter for query params
	mapA := make(map[string]int)
	mapB := make(map[string]int)
	for _, v := range a {
		mapA[v]++
	}
	for _, v := range b {
		mapB[v]++
	}
	if len(mapA) != len(mapB) { // Different unique values
		return false
	}
	for k, v := range mapA {
		if mapB[k] != v {
			return false
		}
	}
	return true
}

func (s *SearchServer) ExpectPathAndQuery(path string, query url.Values) error {
	err := s.expectPath(path)
	if err != nil {
		return err
	}
	return s.expectQuery(query)
}

func GetErrorServer(t *testing.T) *Client {
	t.Helper()
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer errorServer.Close()

	errorServerURL, _ := url.Parse(errorServer.URL)
	return NewClient(NewClientOptions{
		BaseURL: errorServerURL,
	})
}
