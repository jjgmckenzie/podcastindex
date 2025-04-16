package podcastindex

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestEpisode_UnmarshalJSON_Errors(t *testing.T) {
	testCases := []struct {
		name        string
		jsonData    string
		expectedErr string
	}{
		{
			name:        "Invalid Base JSON",
			jsonData:    `{invalid json`,
			expectedErr: "invalid character 'i' looking for beginning of object key string",
		},
		{
			name:        "Invalid Link URL",
			jsonData:    `{"id":1,"link":"invalid url ://"}`,
			expectedErr: "failed to parse Link 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
		{
			name:        "Invalid Enclosure URL",
			jsonData:    `{"id":1,"link":"http://example.com","enclosureUrl":"invalid url ://"}`,
			expectedErr: "failed to parse EnclosureURL 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
		{
			name:        "Invalid Image URL",
			jsonData:    `{"id":1,"link":"http://example.com","enclosureUrl":"http://example.com/enclosure.mp3","image":"invalid url ://"}`,
			expectedErr: "failed to parse Image URL 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
		{
			name:        "Invalid Feed Image URL",
			jsonData:    `{"id":1,"link":"http://example.com","enclosureUrl":"http://example.com/enclosure.mp3","image":"http://example.com/image.jpg","feedImage":"invalid url ://"}`,
			expectedErr: "failed to parse FeedImage URL 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
		{
			name:        "Invalid Feed URL",
			jsonData:    `{"id":1,"link":"http://example.com","enclosureUrl":"http://example.com/enclosure.mp3","image":"http://example.com/image.jpg","feedImage":"http://example.com/feedimage.jpg","feedUrl":"invalid url ://"}`,
			expectedErr: "failed to parse FeedURL 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
		{
			name:        "Invalid Chapters URL",
			jsonData:    `{"id":1,"link":"http://example.com","enclosureUrl":"http://example.com/enclosure.mp3","image":"http://example.com/image.jpg","feedImage":"http://example.com/feedimage.jpg","feedUrl":"http://example.com/feed.xml","chaptersUrl":"invalid url ://"}`,
			expectedErr: "failed to parse ChaptersURL 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
		{
			name:        "Invalid Transcript URL",
			jsonData:    `{"id":1,"link":"http://example.com","enclosureUrl":"http://example.com/enclosure.mp3","image":"http://example.com/image.jpg","feedImage":"http://example.com/feedimage.jpg","feedUrl":"http://example.com/feed.xml","chaptersUrl":"http://example.com/chapters.json","transcriptUrl":"invalid url ://"}`,
			expectedErr: "failed to parse TranscriptURL 'invalid url ://': parse \"invalid url ://\": first path segment in URL cannot contain colon",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var e Episode
			err := json.Unmarshal([]byte(tc.jsonData), &e)

			if err == nil {
				t.Errorf("expected an error but got nil")
				return // No further checks if no error
			}

			// Check if the error message contains the expected substring for more specific checks
			if !strings.Contains(err.Error(), tc.expectedErr) {
				t.Errorf("expected error message '%s' to contain '%s'", err.Error(), tc.expectedErr)
			}
		})
	}
}
