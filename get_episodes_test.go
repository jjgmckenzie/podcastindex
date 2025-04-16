package podcastindex

import (
	"context"
	"testing"
)

func TestGetEpisodesIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get podcast by feed ID", func(t *testing.T) {
		client := authenticatedClient(t)
		Podcast, err := client.GetPodcastByFeedID(context.Background(), 743229)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		params := GetEpisodesParams{
			Max: 10,
		}
		p, err := client.GetEpisodes(context.Background(), *Podcast, &params)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}
	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		authenticatedClient := authenticatedClient(t)
		Podcast, err := authenticatedClient.GetPodcastByFeedID(context.Background(), 743229)

		unAuthenticatedClient := unauthenticatedClient()
		_, err = unAuthenticatedClient.GetEpisodes(context.Background(), *Podcast, nil)
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
