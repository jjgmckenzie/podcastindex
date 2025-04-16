package podcastindex

import (
	"context"
	"net/url"

	"github.com/jjgmckenzie/podcastindex/podcast"
)

func (c *Client) GetPodcastByITunesID(ctx context.Context, itunesID podcast.ITunesID) (*Podcast, error) {
	var response getPodcastResponse
	params := url.Values{"id": {string(itunesID)}}
	err := c.api.Get(ctx, "/podcasts/byitunesid", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Feed, nil
}
