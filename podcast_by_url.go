package podcastindex

import (
	"context"
	"net/url"
)

func (c *Client) GetPodcastByURL(ctx context.Context, feedURL url.URL) (*Podcast, error) {
	var response getPodcastResponse
	params := url.Values{"url": {feedURL.String()}}
	err := c.api.Get(ctx, "/podcasts/byfeedurl", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Feed, nil
}
