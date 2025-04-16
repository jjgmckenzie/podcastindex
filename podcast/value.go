package podcast

import (
	"github.com/jjgmckenzie/podcastindex/podcast/value"
)

// Value contains the information for supporting the podcast via one of the
// "Value for Value" methods from the PodcastIndex ID.
type Value struct {
	// Model is the description of the method for providing "Value for Value" payments
	Model value.Model `json:"model"`
	// Destinations is the list of destinations where "Value for Value" payments should be sent.
	Destinations []value.Destination `json:"destinations"`
}
