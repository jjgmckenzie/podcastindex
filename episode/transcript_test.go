package episode

import (
	"encoding/json"
	"net/url"
	"reflect"
	"testing"
)

func TestTranscript_MarshalUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		transcript   Transcript
		expectedJSON string
		expectErr    bool // Expect error during unmarshal
	}{
		{
			name: "Valid Transcript",
			transcript: Transcript{
				URL:  url.URL{Scheme: "https", Host: "example.com", Path: "/transcript.vtt"},
				Type: TranscriptVTT,
			},
			expectedJSON: `{"URL":"https://example.com/transcript.vtt","Type":"text/vtt"}`,
			expectErr:    false,
		},
		{
			name: "Plaintext Transcript",
			transcript: Transcript{
				URL:  url.URL{Scheme: "http", Host: "another.org", Path: "/plain.txt"},
				Type: TranscriptPlaintext,
			},
			expectedJSON: `{"URL":"http://another.org/plain.txt","Type":"text/plain"}`,
			expectErr:    false,
		},
		{
			name: "Empty URL", // Still technically valid URL, marshals fine
			transcript: Transcript{
				URL:  url.URL{},
				Type: TranscriptJSON,
			},
			expectedJSON: `{"URL":"","Type":"application/json"}`,
			expectErr:    false,
		},
		{
			name:         "Invalid URL string for unmarshal",
			transcript:   Transcript{}, // Used only for unmarshal target
			expectedJSON: `{"URL":"://invalid url","Type":"text/html"}`,
			expectErr:    true, // Expect error during unmarshal
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test Marshal
			if !tt.expectErr { // Only test marshalling for valid initial transcripts
				jsonData, err := json.Marshal(tt.transcript)
				if err != nil {
					t.Fatalf("MarshalJSON() error = %v", err)
				}
				if string(jsonData) != tt.expectedJSON {
					t.Errorf("MarshalJSON() got = %v, want %v", string(jsonData), tt.expectedJSON)
				}
			}

			// Test Unmarshal
			var unmarshalledTranscript Transcript
			err := json.Unmarshal([]byte(tt.expectedJSON), &unmarshalledTranscript)

			if tt.expectErr {
				if err == nil {
					t.Fatalf("UnmarshalJSON() expected an error, but got none")
				}
				// Can add more specific error checking here if needed
				return // Stop testing this case after expected error
			}

			if err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}

			// Re-parse the expected URL for comparison, as url.URL comparison needs care
			expectedURL, _ := url.Parse(tt.transcript.URL.String())

			if !reflect.DeepEqual(unmarshalledTranscript.URL, *expectedURL) || unmarshalledTranscript.Type != tt.transcript.Type {
				t.Errorf("UnmarshalJSON() got = %+v, want %+v", unmarshalledTranscript, tt.transcript)
			}
		})
	}
}
