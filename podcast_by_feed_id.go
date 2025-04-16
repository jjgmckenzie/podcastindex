package podcastindex

import (
	"context"
	"net/url"

	"github.com/jjgmckenzie/podcastindex/podcast"
	"strconv"
)

func (c *Client) GetPodcastByFeedID(ctx context.Context, feedID podcast.ID) (*Podcast, error) {
	var response getPodcastResponse
	params := url.Values{"id": {strconv.Itoa(int(feedID))}}
	err := c.api.Get(ctx, "/podcasts/byfeedid", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Feed, nil
}
