package podcastindex

import (
	"context"
	"net/url"
	"podcastindex/podcast"
	"strconv"
)

type podcastByFeedIdResponse struct {
	// Status indicates API request status; either "true" or "false"
	Status string `json:"status"`
	// Query is the object containing the input query data
	Query struct{} `json:"query"`
	// Feed is the known details of podcast feed; type Podcast - this is the response.
	Feed Podcast `json:"feed"`
	//  Description is the description of the response
	Description string `json:"description"`
}

func (c *Client) GetPodcastByFeedID(ctx context.Context, feedID podcast.ID) (*Podcast, error) {
	var response podcastByFeedIdResponse
	params := url.Values{"id": {strconv.Itoa(int(feedID))}}
	err := c.api.Get(ctx, "/podcasts/byfeedid", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Feed, nil
}
