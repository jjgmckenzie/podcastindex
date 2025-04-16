package podcastindex

import (
	"context"
	"net/url"
	"strconv"
)

type LiveEpisodesParams struct {
	// Max is the maximum number of episodes to return
	Max int
}

// Get all episodes that have been found in the podcast:liveitem from the feeds.
func (c *Client) GetLiveEpisodes(ctx context.Context, params *LiveEpisodesParams) (*[]Episode, error) {
	var response getEpisodeResponse
	urlParams := url.Values{}
	if params != nil {
		if params.Max != 0 {
			urlParams.Set("max", strconv.Itoa(params.Max))
		}
	}
	err := c.api.Get(ctx, "/episodes/live", urlParams, &response)
	if err != nil {
		return nil, err
	}
	return &response.Items, nil
}
