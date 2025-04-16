package podcastindex

import (
	"net/url"
	"podcastindex/episode"
	"podcastindex/podcast"
	"time"

	"golang.org/x/text/language"
)

// Episode is an episode from the PodcastIndex.org API
type Episode struct {
	// ID is the internal PodcastIndex.org episode ID.
	ID episode.ID
	// Title is the name of the feed
	Title string
	// Link is the channel-level link in the feed
	Link url.URL
	// Description is  The item-level description of the episode.
	Description string
	// GUID is the unique identifier for the episode
	GUID episode.GUID
	// DatePublished is the date and time the episode was published
	DatePublished time.Time
	// DatePublishedPretty The date and time the episode was published formatted as a human readable string.
	//
	// Note: uses the PodcastIndex server local time to do conversion.
	DatePublishedPretty string
	// DateCrawled is the time this episode was found in the feed
	DateCrawled time.Time
	// EnclosureURL is the URL/link to the episode file
	EnclosureURL url.URL
	// EnclosureType is the Content-Type for the item specified by the enclosureUrl
	EnclosureType string
	// EnclosureLength is the length of the enclosure in bytes
	EnclosureLength int
	// Explicit : is feed or episode marked as explicit
	Explicit bool
	// EpisodeNumber is the episode number, may be null.
	EpisodeNumber *int
	EpisodeType   *string
	Image         url.URL
	FeedITunesID  podcast.ITunesID
	FeedURL       url.URL
	FeedImage     url.URL
	FeedID        podcast.ID
	FeedLanguage  language.Tag
	ChaptersURL   url.URL
	TranscriptURL *url.URL
	Transcripts   *[]episode.Transcript
}
