package podcastindex

import (
	"context"
	"net/url"
	"testing"
)

func TestGetPodcastByURLIntegration(t *testing.T) {
	// https://feeds.theincomparable.com/batmanuniversity
	var testValidFeedURL = url.URL{
		Scheme: "https",
		Host:   "feeds.theincomparable.com",
		Path:   "/batmanuniversity",
	}
	var testInvalidFeedURL = url.URL{
		Scheme: "https",
		Host:   "localhost",
		Path:   "/not-a-valid-path",
	}
	t.Run("Integration test: Client should be able to get podcast by feed URL", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByURL(context.Background(), testValidFeedURL)
		if err != nil {
			t.Fatalf("failed to get podcast: %v", err)
		}
		if p == nil {
			t.Fatalf("podcast is nil")
		}

	})
	t.Run("Integration test: Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.GetPodcastByURL(context.Background(), testValidFeedURL)
		if err == nil {
			t.Fatalf("expected error to be returned")
		}
	})
	t.Run("Integration test: Invalid feed URL should return an error", func(t *testing.T) {
		client := authenticatedClient(t)
		p, err := client.GetPodcastByURL(context.Background(), testInvalidFeedURL)
		if err == nil {
			t.Logf("error: %s", p.Title)
			t.Fatalf("expected error to be returned")
		}
		t.Logf("error: %v", err)
	})
}
