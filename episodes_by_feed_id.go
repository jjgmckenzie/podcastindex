package podcastindex

import (
	"context"
	"net/url"
	"podcastindex/podcast"
	"strconv"
)

func (c *Client) GetEpisodesByFeedID(ctx context.Context, feedID podcast.ID, params *GetEpisodesParams) (*[]Episode, error) {
	var response getEpisodeResponse
	urlParams := url.Values{"id": {strconv.Itoa(int(feedID))}}
	if params != nil {
		if params.Max != 0 {
			urlParams.Set("max", strconv.Itoa(params.Max))
		}
		if params.FullText {
			urlParams.Add("fulltext", "")
		}
	}
	err := c.api.Get(ctx, "/episodes/byfeedid", urlParams, &response)
	if err != nil {
		return nil, err
	}
	return &response.Items, nil
}
