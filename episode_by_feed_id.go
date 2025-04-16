package podcastindex

import (
	"context"
	"net/url"

	"github.com/jjgmckenzie/podcastindex/episode"

	"strconv"
)

func (c *Client) GetEpisodeByID(ctx context.Context, feedID episode.ID) (*Podcast, error) {
	var response getPodcastResponse
	params := url.Values{"id": {strconv.Itoa(int(feedID))}}
	err := c.api.Get(ctx, "/episodes/byid", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Feed, nil
}
