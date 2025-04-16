package episode

// Soundbite is a segment of an episode
type Soundbite struct {
	// StartTime  The time where the soundbite begins in the item specified by the enclosureUrl
	StartTime int `json:"startTime"`
	// Duration  The duration of the soundbite in milliseconds
	Duration int `json:"duration"`
	// Title  The title of the soundbite
	Title string `json:"title"`
}
