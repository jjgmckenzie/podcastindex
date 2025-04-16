package podcastindex

import (
	"context"

	"github.com/jjgmckenzie/podcastindex/podcast"
)

// categoriesResponse is the response from the categories/list endpoint on a 200/OK response.
//
// https://podcastindex-org.github.io/docs-api/#tag--Categories
type categoriesResponse struct {
	// Status is a boolean value that indicates if the request was successful.
	// This can be either "true" or "false", but is passed through as a string.
	Status string `json:"status"`
	// Result is a list of categories - this is the only field we care about.
	Result []podcast.Category `json:"feeds"`
	// Count is the number of items returned in the result.
	Count int `json:"count"`
	// Description is the description of the response.
	Description string `json:"description"`
}

// Categories returns all the possible categories supported by the index.
//
// Returns: A list of categories in []podcast.Category format
func (c *Client) Categories(ctx context.Context) ([]podcast.Category, error) {
	var response categoriesResponse
	err := c.api.Get(ctx, "/categories/list", nil, &response)
	if err != nil {
		return nil, err
	}
	return response.Result, nil
}
