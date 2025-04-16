package episode

import "net/url"

type PersonID int

type Person struct {
	// ID is the internal PodcastIndex.org episode.PersonID.
	ID PersonID `json:"id"`
	// Name is the name of the person.
	//
	// See the podcast namespace spec for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
	Name string `json:"name"`
	// Role is used to identify what role the person serves on the show or episode.
	//
	// value should be an official role from https://github.com/Podcastindex-org/podcast-namespace/blob/main/taxonomy.json
	//
	// See the podcast namespace spec for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
	Role string `json:"role"`
	// Group is used to identify what role the person serves on the show or episode.
	//
	// value should be an official group from https://github.com/Podcastindex-org/podcast-namespace/blob/main/taxonomy.json
	//
	// See the podcast namespace spec for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
	Group string `json:"group"`
	// Href is the url to a relevant resource of information about the person, such as a homepage or third-party profile platform.
	//
	// See the podcast namespace spec for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
	Href url.URL `json:"href"`
	// Image is the URL to a picture or avatar of the person.
	//
	// See the podcast namespace spec for more information.
	// https://github.com/Podcastindex-org/podcast-namespace/blob/main/docs/1.0.md#person
	Image url.URL `json:"img"`
}
