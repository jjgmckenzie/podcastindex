package podcastindex

import (
	"encoding/json"
	"net/url"
	"os"
	"podcastindex/podcast"
	"reflect"
	"sort"
	"testing"
	"time"

	"golang.org/x/text/language"
)

// TestUnmarshalPodcastFromTitleJSON verifies that the JSON data from the
// /podcasts/bytitle endpoint (or similar search endpoints returning podcast feeds)
// can be correctly unmarshalled into the Podcast struct.
func TestUnmarshalPodcastFromTitleJSON(t *testing.T) {
	// Helper to parse URLs and handle errors inline
	parseURL := func(rawURL string) url.URL {
		u, err := url.Parse(rawURL)
		if err != nil {
			t.Fatalf("Failed to parse URL '%s': %v", rawURL, err)
		}
		return *u
	}

	// Define expected time values based on the JSON data
	lastUpdateTime := time.Unix(1744427062, 0).UTC()
	lastCrawlTime := time.Unix(1744427030, 0).UTC()
	lastParseTime := time.Unix(1744427084, 0).UTC()
	lastGoodHTTPStatusTime := time.Unix(1744427030, 0).UTC()
	newestItemPubDate := time.Unix(1546399813, 0).UTC()

	// Define the expected Go struct slice matching the JSON structure
	expectedPodcasts := []*Podcast{
		{
			ID:                     podcast.ID(75075),
			Title:                  "Batman University",
			URL:                    parseURL("https://feeds.theincomparable.com/batmanuniversity"),
			OriginalURL:            parseURL("https://feeds.theincomparable.com/batmanuniversity"),
			Link:                   parseURL("https://www.theincomparable.com/batmanuniversity/"),
			Description:            `Batman University is a seasonal podcast about you know who. It began with an analysis of episodes of “Batman: The Animated Series” but has now expanded to cover other series, movies, and media. Your professor is Tony Sindelar.`,
			Author:                 "Tony Sindelar",
			OwnerName:              "",
			Image:                  parseURL("https://www.theincomparable.com/imgs/logos/logo-batmanuniversity-3x.jpg"),
			Artwork:                parseURL("https://www.theincomparable.com/imgs/logos/logo-batmanuniversity-3x.jpg"),
			LastUpdateTime:         lastUpdateTime,
			LastCrawlTime:          lastCrawlTime,
			LastParseTime:          lastParseTime,
			InPollingQueue:         nil, // Assuming nil if not present or 0/false
			Priority:               0,   // Assuming 0 if not present
			LastGoodHTTPStatusTime: lastGoodHTTPStatusTime,
			LastHTTPStatus:         200,
			ContentType:            "application/rss+xml",
			Language:               language.Make("en"),
			Explicit:               false,
			EpisodeCount:           19,
			ITunesID:               podcast.ITunesID(1441923632),
			Generator:              "", // Assuming empty if not present or null
			Categories: []podcast.Category{
				{ID: 104, Name: "Tv"},
				{ID: 105, Name: "Film"},
				{ID: 107, Name: "Reviews"},
			},
			GUID:              podcast.GUID("ac9907f2-a748-59eb-a799-88a9c8bfb9f5"),
			ITunesType:        "", // Assuming empty if not present or null
			Type:              0,
			Medium:            "podcast",
			Dead:              false,
			CrawlErrors:       0,
			ParseErrors:       0,
			Locked:            false,
			ImageURLHash:      1702747127,
			NewestItemPubDate: &newestItemPubDate,
			Value:             nil, // Assuming nil if not present
		},
	}

	// Read the mock JSON from file
	jsonBytes, err := os.ReadFile("testdata/example_podcasts_by_title.json")
	if err != nil {
		t.Fatalf("Failed to read mock JSON file: %v", err)
	}

	// Define a struct that matches the overall JSON response structure
	// This assumes the podcasts are nested under a "feeds" key
	var response struct {
		Status      string     `json:"status"`
		Feeds       []*Podcast `json:"feeds"`
		Count       int        `json:"count"`
		Query       string     `json:"query"` // Or map[string]string if query is complex
		Description string     `json:"description"`
	}

	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	actualPodcasts := response.Feeds
	for _, p := range actualPodcasts {
		p.LastUpdateTime = p.LastUpdateTime.UTC()
		p.LastCrawlTime = p.LastCrawlTime.UTC()
		p.LastParseTime = p.LastParseTime.UTC()
		p.LastGoodHTTPStatusTime = p.LastGoodHTTPStatusTime.UTC()
		if p.NewestItemPubDate != nil {
			newestPubDate := p.NewestItemPubDate.UTC()
			p.NewestItemPubDate = &newestPubDate
		}
		sort.Slice(p.Categories, func(i, j int) bool {
			return p.Categories[i].ID < p.Categories[j].ID
		})
	}

	sort.Slice(expectedPodcasts[0].Categories, func(i, j int) bool {
		return expectedPodcasts[0].Categories[i].ID < expectedPodcasts[0].Categories[j].ID
	})

	if !reflect.DeepEqual(actualPodcasts, expectedPodcasts) {
		t.Errorf("Unmarshalled podcast data does not match expected data.\nExpected: %#v\nGot:      %#v", expectedPodcasts[0], actualPodcasts[0])
	}
}
