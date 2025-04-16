package podcastindex

import (
	"context"
	"testing"
)

func TestGetEpisodesByFeedIDIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get podcast by feed ID", func(t *testing.T) {
		client := authenticatedClient(t)
		params := GetEpisodesByFeedIDParams{
			Max: 999,
		}
		p, err := client.GetEpisodesByFeedID(context.Background(), 743229, &params)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}
	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetEpisodesByFeedID(context.Background(), testValidFeedID, nil)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
	t.Run("Integration test: Invalid feed ID should return an error", func(t *testing.T) {
		client := authenticatedClient(t)
		_, err := client.GetEpisodesByFeedID(context.Background(), testInvalidFeedID, nil)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
		t.Logf("error: %v", err)
	})
}
