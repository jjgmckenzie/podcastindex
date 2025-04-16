package podcastindex

type searchResponse struct {
	Status      string     `json:"status"`
	Feeds       []*Podcast `json:"feeds"`
	Count       int        `json:"count"`
	Query       *string    `json:"query"`
	Description string     `json:"description"`
}

type getPodcastResponse struct {
	// Status indicates API request status; either "true" or "false"
	Status string `json:"status"`
	// Query is the object containing the input query data
	Query struct{} `json:"query"`
	// Feed is the known details of podcast feed; type Podcast - this is the response.
	Feed Podcast `json:"feed"`
	//  Description is the description of the response
	Description string `json:"description"`
}

type getEpisodeResponse struct {
	// Status indicates API request status; either "true" or "false"
	Status string `json:"status"`
	// Query is the object containing the input query data
	Query string `json:"query"`
	// LiveItems is the List of live episodes for feed
	LiveItems []Episode `json:"liveItems"`
	// Items is the list of episodes matching request
	Items []Episode `json:"items"`
	Count int       `json:"count"`
	//  Description is the description of the response
	Description string `json:"description"`
}
