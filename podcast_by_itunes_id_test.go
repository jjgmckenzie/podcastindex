package podcastindex

import (
	"context"
	"github.com/jjgmckenzie/podcastindex/podcast"
	"testing"
)

const validITunesID = "id1441923632"
const testValidITunesID = podcast.ITunesID(validITunesID)
const testInvalidITunesID = podcast.ITunesID("0")

func TestGetPodcastByITunesIDIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get podcast by iTunes ID", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByITunesID(context.Background(), testValidITunesID)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}

	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetPodcastByITunesID(context.Background(), testValidITunesID)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
	t.Run("Integration test: Invalid iTunes ID should return an error", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByITunesID(context.Background(), testInvalidITunesID)
		if err == nil {
			t.Logf("error: %s", p.Title)
			t.Fatalf("expected error to be returned")
		}
		t.Logf("error: %v", err)
	})
}
