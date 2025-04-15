package podcastindex

import (
	"encoding/json"
	"fmt"
	"net/url"
	"podcastindex/podcast"
	"strconv"
	"time"

	"golang.org/x/text/language"
)

// Podcast contains all the information returned by the PodcastIndex API
//
// https://podcastindex-org.github.io/docs-api/#tag--Podcasts
// They can be retrieved by Client.GetPodcastByID(podcast.ID), Client.GetPodcastByITunesID(podcast.ITunesID), or
// Client.GetPodcastByGUID(GUID)
type Podcast struct {

	// ID is The internal PodcastIndex.org Feed ID.
	ID podcast.ID `json:"id"`
	// GUID is the GUID from the podcast:guid tag in the feed. This value is a unique, global identifier for the podcast.
	//
	// See the namespace spec for GUID for details.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#guid
	GUID podcast.GUID `json:"podcastGUID"`
	// Title is the name of the feed
	Title string `json:"title"`
	// URL is the current feed url
	URL url.URL `json:"url"`
	// OriginalURL is the url of the feed, before it changed to the current url value.
	OriginalURL url.URL `json:"originalURL"`
	// Link is the channel-level link in the feed
	Link url.URL `json:"link"`
	// Description is the channel-level description
	//
	// Uses the longer of the possible fields in the feed: <description>, <itunes:summary> and <content:encoded>
	Description string `json:"description"`
	// Author is the channel-level author element.
	//
	// Usually iTunes specific, but could be from another namespace if not present.
	Author string `json:"author"`
	// OwnerName is the channel-level owner:name element.
	//
	// Usually iTunes specific, but could be from another namespace if not present.
	OwnerName string `json:"ownerName"`
	// Image is the channel-level image element.
	Image url.URL `json:"image"`
	// Artwork is seemingly the best artwork we can find for the feed.
	//
	// Might be the same as image in most instances
	Artwork url.URL `json:"artwork"`
	// LastUpdateTime is the channel-level pubDate for the feed, if it's sane.
	//
	// If not, this is a heuristic value, arrived at by analyzing other parts of the feed, like item-level pubDates.
	LastUpdateTime time.Time `json:"lastUpdateTime"`
	// LastCrawlTime is the last time we attempted to pull this feed from its url.
	LastCrawlTime time.Time `json:"lastCrawlTime"`
	// LastParseTime is the last time we tried to parse the downloaded feed content.
	LastParseTime time.Time `json:"lastParseTime"`
	// InPollingQueue indicates if feed is currently scheduled to be polled/checked for new episodes.
	InPollingQueue *int `json:"inPollingQueue"`
	// Priority is How often the feed is checked for updates and new episodes
	//
	// A value of -1 means never check. A value of 5 means check the most.
	//
	// Allowed: -1┃0┃1┃2┃3┃4┃5
	Priority int `json:"priority"`
	// LastGoodHTTPStatusTime is the timestamp of the last time we got a "good", meaning non-4xx/non-5xx, status code when pulling this feed from its url.
	LastGoodHTTPStatusTime time.Time `json:"lastGoodHttpStatusTime"`
	// LastHTTPStatus is the last http status code we got when pulling this feed from its url.
	//
	// You will see some made up status codes sometimes. These are what we use to track state within the feed puller. These all start with 9xx.
	LastHTTPStatus int `json:"lastHTTPStatus"`
	// ContentType is The Content-Type header from the last time we pulled this feed from its url.
	ContentType string `json:"contentType"`
	// ITunesID is The iTunes id of this feed if there is one, and we know what it is.
	// Note this CAN be null if not found.
	ITunesID podcast.ITunesID `json:"itunesID"`
	// ITunesType is the type of iTunes feed.
	//
	// Possible values: episodic, serial
	ITunesType string `json:"itunesType"`
	// Generator is the channel-level generator element if there is one.
	Generator string `json:"generator"`
	// Language is the channel-level language specification of the feed.
	//
	//Languages accord with the RSS language Spec.
	Language language.Tag `json:"language"`
	// Explicit is whether the feed is marked as explicit
	Explicit bool `json:"explicit"`
	// Type of source feed where:
	//
	//    0: RSS - podcast.FeedRSS
	//    1: Atom - podcast.FeedAtom
	//
	// Allowed: 0┃1
	Type int `json:"type"`
	// Medium is the value of the podcast:medium attribute for the feed.
	//
	// See the medium description in the podcast namespace for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#medium
	Medium string `json:"medium"`
	// Dead: At some point, we give up trying to process a feed and mark it as dead. This is usually after 1000 errors without a successful pull/parse cycle. Once the feed is marked dead, we only check it once per month.
	Dead bool `json:"dead"`
	// EpisodeCount is the number of episodes in the feed known to the index.
	EpisodeCount int `json:"episodeCount"`
	// CrawlErrors is the number of errors we've encountered trying to pull a copy of the feed. Errors are things like a 500 or 404 response, a server timeout, bad encoding, etc.
	CrawlErrors int `json:"crawlErrors"`
	// ParseErrors is  The number of errors we've encountered trying to parse the feed content. Errors here are things like not well-formed xml, bad character encoding, etc.
	// We fix many of these types of issues on the fly when parsing. We only increment the errors count when we can't fix it.
	ParseErrors int `json:"parseErrors"`
	// Categories is an array of categories, where the index is Category ID, and the value is Category Name.
	// All Category numbers and names are returned by the categories/list endpoint.
	Categories []podcast.Category `json:"categories"`
	// Locked: Tell other podcast platforms whether they are allowed to import this feed. A value of true means that
	// any attempt to import this feed into a new platform should be rejected.
	Locked bool `json:"locked"`
	// ImageURLHash is a CRC32 hash of the image url with the protocol (http://, https://) removed. 64bit integer.
	ImageURLHash int `json:"imageURLHash"`
	// NewestItemPubDate is the time the most recent episode in the feed was published.
	// Note: some endpoints use newestItemPubdate while others use newestItemPublishTime. They return the same information.
	// See https://github.com/Podcastindex-org/api/issues/3 to track when the property name is updated.
	NewestItemPubDate *time.Time `json:"newestItemPubDate"`
	// Value is the "Value for Value" payment information for the podcast. Will be nil if not reported.
	Value *podcast.Value `json:"value"`
}

// podcastJSON is an intermediary struct used for unmarshalling Podcast data,
// handling Unix timestamps for time fields, converting category IDs to ints, and parsing URLs.
type podcastJSON struct {
	ID                     int               `json:"id"`
	GUID                   string            `json:"podcastGUID"`
	Title                  string            `json:"title"`
	URL                    string            `json:"url"`
	OriginalURL            string            `json:"originalURL"`
	Link                   string            `json:"link"`
	Description            string            `json:"description"`
	Author                 string            `json:"author"`
	OwnerName              string            `json:"ownerName"`
	Image                  string            `json:"image"`
	Artwork                string            `json:"artwork"`
	LastUpdateTime         int64             `json:"lastUpdateTime"`
	LastCrawlTime          int64             `json:"lastCrawlTime"`
	LastParseTime          int64             `json:"lastParseTime"`
	LastGoodHTTPStatusTime int64             `json:"lastGoodHttpStatusTime"`
	LastHTTPStatus         int               `json:"lastHTTPStatus"`
	ContentType            string            `json:"contentType"`
	ITunesID               *int              `json:"itunesID"`
	ITunesType             string            `json:"itunesType"`
	Generator              string            `json:"generator"`
	Language               string            `json:"language"`
	Explicit               bool              `json:"explicit"`
	Type                   int               `json:"type"`
	Medium                 string            `json:"medium"`
	Dead                   int               `json:"dead"`
	EpisodeCount           int               `json:"episodeCount"`
	CrawlErrors            int               `json:"crawlErrors"`
	ParseErrors            int               `json:"parseErrors"`
	InPollingQueue         *int              `json:"inPollingQueue"`
	Priority               int               `json:"priority"`
	Categories             map[string]string `json:"categories"`
	Locked                 int               `json:"locked"`
	ImageURLHash           int               `json:"imageURLHash"`
	NewestItemPubDate      int64             `json:"newestItemPubDate"`
	Value                  *podcast.Value    `json:"value"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for Podcast.
// It handles the conversion of Unix timestamps (int64) to time.Time,
// as well as converting category IDs to ints, mistyped booleans, and parsing URLs.
func (p *Podcast) UnmarshalJSON(data []byte) error {
	var aux podcastJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// --- Direct field assignments & simple conversions ---
	p.ID = podcast.ID(aux.ID)
	p.GUID = podcast.GUID(aux.GUID)
	p.Title = aux.Title
	p.Description = aux.Description
	p.Author = aux.Author
	p.OwnerName = aux.OwnerName
	p.LastHTTPStatus = aux.LastHTTPStatus
	p.ContentType = aux.ContentType
	if aux.ITunesID != nil {
		p.ITunesID = podcast.ITunesID(*aux.ITunesID)
	}
	p.ITunesType = aux.ITunesType
	p.Generator = aux.Generator
	p.Language = language.Make(aux.Language)
	p.Explicit = aux.Explicit
	p.Type = aux.Type
	p.Medium = aux.Medium
	p.Dead = aux.Dead == 1 // Convert int to bool
	p.EpisodeCount = aux.EpisodeCount
	p.CrawlErrors = aux.CrawlErrors
	p.ParseErrors = aux.ParseErrors
	p.InPollingQueue = aux.InPollingQueue
	p.Priority = aux.Priority
	p.Locked = aux.Locked == 1 // Convert int to bool
	p.ImageURLHash = aux.ImageURLHash
	p.Value = aux.Value // Assign pointer directly

	// --- URL parsing ---
	var err error
	var parsedURL *url.URL

	parsedURL, err = url.Parse(aux.URL)
	if err != nil {
		return fmt.Errorf("failed to parse URL '%s': %w", aux.URL, err)
	}
	p.URL = *parsedURL

	parsedURL, err = url.Parse(aux.OriginalURL)
	if err != nil {
		return fmt.Errorf("failed to parse OriginalURL '%s': %w", aux.OriginalURL, err)
	}
	p.OriginalURL = *parsedURL

	parsedURL, err = url.Parse(aux.Link)
	if err != nil {
		return fmt.Errorf("failed to parse Link '%s': %w", aux.Link, err)
	}
	p.Link = *parsedURL

	parsedURL, err = url.Parse(aux.Image)
	if err != nil {
		return fmt.Errorf("failed to parse Image URL '%s': %w", aux.Image, err)
	}
	p.Image = *parsedURL

	parsedURL, err = url.Parse(aux.Artwork)
	if err != nil {
		return fmt.Errorf("failed to parse Artwork URL '%s': %w", aux.Artwork, err)
	}
	p.Artwork = *parsedURL

	// --- Category mapping ---
	categories := make([]podcast.Category, 0, len(aux.Categories))
	for id, name := range aux.Categories {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return fmt.Errorf("failed to convert category ID '%s' to int: %w", id, err)
		}
		categories = append(categories, podcast.Category{ID: podcast.CategoryID(idInt), Name: name})
	}
	p.Categories = categories

	// --- Time conversions (Unix timestamp to time.Time) ---
	p.LastUpdateTime = time.Unix(aux.LastUpdateTime, 0)
	p.LastCrawlTime = time.Unix(aux.LastCrawlTime, 0)
	p.LastParseTime = time.Unix(aux.LastParseTime, 0)
	p.LastGoodHTTPStatusTime = time.Unix(aux.LastGoodHTTPStatusTime, 0)
	newestItemPubDateTime := time.Unix(aux.NewestItemPubDate, 0)
	p.NewestItemPubDate = &newestItemPubDateTime

	return nil
}
