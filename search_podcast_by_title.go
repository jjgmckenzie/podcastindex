package podcastindex

import (
	"context"
	"net/url"
	"strconv"
)

// SearchPodcastsByTitleParams is a struct that contains the optional parameters for the SearchPodcastsByTitle method.
//
// The optional parameters are:
type SearchPodcastsByTitleParams struct {
	// Max is the maximum number of podcasts to return; default 10, maximum 1000
	Max int
	// Clean : If present, only non-explicit feeds will be returned. Meaning, feeds where the itunes:explicit flag is set to false.
	Clean bool
	// FullText : If present, return the full text value of any text fields (ex: description). If not provided, field value is truncated to 100 words.
	FullText bool
	// Similar : If present, include similar matches in search response
	Similar bool
	// Value : Only returns feeds with a Value block of the specified type. Use any to return feeds with any value block.
	//
	// Valid values are: podcast.value.PaymentAny, podcast.value.PaymentLightning, podcast.value.PaymentHive, or podcast.value.PaymentWebMonetization.
	Value string
}

// SearchPodcastsByTitle returns a list of podcasts that match the given title.
//
// Also accepts optional parameters to filter the results, see SearchPodcastsByTitleParams for more details.
func (c *Client) SearchPodcastsByTitle(ctx context.Context, title string, params *SearchPodcastsByTitleParams) ([]*Podcast, error) {
	var response searchResponse
	urlParams := url.Values{"q": {title}, "max": {"10"}}
	if params != nil {
		if params.Max != 0 {
			urlParams.Set("max", strconv.Itoa(params.Max))
		}
		if params.Clean {
			urlParams.Add("clean", "")
		}
		if params.FullText {
			urlParams.Add("fulltext", "")
		}
		if params.Similar {
			urlParams.Add("similar", "")
		}
		if params.Value != "" {
			urlParams.Add("val", params.Value)
		}
	}
	err := c.api.Get(ctx, "/search/bytitle", urlParams, &response)
	if err != nil {
		return nil, err
	}
	return response.Feeds, nil
}
