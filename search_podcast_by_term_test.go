package podcastindex

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"podcastindex/podcast/value"
	"testing"
)

// TestGetPodcastsByTitleRequest verifies that the client sends the correct
// request path and query parameters for the GetPodcastsByTitle method.
func TestGetPodcastsByTermRequest(t *testing.T) {
	t.Run("the client sends the correct request path and query parameters", func(t *testing.T) {
		searchServer := GetSearchServer(t)
		defer searchServer.Server.Close()

		serverURL, _ := url.Parse(searchServer.Server.URL)
		client := NewClient(NewClientOptions{
			BaseURL: serverURL,
		})

		term := "batman university"
		params := SearchPodcastsByTermParams{
			Max:      5,
			Clean:    true,
			FullText: true,
			Similar:  true,
			APOnly:   true,
			Value:    value.PaymentAny,
		}

		expectedQuery := url.Values{
			"q":        {term},
			"max":      {"5"},
			"clean":    {},
			"fulltext": {},
			"similar":  {},
			"val":      {value.PaymentAny},
			"aponly":   {"true"},
		}

		_, _ = client.SearchPodcastsByTerm(context.Background(), term, &params)
		if err := searchServer.ExpectPathAndQuery("/search/byterm", expectedQuery); err != nil {
			t.Error(err.Error())
		}
	})
	t.Run("the client errors if the server returns an error", func(t *testing.T) {
		errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer errorServer.Close()

		errorServerURL, _ := url.Parse(errorServer.URL)
		errorClient := NewClient(NewClientOptions{
			BaseURL: errorServerURL,
		})

		// Make sure an error is returned from the client when the server errors
		_, err := errorClient.SearchPodcastsByTerm(context.Background(), "error test", nil)
		if err == nil {
			t.Errorf("Expected an error when server returns status 500, but got nil")
		}
	})
}

func TestGetPodcastsByTermIntegration(t *testing.T) {
	client := authenticatedClient(t)
	_, err := client.SearchPodcastsByTerm(context.Background(), "batman university", nil)
	if err != nil {
		t.Fatalf("GetPodcastsByTerm failed: %v", err)
	}
}
