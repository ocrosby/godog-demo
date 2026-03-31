package models

// Comment represents a comment on a post as returned by the JSONPlaceholder
// /comments endpoint.
type Comment struct {
	// PostID is the identifier of the post this comment belongs to.
	PostID int `json:"postId"`

	// ID is the unique identifier of the comment.
	ID int `json:"id"`

	// Name is the display name provided by the commenter.
	Name string `json:"name"`

	// Email is the email address of the commenter.
	Email string `json:"email"`

	// Body is the full text content of the comment.
	Body string `json:"body"`
}
