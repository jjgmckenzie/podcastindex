package podcastindex

import (
	"context"
	"net/url"

	"github.com/jjgmckenzie/podcastindex/episode"

	"strconv"
)

type getSingleEpisodeResponse struct {
	Id          string  `json:"id"`
	Episode     Episode `json:"episode"`
	Description string  `json:"description"`
}

func (c *Client) GetEpisodeByID(ctx context.Context, feedID episode.ID) (*Episode, error) {
	var response getSingleEpisodeResponse
	params := url.Values{"id": {strconv.Itoa(int(feedID))}}
	err := c.api.Get(ctx, "/episodes/byid", params, &response)
	if err != nil {
		return nil, err
	}
	return &response.Episode, nil
}
