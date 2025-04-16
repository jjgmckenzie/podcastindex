package podcastindex

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/jjgmckenzie/podcastindex/episode"
	"github.com/jjgmckenzie/podcastindex/internal"
	"github.com/jjgmckenzie/podcastindex/podcast"

	"golang.org/x/text/language"
)

// Episode is an episode from the PodcastIndex.org API
type Episode struct {
	// ID is the internal PodcastIndex.org episode.ID.
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
	// EpisodeType is the type of episode. May be null for liveItem.
	//
	// Could be episode.EpisodeFull episode.EpisodeTrailer episode.EpisodeBonus
	EpisodeType *episode.EpisodeType
	// Season is the season number. May be null for liveItem.
	Season *int
	// Image is the item-level image for the episode
	Image url.URL
	// FeedITunesID is the iTunes id of this feed if there is one, and we know what it is.
	FeedITunesID *podcast.ITunesID
	// FeedURL is the Current feed URL
	FeedURL url.URL
	// FeedImage is the channel-level image element.
	FeedImage url.URL
	// FeedID is the internal PodcastIndex.org Feed ID.
	FeedID podcast.ID
	// FeedGUID is the podcast.GUID from the podcast:guid tag in the feed. This value is a unique, global identifier for the podcast.
	FeedGUID podcast.GUID
	// FeedLanguage is the channel-level language specification of the feed.
	FeedLanguage language.Tag
	// FeedDead : At some point, we give up trying to process a feed and mark it as dead. This is usually after 1000 errors without a successful pull/parse cycle. Once the feed is marked dead, we only check it once per month.
	FeedDead bool
	// FeedDuplicateOf :The internal PodcastIndex.org Feed id this feed duplicates. May be null except in podcasts/dead.
	FeedDuplicateOf *podcast.ID
	// ChaptersURL is the link to the JSON file containing the episode chapters
	ChaptersURL *url.URL
	// TranscriptURL is the link to the file containing the episode transcript
	//
	// NB: In most cases, Transcripts should be used instead.
	TranscriptURL *url.URL
	// Transcripts : list of transcripts for the episode. May not be reported.
	Transcripts *[]episode.Transcript
	// Soundbite is the soundbite for episode. May not be reported.
	Soundbite *episode.Soundbite
	// Soundbites are the soundbites for episode. May not be reported.
	Soundbites *[]episode.Soundbite
	// Persons is the List of people with an interest in this episode. May not be reported.
	//
	// See the podcast namespace spec for more information.
	Persons *[]episode.Person
	// SocialInteract is a lList the social interact data found in the podcast feed. May not be reported.
	//
	// See the podcast namespace spec for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#social-interact
	SocialInteract *[]episode.SocialInteract
	// Value is the Information for supporting the podcast via one of the "Value for Value" methods. May not be reported.
	Value *podcast.Value
	// live items also have these additional variables:

	// LivestreamStatus is the status of the livestream
	// Allowed: ended┃live
	// Is null if not live.
	LivestreamStatus *episode.LivestreamStatus
	// StartTime is the time the livestream starts
	StartTime *time.Time
	// EndTime is the time the livestream ends
	EndTime *time.Time
	// ContentLink (???)
	ContentLink *string
	// Duration is the duration of the episode in seconds
	Duration *int
}

// episodeJSON is an intermediary struct used for unmarshalling Episode data,
// handling Unix timestamps, URLs, and other data types that need conversion.
type episodeJSON struct {
	ID                  int                       `json:"id"`
	Title               string                    `json:"title"`
	Link                string                    `json:"link"`
	Description         string                    `json:"description"`
	GUID                string                    `json:"guid"`
	DatePublished       int64                     `json:"datePublished"`
	DatePublishedPretty string                    `json:"datePublishedPretty"`
	DateCrawled         int64                     `json:"dateCrawled"`
	EnclosureURL        string                    `json:"enclosureUrl"`
	EnclosureType       string                    `json:"enclosureType"`
	EnclosureLength     int                       `json:"enclosureLength"`
	Explicit            int                       `json:"explicit"`
	Episode             *int                      `json:"episode"`
	EpisodeType         *string                   `json:"episodeType"`
	Season              *int                      `json:"season,omitempty"`
	Image               string                    `json:"image"`
	FeedITunesID        *int                      `json:"feedItunesId,omitempty"`
	FeedURL             string                    `json:"feedUrl,omitempty"`
	FeedImage           string                    `json:"feedImage"`
	FeedID              int                       `json:"feedId"`
	PodcastGUID         string                    `json:"podcastGuid,omitempty"`
	FeedLanguage        string                    `json:"feedLanguage"`
	FeedDead            int                       `json:"feedDead"`
	FeedDuplicateOf     *int                      `json:"feedDuplicateOf"`
	ChaptersURL         *string                   `json:"chaptersUrl"`
	TranscriptURL       *string                   `json:"transcriptUrl"`
	Transcripts         *[]episode.Transcript     `json:"transcripts,omitempty"`
	Soundbite           *episode.Soundbite        `json:"soundbite,omitempty"`
	Soundbites          *[]episode.Soundbite      `json:"soundbites,omitempty"`
	Persons             *[]episode.Person         `json:"persons,omitempty"`
	SocialInteract      *[]episode.SocialInteract `json:"socialInteract,omitempty"`
	Value               *podcast.Value            `json:"value,omitempty"`
	LivestreamStatus    *string                   `json:"status,omitempty"`
	StartTime           *int64                    `json:"startTime,omitempty"`
	EndTime             *int64                    `json:"endTime,omitempty"`
	ContentLink         *string                   `json:"contentLink,omitempty"`
	Duration            *int                      `json:"duration,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for Episode.
// It handles the conversion of Unix timestamps to time.Time, URLs and other data conversions.
func (e *Episode) UnmarshalJSON(data []byte) error {
	var aux episodeJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Direct field assignments
	e.ID = episode.ID(aux.ID)
	e.Title = aux.Title
	e.Description = aux.Description
	e.GUID = episode.GUID(aux.GUID)
	e.DatePublishedPretty = aux.DatePublishedPretty
	e.EnclosureType = aux.EnclosureType
	e.EnclosureLength = aux.EnclosureLength
	e.Explicit = aux.Explicit == 1
	e.EpisodeNumber = aux.Episode
	e.Season = aux.Season
	e.FeedID = podcast.ID(aux.FeedID)
	if aux.PodcastGUID != "" {
		e.FeedGUID = podcast.GUID(aux.PodcastGUID)
	}
	e.FeedLanguage = language.Make(aux.FeedLanguage)
	e.FeedDead = aux.FeedDead == 1
	if aux.FeedDuplicateOf != nil {
		feedDupID := podcast.ID(*aux.FeedDuplicateOf)
		e.FeedDuplicateOf = &feedDupID
	}
	e.ContentLink = aux.ContentLink
	e.Duration = aux.Duration

	// Handle complex nested types that are managed elsewhere
	e.Transcripts = aux.Transcripts
	e.Soundbite = aux.Soundbite
	e.Soundbites = aux.Soundbites
	e.Persons = aux.Persons
	e.SocialInteract = aux.SocialInteract
	e.Value = aux.Value

	if aux.EpisodeType != nil {
		episodeType := episode.EpisodeType(*aux.EpisodeType)
		e.EpisodeType = &episodeType
	}

	if aux.LivestreamStatus != nil {
		livestreamStatus := episode.LivestreamStatus(*aux.LivestreamStatus)
		e.LivestreamStatus = &livestreamStatus
	}

	// Handle iTunes ID
	if aux.FeedITunesID != nil {
		itunesID := podcast.ITunesID(strconv.Itoa(*aux.FeedITunesID))
		e.FeedITunesID = &itunesID
	}

	// Time conversions
	e.DatePublished = time.Unix(aux.DatePublished, 0)
	e.DateCrawled = time.Unix(aux.DateCrawled, 0)

	if aux.StartTime != nil {
		startTime := time.Unix(*aux.StartTime, 0)
		e.StartTime = &startTime
	}

	if aux.EndTime != nil {
		endTime := time.Unix(*aux.EndTime, 0)
		e.EndTime = &endTime
	}

	// URL parsing
	var err error
	var parsedURL *url.URL

	// Parse Link URL
	parsedURL, err = url.Parse(aux.Link)
	if err != nil {
		return fmt.Errorf("failed to parse Link '%s': %w", aux.Link, err)
	}
	e.Link = *parsedURL

	// Parse EnclosureURL
	parsedURL, err = url.Parse(aux.EnclosureURL)
	if err != nil {
		return fmt.Errorf("failed to parse EnclosureURL '%s': %w", aux.EnclosureURL, err)
	}
	e.EnclosureURL = *parsedURL

	// Parse Image URL
	parsedURL, err = url.Parse(aux.Image)
	if err != nil {
		return fmt.Errorf("failed to parse Image URL '%s': %w", aux.Image, err)
	}
	e.Image = *parsedURL

	// Parse FeedImage URL
	parsedURL, err = url.Parse(aux.FeedImage)
	if err != nil {
		return fmt.Errorf("failed to parse FeedImage URL '%s': %w", aux.FeedImage, err)
	}
	e.FeedImage = *parsedURL

	// Parse FeedURL if present
	if aux.FeedURL != "" {
		parsedURL, err = url.Parse(aux.FeedURL)
		if err != nil {
			return fmt.Errorf("failed to parse FeedURL '%s': %w", aux.FeedURL, err)
		}
		e.FeedURL = *parsedURL
	}

	// Parse ChaptersURL if present
	if aux.ChaptersURL != nil {
		parsedURL, err = url.Parse(*aux.ChaptersURL)
		if err != nil {
			return fmt.Errorf("failed to parse ChaptersURL '%s': %w", *aux.ChaptersURL, err)
		}
		e.ChaptersURL = parsedURL
	}

	// Parse TranscriptURL if present
	if aux.TranscriptURL != nil {
		parsedURL, err = url.Parse(*aux.TranscriptURL)
		if err != nil {
			return fmt.Errorf("failed to parse TranscriptURL '%s': %w", *aux.TranscriptURL, err)
		}
		e.TranscriptURL = parsedURL
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface for Episode.
// It converts time.Time to Unix timestamps, URLs to strings, and handles other data conversions.
func (e *Episode) MarshalJSON() ([]byte, error) {

	aux := episodeJSON{
		ID:                  int(e.ID),
		Title:               e.Title,
		Link:                e.Link.String(),
		Description:         e.Description,
		GUID:                string(e.GUID),
		DatePublished:       e.DatePublished.Unix(),
		DatePublishedPretty: e.DatePublishedPretty,
		DateCrawled:         e.DateCrawled.Unix(),
		EnclosureURL:        e.EnclosureURL.String(),
		EnclosureType:       e.EnclosureType,
		EnclosureLength:     e.EnclosureLength,
		Explicit:            internal.BoolToInt(e.Explicit),
		Image:               e.Image.String(),
		FeedImage:           e.FeedImage.String(),
		FeedID:              int(e.FeedID),
		FeedLanguage:        e.FeedLanguage.String(),
		FeedDead:            internal.BoolToInt(e.FeedDead),
		Transcripts:         e.Transcripts,
		Soundbite:           e.Soundbite,
		Soundbites:          e.Soundbites,
		Persons:             e.Persons,
		SocialInteract:      e.SocialInteract,
		Value:               e.Value,
		ContentLink:         e.ContentLink,
		Duration:            e.Duration,
	}

	// Handle optional fields
	if e.EpisodeNumber != nil {
		aux.Episode = e.EpisodeNumber
	}

	if e.Season != nil {
		aux.Season = e.Season
	}

	if e.FeedITunesID != nil {
		itunesID, err := strconv.Atoi(string(*e.FeedITunesID))
		if err != nil {
			return nil, fmt.Errorf("failed to convert ITunesID to int: %w", err)
		}
		aux.FeedITunesID = &itunesID
	}

	if e.FeedURL.String() != "" {
		aux.FeedURL = e.FeedURL.String()
	}

	if e.FeedGUID != "" {
		aux.PodcastGUID = string(e.FeedGUID)
	}

	if e.FeedDuplicateOf != nil {
		dupID := int(*e.FeedDuplicateOf)
		aux.FeedDuplicateOf = &dupID
	}

	if e.ChaptersURL != nil {
		chaptersURL := e.ChaptersURL.String()
		aux.ChaptersURL = &chaptersURL
	}

	if e.TranscriptURL != nil {
		transcriptURL := e.TranscriptURL.String()
		aux.TranscriptURL = &transcriptURL
	}

	if e.EpisodeType != nil {
		episodeType := string(*e.EpisodeType)
		aux.EpisodeType = &episodeType
	}

	if e.LivestreamStatus != nil {
		status := string(*e.LivestreamStatus)
		aux.LivestreamStatus = &status
	}

	if e.StartTime != nil {
		startTime := e.StartTime.Unix()
		aux.StartTime = &startTime
	}

	if e.EndTime != nil {
		endTime := e.EndTime.Unix()
		aux.EndTime = &endTime
	}

	return json.Marshal(aux)
}
