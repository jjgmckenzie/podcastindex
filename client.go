package podcastindex

import (
	"context"
	"net/http"
	"net/url"

	"github.com/jjgmckenzie/podcastindex/internal"
)

// DefaultBaseURL is the base URL for the PodcastIndex API
// This is the URL that will be used if no baseURL is provided to the NewClient function
var DefaultBaseURL = url.URL{
	Scheme: "https",
	Host:   "api.podcastindex.org",
	Path:   "/api/1.0/",
}

// api is the interface for the API, so we can mock the API for testing.
type api interface {
	// Get makes a GET request to the PodcastIndex API
	//
	// Returns: error if the request fails, or the response body is not valid JSON
	Get(ctx context.Context, path string, params url.Values, result any) error
	// GetRawJSON makes a GET request to the PodcastIndex API and returns the raw JSON response;
	//
	// This is useful for debugging and testing
	//
	// Returns: the raw JSON response, or an error if the request fails
}

// Client is the client for the PodcastIndex Library
type Client struct {
	api api
}

// NewClientOptions is the options for the NewClient function
//
// UserAgent: Please identify the system/product you are using to make this request.
// Example: SuperPodcastPlayer/1.3
//
// APIKey: Your API key string
// Example: UXKCGDSYGUUEVQJSYDZH
//
// APISecret: Your API secret string
// Example: yzJe2eE7XV-3eY576dyRZ6wXyAbndh6LUrCZ8KN|
//
// BaseURL (Optional) is the base URL to use for the request; defaults to https://api.podcastindex.org/api/1.0/
// Useful if you are using a proxy or a local server for testing.
//
// HTTPClient (Optional) is the HTTP client to use for all api requests; defaults to http.DefaultClient
// Useful for custom timeouts, proxies, or TLS configuration
type NewClientOptions struct {
	// UserAgent: Please identify the system/product you are using to make this request.
	// Example: SuperPodcastPlayer/1.3
	UserAgent string
	// APIKey: Your API key string
	// Example: UXKCGDSYGUUEVQJSYDZH
	APIKey string
	// APISecret is the API secret to use for the request, eg yzJe2eE7XV-3eY576dyRZ6wXyAbndh6LUrCZ8KN|
	APISecret string
	// BaseURL (Optional) is the base URL to use for the request; defaults to https://api.podcastindex.org/api/1.0/
	BaseURL *url.URL
	// HTTPClient (Optional) is the HTTP client to use for all api requests; defaults to http.DefaultClient
	HTTPClient *http.Client
}

// NewClient creates a new Client, and takes a NewClientOptions struct as an argument
//
// Returns: A new PodcastIndex API Client
func NewClient(options NewClientOptions) *Client {
	if options.BaseURL == nil {
		options.BaseURL = &DefaultBaseURL
	}
	if options.HTTPClient == nil {
		options.HTTPClient = http.DefaultClient
	}
	return &Client{
		api: &internal.PodcastIndexAPI{
			BaseURL:    options.BaseURL,
			APIKey:     options.APIKey,
			APISecret:  options.APISecret,
			UserAgent:  options.UserAgent,
			HTTPClient: options.HTTPClient,
		},
	}
}
