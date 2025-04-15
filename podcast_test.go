package podcastindex

import (
	"strings"
	"testing"
)

func TestPodcast_UnmarshalJSON_Errors(t *testing.T) {
	testCases := []struct {
		name          string
		inputJSON     string
		expectedError string // Check if the error message contains this string
	}{
		{
			name:          "Invalid Base JSON",
			inputJSON:     `{invalid json`,
			expectedError: "invalid character 'i' looking for beginning of object key string",
		},
		{
			name:          "Invalid URL",
			inputJSON:     `{"id": 1, "url": ":invalid"}`,
			expectedError: "failed to parse URL ':invalid'",
		},
		{
			name:          "Invalid OriginalURL",
			inputJSON:     `{"id": 1, "url": "http://example.com", "originalURL": ":invalid"}`,
			expectedError: "failed to parse OriginalURL ':invalid'",
		},
		{
			name:          "Invalid Link",
			inputJSON:     `{"id": 1, "url": "http://example.com", "originalURL": "http://example.com", "link": ":invalid"}`,
			expectedError: "failed to parse Link ':invalid'",
		},
		{
			name:          "Invalid Image URL",
			inputJSON:     `{"id": 1, "url": "http://example.com", "originalURL": "http://example.com", "link": "http://example.com", "image": ":invalid"}`,
			expectedError: "failed to parse Image URL ':invalid'",
		},
		{
			name:          "Invalid Artwork URL",
			inputJSON:     `{"id": 1, "url": "http://example.com", "originalURL": "http://example.com", "link": "http://example.com", "image": "http://example.com/image.png", "artwork": ":invalid"}`,
			expectedError: "failed to parse Artwork URL ':invalid'",
		},
		{
			name:          "Invalid Category ID",
			inputJSON:     `{"id": 1, "url": "http://example.com", "originalURL": "http://example.com", "link": "http://example.com", "image": "http://example.com/image.png", "artwork": "http://example.com/artwork.png", "categories": {"not_a_number": "Category Name"}}`,
			expectedError: "failed to convert category ID 'not_a_number' to int",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var p Podcast
			err := p.UnmarshalJSON([]byte(tc.inputJSON))

			if err == nil {
				t.Errorf("Expected an error, but got nil")
				return
			}

			if !strings.Contains(err.Error(), tc.expectedError) {
				t.Errorf("Expected error containing '%s', but got: %v", tc.expectedError, err)
			}
		})
	}
}
