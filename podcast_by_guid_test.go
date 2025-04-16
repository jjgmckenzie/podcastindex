package podcastindex

import (
	"context"
	"github.com/jjgmckenzie/podcastindex/podcast"
	"testing"
)

const validGUID = "9b024349-ccf0-5f69-a609-6b82873eab3c"
const testValidGUID = podcast.GUID(validGUID)
const testInvalidGUID = podcast.GUID("")

func TestGetPodcastByGUIDIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get podcast by GUID", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByGUID(context.Background(), testValidGUID)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}

	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetPodcastByGUID(context.Background(), testValidGUID)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
	t.Run("Integration test: Invalid GUID should return an error", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByGUID(context.Background(), testInvalidGUID)
		if err == nil {
			t.Logf("error: %s", p.Title)
			t.Fatalf("expected error to be returned")
		}
		t.Logf("error: %v", err)
	})
}
