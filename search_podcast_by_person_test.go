package podcastindex

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// TestGetPodcastsByTitleRequest verifies that the client sends the correct
// request path and query parameters for the GetPodcastsByTitle method.
func TestGetPodcastsByPersonRequest(t *testing.T) {
	t.Run("the client sends the correct request path and query parameters", func(t *testing.T) {
		searchServer := GetSearchServer(t)
		defer searchServer.Server.Close()

		serverURL, _ := url.Parse(searchServer.Server.URL)
		client := NewClient(NewClientOptions{
			BaseURL: serverURL,
		})

		person := "batman university"
		params := SearchPodcastsByPersonParams{
			Max:      5,
			FullText: true,
		}

		expectedQuery := url.Values{
			"q":        {person},
			"max":      {"5"},
			"fulltext": {},
		}
		
		_, _ = client.SearchPodcastsByPerson(context.Background(), person, &params)
		if err := searchServer.ExpectPathAndQuery("/search/byperson", expectedQuery); err != nil {
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
		_, err := errorClient.SearchPodcastsByPerson(context.Background(), "error test", nil)
		if err == nil {
			t.Errorf("Expected an error when server returns status 500, but got nil")
		}
	})
}

func TestGetPodcastsByPersonIntegration(t *testing.T) {
	client := authenticatedClient(t)
	_, err := client.SearchPodcastsByPerson(context.Background(), "batman university", nil)
	if err != nil {
		t.Fatalf("GetPodcastsByPerson failed: %v", err)
	}
}
