package internal

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

// setupTestAPI creates a test server and a PodcastIndexAPI instance configured to use it.
func setupTestAPI(t *testing.T, handler http.HandlerFunc) (*PodcastIndexAPI, *httptest.Server) {
	t.Helper()

	server := httptest.NewServer(handler)
	serverURL, _ := url.Parse(server.URL)
	api := &PodcastIndexAPI{
		APIKey:     "testKey",
		APISecret:  "testSecret",
		UserAgent:  "testAgent",
		BaseURL:    serverURL,
		HTTPClient: server.Client(),
	}
	t.Cleanup(server.Close)
	return api, server
}

// newTestAPI creates a PodcastIndexAPI instance with default valid values for tests
// that don't require a live server.
func newTestAPI() *PodcastIndexAPI {
	placeHolderURL := url.URL{
		Scheme: "http",
		Host:   "localhost",
	}
	return &PodcastIndexAPI{
		APIKey:     "testKey",
		BaseURL:    &placeHolderURL,
		APISecret:  "testSecret",
		UserAgent:  "testAgent",
		HTTPClient: http.DefaultClient,
	}
}

func TestGetRequestURL(t *testing.T) {
	var capturedRequest *http.Request

	handler := func(w http.ResponseWriter, r *http.Request) {
		capturedRequest = r
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status": "true", "description": "Success"}`)) // Send a minimal valid JSON response
	}
	api, _ := setupTestAPI(t, handler)

	params := url.Values{}
	params.Add("query", "test search")
	params.Add("max", "10")
	params.Add("lang", "en")
	var result struct{} // Dummy result struct

	err := api.Get(context.Background(), "search/byterm", params, &result)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	t.Run("CorrectlyEncodesQueryParams", func(t *testing.T) {
		if capturedRequest == nil {
			t.Fatal("Request was not made to the test server")
		}
		expectedQuery := "lang=en&max=10&query=test+search" // url.Values.Encode sorts parameters alphabetically
		actualQuery := capturedRequest.URL.RawQuery
		if actualQuery != expectedQuery {
			t.Errorf("Expected query string %q, got %q", expectedQuery, actualQuery)
		}
	})

	t.Run("UsesCorrectPath", func(t *testing.T) {
		if capturedRequest == nil {
			t.Fatal("Request was not made to the test server")
		}
		expectedPath := "/search/byterm" // Base URL from server includes the host and port
		actualPath := capturedRequest.URL.Path
		if actualPath != expectedPath {
			t.Errorf("Expected path %q, got %q", expectedPath, actualPath)
		}
	})
}

// errorReader is a helper type that implements io.Reader and always returns an error.
type errorReader struct{}

func (er *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("simulated read error")
}

// TestGetResponseErrorHandling covers various non-200 status codes and invalid responses.
func TestGetResponseErrorHandling(t *testing.T) {
	testCases := []struct {
		name            string
		statusCode      int
		responseBody    []byte // Reverted type back to []byte
		expectedErrMsg  string
		setupClient     func(serverURL string) *PodcastIndexAPI // Optional setup override
		skipAuthHeaders bool                                    // If true, don't set auth headers
	}{
		{
			name:           "Unauthorized",
			statusCode:     http.StatusUnauthorized,
			expectedErrMsg: "authentication error",
		},
		{
			name:           "BadRequest",
			statusCode:     http.StatusBadRequest,
			expectedErrMsg: "returned status code 400",
		},
		{
			name:           "ServerError",
			statusCode:     http.StatusInternalServerError,
			responseBody:   []byte("internal server error"),
			expectedErrMsg: "returned status code 500 with response: internal server error",
		},
		{
			name:           "InvalidJSONResponse",
			statusCode:     http.StatusOK,
			responseBody:   []byte(`this is not json`),
			expectedErrMsg: "failed to decode response",
			// Need custom setup to avoid auth headers being checked for a 200 response
			setupClient: func(serverURL string) *PodcastIndexAPI {
				url, _ := url.Parse(serverURL)
				return &PodcastIndexAPI{
					BaseURL: url,
					// No APIKey/Secret needed here as we expect a decode error after 200 OK
					UserAgent:  "testAgent",
					HTTPClient: nil, // Will be set later
				}
			},
			skipAuthHeaders: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			handler := func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.statusCode)
				if tc.responseBody != nil {
					_, _ = w.Write(tc.responseBody)
				}
			}

			var api *PodcastIndexAPI
			if tc.setupClient != nil {
				tempServer := httptest.NewServer(http.HandlerFunc(handler))
				api = tc.setupClient(tempServer.URL)
				api.HTTPClient = tempServer.Client()
				t.Cleanup(tempServer.Close)
			} else {
				api, _ = setupTestAPI(t, handler)
			}

			var result interface{}
			err := api.Get(context.Background(), "test/errorhandling", nil, &result)

			if err == nil {
				t.Fatalf("Expected an error for status %d, but got nil", tc.statusCode)
			}
			if !strings.Contains(err.Error(), tc.expectedErrMsg) {
				t.Errorf("Expected error message to contain '%s', but got '%s'", tc.expectedErrMsg, err.Error())
			}
		})
	}
}

var _ http.RoundTripper = &MockRoundTripper{}

type MockRoundTripper struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return m.RoundTripFunc(req)
}

func TestGetRequestFailure(t *testing.T) {
	mockRequestFailClient := &http.Client{
		Transport: &MockRoundTripper{
			RoundTripFunc: func(req *http.Request) (*http.Response, error) {
				return nil, fmt.Errorf("simulated request failure")
			},
		},
	}

	api := newTestAPI()
	api.HTTPClient = mockRequestFailClient

	var result interface{}
	err := api.Get(context.Background(), "test/requestfailure", nil, &result)

	if err == nil {
		t.Fatalf("Expected an error due to request failure, but got nil")
	}
	expectedErrMsg := "failed to execute HTTP request:"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrMsg, err.Error())
	}
}

func TestGetContextTimeout(t *testing.T) {
	api := newTestAPI()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()
	time.Sleep(2 * time.Millisecond)
	err := api.Get(ctx, "test/endpoint", nil, nil)

	if err == nil {
		t.Fatalf("Expected an error due to nil context, but got nil")
	}

	expectedErrMsg := "failed to execute HTTP request:"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrMsg, err.Error())
	}

	expectedWrappedMsg := "context deadline exceeded"
	if !strings.Contains(err.Error(), expectedWrappedMsg) {
		t.Errorf("Expected error message to contain wrapped error '%s', but got '%s'", expectedWrappedMsg, err.Error())
	}
}

func TestNoHTTPClient(t *testing.T) {
	api := newTestAPI()
	api.HTTPClient = nil
	err := api.Get(context.Background(), "test/endpoint", nil, nil)
	if err == nil {
		t.Fatalf("Expected an error due to nil http client, but got nil")
	}
	expectedErr := "HTTPClient is nil"
	if !strings.Contains(err.Error(), expectedErr) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErr, err.Error())
	}
}

func TestNilContext(t *testing.T) {
	api := newTestAPI() // Use the helper (default client is fine)
	//lint:ignore SA1012 - we are testing the nil context case
	err := api.Get(nil, "test/endpoint", nil, nil)
	if err == nil {
		t.Fatalf("Expected an error due to nil context, but got nil")
	}
	expectedErrMsg := "failed to execute HTTP request:"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrMsg, err.Error())
	}
}

// Added new test function TestGetResponseReadFailure
func TestGetResponseReadFailure(t *testing.T) {
	mockClientWithUnreadableBody := &http.Client{
		Transport: &MockRoundTripper{
			RoundTripFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusInternalServerError,
					Body:       io.NopCloser(&errorReader{}),
					Request:    req,
				}, nil
			},
		},
	}

	api := newTestAPI()
	api.HTTPClient = mockClientWithUnreadableBody

	var result interface{}
	err := api.Get(context.Background(), "test/readfailure", nil, &result)

	if err == nil {
		t.Fatalf("Expected an error due to response body read failure, but got nil")
	}

	// Check that the specific error message for read failure is returned
	expectedErrMsg := "failed to read error response: simulated read error"
	if !strings.Contains(err.Error(), expectedErrMsg) {
		t.Errorf("Expected error message to contain '%s', but got '%s'", expectedErrMsg, err.Error())
	}
}
