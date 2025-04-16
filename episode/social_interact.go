package episode

import "net/url"

type SocialInteract struct {
	// URL is  The uri/url of the root post comment
	URL url.URL `json:"url"`
	// Protocol is the protocol in use for interacting with the comment root post.
	//
	// For the most up-to-date list of options, see https://github.com/Podcastindex-org/podcast-namespace/blob/main/socialprotocols.txt
	Protocol string `json:"protocol"`
	// AccountID is the account id (on the commenting platform) of the account that created this root post.
	AccountID string `json:"accountId"`
	// AccountURL is the public url (on the commenting platform) of the account that created this root post.
	AccountURL url.URL `json:"accountUrl"`
	// Priority : When multiple socialInteract tags are present, this integer gives order of priority. A lower number means higher priority.
	Priority int `json:"priority"`
}
