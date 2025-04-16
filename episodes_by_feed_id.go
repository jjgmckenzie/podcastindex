package podcastindex

import (
	"context"
	"net/url"
	"podcastindex/podcast"
	"strconv"
)

type GetEpisodesByFeedIDParams struct {
	// Max is the maximum number of episodes to return
	Max int
	//If present, return the full text value of any text fields (ex: description). If not provided, field value is truncated to 100 words.
	FullText bool
}

func (c *Client) GetEpisodesByFeedID(ctx context.Context, feedID podcast.ID, params *GetEpisodesByFeedIDParams) (*[]Episode, error) {
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
