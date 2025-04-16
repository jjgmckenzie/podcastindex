package podcastindex

import (
	"context"
	"net/url"

	"github.com/jjgmckenzie/podcastindex/podcast"
)

func (c *Client) GetPodcastByGUID(ctx context.Context, guid podcast.GUID) (*Podcast, error) {
	var response getPodcastResponse
	params := url.Values{"guid": {string(guid)}}
	err := c.api.Get(ctx, "/podcasts/byguid", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Feed, nil
}
