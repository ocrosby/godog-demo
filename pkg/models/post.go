package models

// Post represents a blog post as returned by the JSONPlaceholder /posts endpoint.
type Post struct {
	// UserID is the identifier of the user who authored the post.
	UserID int `json:"userId"`

	// ID is the unique identifier of the post.
	ID int `json:"id"`

	// Title is the headline of the post.
	Title string `json:"title"`

	// Body is the full text content of the post.
	Body string `json:"body"`
}
