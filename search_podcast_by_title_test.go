package podcastindex

import (
	"context"
	"net/url"
	"podcastindex/podcast/value"
	"testing"
)

// TestGetPodcastsByTitleRequest verifies that the client sends the correct
// request path and query parameters for the GetPodcastsByTitle method.
func TestGetPodcastsByTitleRequest(t *testing.T) {
	t.Run("the client sends the correct request path and query parameters", func(t *testing.T) {
		searchServer := GetSearchServer(t)
		defer searchServer.Server.Close()

		serverURL, _ := url.Parse(searchServer.Server.URL)
		client := NewClient(NewClientOptions{
			BaseURL: serverURL,
		})

		title := "batman university"
		params := SearchPodcastsByTitleParams{
			Max:      5,
			Clean:    true,
			FullText: true,
			Similar:  true,
			Value:    value.PaymentAny,
		}

		_, _ = client.SearchPodcastsByTitle(context.Background(), title, &params)

		expectedQuery := url.Values{
			"q":        {title},
			"max":      {"5"},
			"clean":    {},
			"fulltext": {},
			"similar":  {},
			"val":      {value.PaymentAny},
		}

		if err := searchServer.ExpectPathAndQuery("/search/bytitle", expectedQuery); err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("the client truncates Max parameter to 99 when it's >= 100", func(t *testing.T) {
		searchServer := GetSearchServer(t)
		defer searchServer.Server.Close()

		serverURL, _ := url.Parse(searchServer.Server.URL)
		client := NewClient(NewClientOptions{
			BaseURL: serverURL,
		})

		title := "batman university"
		params := SearchPodcastsByTitleParams{
			Max: 100, // Testing with Max = 100
		}

		_, _ = client.SearchPodcastsByTitle(context.Background(), title, &params)

		// Verify that the max param was truncated to 99
		expectedQuery := url.Values{
			"q":   {title},
			"max": {"99"}, // Should be truncated to 99
		}

		if err := searchServer.ExpectPathAndQuery("/search/bytitle", expectedQuery); err != nil {
			t.Error(err.Error())
		}

		// Also test with a higher value
		params = SearchPodcastsByTitleParams{
			Max: 1000, // Testing with Max = 1000
		}

		_, _ = client.SearchPodcastsByTitle(context.Background(), title, &params)

		// Verify that the max param was truncated to 99
		if err := searchServer.ExpectPathAndQuery("/search/bytitle", expectedQuery); err != nil {
			t.Error(err.Error())
		}
	})

	t.Run("the client errors if the server returns an error", func(t *testing.T) {
		// Make sure an error is returned from the client when the server errors
		_, err := GetErrorServer(t).SearchPodcastsByTitle(context.Background(), "error test", nil)
		if err == nil {
			t.Errorf("Expected an error when server returns status 500, but got nil")
		}
	})
}

func TestGetPodcastsByTitleIntegration(t *testing.T) {
	client := authenticatedClient(t)
	_, err := client.SearchPodcastsByTitle(context.Background(), "batman university", nil)
	if err != nil {
		t.Fatalf("GetPodcastsByTitle failed: %v", err)
	}
}
