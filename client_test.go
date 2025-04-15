package podcastindex

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// getKeyAndSecretFromEnv returns the API key and secret from the environment variables.
// If the environment variables are not set, it will load the .env file and return the values from there.
// If the .env file is not found, it will skip the test.
func getKeyAndSecretFromEnv(t *testing.T) (string, string) {
	key := os.Getenv("PODCASTINDEX_API_KEY")
	secret := os.Getenv("PODCASTINDEX_API_SECRET")

	if key == "" || secret == "" {
		err := godotenv.Load(".env")
		if err != nil {
			t.Skip("PODCASTINDEX_API_KEY and PODCASTINDEX_API_SECRET not set, skipping test")
		}
		key = os.Getenv("PODCASTINDEX_API_KEY")
		secret = os.Getenv("PODCASTINDEX_API_SECRET")
	}
	return key, secret
}

func authenticatedClient(t *testing.T) *Client {
	key, secret := getKeyAndSecretFromEnv(t)
	return NewClient(NewClientOptions{
		UserAgent: "podcastindex-go/testrunner",
		APIKey:    key,
		APISecret: secret,
	})
}

func unauthenticatedClient() *Client {
	return NewClient(NewClientOptions{
		UserAgent: "podcastindex-go/testrunner",
	})
}

func TestPodcastIndexClientIntegration(t *testing.T) {
	client := authenticatedClient(t)
	t.Run("Client should be able to make a request which requires authentication", func(t *testing.T) {
		// arbitrary request to check if the client is working
		var response struct{}
		err := client.api.Get(context.Background(), "categories/list", nil, &response)
		if err != nil {
			t.Fatal("Failed to make request", err)
		}
	})
}
