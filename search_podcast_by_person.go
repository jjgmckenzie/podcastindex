package podcastindex

import (
	"context"
	"net/url"
	"strconv"
)

// SearchPodcastsByPersonParams is a struct that contains the optional parameters for the SearchPodcastsByPerson method.
//
// The optional parameters are:
type SearchPodcastsByPersonParams struct {
	// Max is the maximum number of podcasts to return; default 10, maximum 1000
	Max int
	// FullText : If present, return the full text value of any text fields (ex: description). If not provided, field value is truncated to 100 words.
	FullText bool
}

// SearchPodcastsByPerson returns all feeds that match the search terms in the `title`, `author` or `owner` fields of the feed.
//
// Also accepts optional parameters to filter the results, see SearchPodcastsByPersonParams for more details.
func (c *Client) SearchPodcastsByPerson(ctx context.Context, person string, params *SearchPodcastsByPersonParams) ([]*Podcast, error) {
	var response searchResponse
	urlParams := url.Values{"q": {person}, "max": {"10"}}
	if params != nil {
		if params.Max != 0 {
			urlParams.Set("max", strconv.Itoa(params.Max))
		}
		if params.FullText {
			urlParams.Add("fulltext", "")
		}
	}
	err := c.api.Get(ctx, "/search/byperson", urlParams, &response)
	if err != nil {
		return nil, err
	}
	return response.Feeds, nil
}
