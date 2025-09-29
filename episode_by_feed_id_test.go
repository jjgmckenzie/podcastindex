package podcastindex

import (
	"context"
	"strings"
	"testing"

	"github.com/jjgmckenzie/podcastindex/episode"
)

const testEpisodeID = episode.ID(43141716087)
const testInvalidEpisodeID = episode.ID(0)

func TestGetEpisodeByFeedIDIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get episode by feed ID", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetEpisodeByID(context.Background(), testEpisodeID)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}
	})
	t.Run("Integration test: Server response should be correct", func(t *testing.T) {
		client := authenticatedClient(t)
		e, err := client.GetEpisodeByID(context.Background(), testEpisodeID)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if e.ID != testEpisodeID {
			t.Fatalf("got episode ID %d, want %d", e.ID, testEpisodeID)
		}
	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetEpisodeByID(context.Background(), testEpisodeID)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
	t.Run("Integration test: Invalid episode ID should return an error", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetEpisodeByID(context.Background(), testInvalidEpisodeID)
		if err == nil {
			t.Logf("error: %s", p.Title)
			t.Fatalf("expected error to be returned")
		}
		t.Logf("error: %v", err)
	})
	t.Run("Integration test: Episode description should not end with '...' and should be more than 100 words", func(t *testing.T) {
		client := authenticatedClient(t)
		e, err := client.GetEpisodeByID(context.Background(), testEpisodeID)
		if err != nil {
			t.Fatalf("failed to get episode: %v", err)
		}
		desc := e.Description
		if len(desc) == 0 {
			t.Fatalf("description is empty")
		}
		if len(desc) >= 3 && desc[len(desc)-3:] == "..." {
			t.Fatalf("description ends with '...'")
		}
		wordCount := len(strings.Fields(desc))
		if wordCount <= 100 {
			t.Fatalf("description has %d words, want > 100", wordCount)
		}
	})
}
