package internal

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// PodcastIndexAPI is the implementation of the api interface used by podcastindex.Client
// it is used to make requests to the PodcastIndex API
type PodcastIndexAPI struct {
	APIKey     string
	APISecret  string
	UserAgent  string
	BaseURL    *url.URL
	HTTPClient *http.Client
}

// doRequest performs the common logic for making a GET request.
// It constructs the URL, calls getRequest, and returns the response.
// The caller is responsible for closing the response body.
func (api *PodcastIndexAPI) doRequest(ctx context.Context, endpoint string, params url.Values) (*http.Response, error) {
	if api.HTTPClient == nil {
		return nil, fmt.Errorf("HTTPClient is nil, please set a valid HTTPClient")
	}
	requestURL := api.BaseURL.JoinPath(endpoint)
	requestURL.RawQuery = params.Encode()

	resp, err := api.getRequest(ctx, *requestURL, http.MethodGet, nil)
	if err != nil {
		// Wrap the error for better context
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}
	return resp, nil
}

// Get makes a GET request to the PodcastIndex API
//
// used internally by the PodcastIndex API Client
//
// IMPORTANT: Remember to close the response body.
//
// context: The context to use for the request
// endpoint: The endpoint to make the request to
// params: The query parameters to include in the request
// result: The struct to unmarshal the response into
//
// Returns: error if the request fails, or the response body is not valid JSON
func (api *PodcastIndexAPI) Get(ctx context.Context, endpoint string, params url.Values, result any) error {
	resp, err := api.doRequest(ctx, endpoint, params)
	if err != nil {
		// Error from doRequest already includes URL and context
		return err
	}
	defer func(body io.ReadCloser) {
		// ignore errors closing the body; we do not care about them once we have read the response body.
		_ = body.Close()
	}(resp.Body)

	statusIsOK := resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusFound

	if !statusIsOK {
		requestURLStr := resp.Request.URL.String() // Get URL from the actual request in the response
		if resp.StatusCode == http.StatusUnauthorized {
			return fmt.Errorf("authentication error when making request to podcast index API (%s), please verify your API key and API secret values are correct", requestURLStr)
		}
		if resp.StatusCode == http.StatusBadRequest {
			return fmt.Errorf("podcast index API at %s returned status code %d. This usually indicates a malformed request, potentially a bug in this client or an API change. Please file an issue at https://github.com/jjgmckenzie/podcast-index/issues with steps to reproduce the error",
				requestURLStr, resp.StatusCode)
		}
		bodyBytes, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			return fmt.Errorf("podcast index API at %s returned status code %d, failed to read error response: %w",
				requestURLStr, resp.StatusCode, readErr)
		}
		respBody := string(bodyBytes)
		return fmt.Errorf("podcast index API at %s returned status code %d with response: %s",
			requestURLStr, resp.StatusCode, respBody)
	}
	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response from podcast index API (%s): %w", resp.Request.URL.String(), err)
	}
	return nil
}

// getRequest creates a new HTTP request with the correct headers for the PodcastIndex API
//
// used internally by the PodcastIndex API Client
//
// context: The context to use for the request
// url: The URL to make the request to
// method: The HTTP method to use for the request
// body: The body to include in the request
//
// Returns: The new HTTP request, or an error if the request fails to create
func (api *PodcastIndexAPI) getRequest(ctx context.Context, url url.URL, method string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url.String(), body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	api.addRequiredHeaders(req)
	return api.HTTPClient.Do(req)
}

func (api *PodcastIndexAPI) addRequiredHeaders(req *http.Request) {
	unixTime := time.Now().Unix()
	req.Header.Set("User-Agent", api.UserAgent)
	req.Header.Set("X-Auth-Key", api.APIKey)
	req.Header.Set("X-Auth-Date", strconv.FormatInt(unixTime, 10))
	// A SHA-1 hash of the X-Auth-Key, the corresponding secret and the X-Auth-Date value concatenated as a string.
	// The resulting hash should be encoded as a hexadecimal value, two digits per byte, using lower case letters
	// for the hex digits "a" through "f".
	hash := sha1.New()
	hash.Write([]byte(api.APIKey + api.APISecret + strconv.FormatInt(unixTime, 10)))
	authHeader := hash.Sum(nil)
	req.Header.Set("Authorization", fmt.Sprintf("%x", authHeader))
}
