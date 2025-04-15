package value

type Destination struct {
	// Name is the name of the destination
	Name string `json:"name"`
	// Address is the address of the destination
	Address string `json:"address"`
	// Type is the type of the destination
	Type string `json:"type"`
	// Split is the split of the destination
	Split int `json:"split"`
	// Fee indicates if destination is included due to a fee being charged. May not be reported.
	Fee *bool `json:"fee"`
	// CustomKey is the name of a custom record key to send along with the payment. May not be reported.
	//
	// See the podcast namespace spec and value specification for more information.
	//
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/value/value.md
	CustomKey *string `json:"customKey"`
	// CustomValue is the custom value of the destination
	//
	// See the podcast namespace spec and value specification for more information.
	//
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#value
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/value/value.md
	CustomValue *string `json:"customValue"`
}
