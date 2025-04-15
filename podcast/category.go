package podcast

type CategoryID int

// Category is a category of a podcast.
//
// https://podcastindex-org.github.io/docs-api/#tag--Categories
type Category struct {
	// ID is the internal podcastindex ID of the category.
	ID CategoryID `json:"id"`
	// Name is the name of the category.
	Name string `json:"name"`
}
