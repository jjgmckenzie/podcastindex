package podcast

import "net/url"

// Funding is the information for donation/funding the podcast. May not be reported.
//
// See the podcast namespace spec for more information.
//
// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#funding
type Funding struct {
	// URL is the URL of the funding page; may be nil.
	URL *url.URL `json:"url"`
	// Message is the description of the funding page.
	Message string `json:"message"`
}
