package podcastindex

import (
	"context"
	"podcastindex/podcast"
	"testing"
)

const validFeedID = 75075
const testValidFeedID = podcast.ID(validFeedID)
const testInvalidFeedID = podcast.ID(0)

func TestGetPodcastByFeedIDIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get podcast by feed ID", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByFeedID(context.Background(), testValidFeedID)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}

	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetPodcastByFeedID(context.Background(), testValidFeedID)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
	t.Run("Integration test: Invalid feed ID should return an error", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByFeedID(context.Background(), testInvalidFeedID)
		if err == nil {
			t.Logf("error: %s", p.Title)
			t.Fatalf("expected error to be returned")
		}
		t.Logf("error: %v", err)
	})
}
