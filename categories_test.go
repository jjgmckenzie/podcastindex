package podcastindex

import (
	"context"
	"testing"
)

func TestCategoriesIntegration(t *testing.T) {
	t.Run("Integration test: Client should be able to get categories", func(t *testing.T) {
		client := authenticatedClient(t)
		categories, err := client.Categories(context.Background())
		if err != nil {
			t.Fatal("Failed to get categories", err)
		}
		if len(categories) == 0 {
			t.Fatal("No categories found")
		}
		t.Logf("Categories: %v", categories)
	})
	t.Run("Unauthenticated client should return an error", func(t *testing.T) {
		client := unauthenticatedClient()
		_, err := client.Categories(context.Background())
		t.Log(err)
		if err == nil {
			t.Fatal("Expected error to be returned")
		}
	})
}
