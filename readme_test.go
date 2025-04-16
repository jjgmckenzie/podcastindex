package podcastindex

import (
	"context"
	"testing"
)

func TestReadme(t *testing.T) {
	client := authenticatedClient(t)
	ctx := context.Background()
	podcasts, err := client.SearchPodcastsByTitle(ctx, "test", nil)
	if err != nil {
		t.Fatalf("failed to search podcasts: %v", err)
	}
	for _, podcast := range podcasts {
		t.Logf("podcast: %s", podcast.Title)
		episodes, err := client.GetEpisodes(ctx, *podcast, nil)
		if err != nil {
			t.Fatalf("failed to get episodes: %v", err)
		}
		for i, episode := range *episodes {
			t.Logf("%d %s : %s\n", i, podcast.Title, episode.Title)
		}
	}
}
