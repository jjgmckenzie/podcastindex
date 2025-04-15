package podcastindex

type searchResponse struct {
	Status      string     `json:"status"`
	Feeds       []*Podcast `json:"feeds"`
	Count       int        `json:"count"`
	Query       *string    `json:"query"`
	Description string     `json:"description"`
}
