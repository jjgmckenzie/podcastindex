package podcastindex

import (
	"context"
	"testing"
)

func TestGetLiveEpisodes(t *testing.T) {
	t.Run("Integration test: Client should be able to get live episodes", func(t *testing.T) {
		client := authenticatedClient(t)
		episodes, err := client.GetLiveEpisodes(context.Background(), &LiveEpisodesParams{
			Max: 999,
		})
		if err != nil {
			t.Fatalf("Failed to get live episodes: %v", err)
		}
		if len(*episodes) == 0 {
			t.Fatalf("No live episodes found")
		}
		for _, episode := range *episodes {
			t.Logf("Episode: %v", episode)
		}
	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetLiveEpisodes(context.Background(), nil)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
}
