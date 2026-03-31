// Package builders provides Builder pattern implementations for constructing
// domain model objects used in BDD scenario steps.
//
// Each builder accumulates field values through a fluent With* API and
// validates required fields in Build(), separating object construction from
// the step that sends the built object to the API. This makes the required vs
// optional field contract explicit and testable independently of the BDD layer.
package builders

import (
	"errors"

	"github.com/ocrosby/godog-demo/pkg/models"
)

// CommentBuilder accumulates field values for a Comment before creating it.
// Use NewCommentBuilder to obtain an instance, call the With* methods to set
// fields, and call Build to obtain a validated *models.Comment.
type CommentBuilder struct {
	postID int
	id     int
	name   string
	email  string
	body   string
}

// NewCommentBuilder returns a new, empty CommentBuilder.
func NewCommentBuilder() *CommentBuilder {
	return &CommentBuilder{}
}

// WithPostID sets the PostID field and returns the builder for method chaining.
func (b *CommentBuilder) WithPostID(postID int) *CommentBuilder {
	b.postID = postID
	return b
}

// WithID sets the ID field and returns the builder for method chaining.
func (b *CommentBuilder) WithID(id int) *CommentBuilder {
	b.id = id
	return b
}

// WithName sets the Name field and returns the builder for method chaining.
func (b *CommentBuilder) WithName(name string) *CommentBuilder {
	b.name = name
	return b
}

// WithEmail sets the Email field and returns the builder for method chaining.
func (b *CommentBuilder) WithEmail(email string) *CommentBuilder {
	b.email = email
	return b
}

// WithBody sets the Body field and returns the builder for method chaining.
func (b *CommentBuilder) WithBody(body string) *CommentBuilder {
	b.body = body
	return b
}

// Build validates the accumulated state and returns a new *models.Comment.
// It returns an error if PostID is zero, as PostID is required by
// JSONPlaceholder's POST /comments endpoint.
func (b *CommentBuilder) Build() (*models.Comment, error) {
	if b.postID == 0 {
		return nil, errors.New("comment requires a non-zero PostID")
	}
	return &models.Comment{
		PostID: b.postID,
		ID:     b.id,
		Name:   b.name,
		Email:  b.email,
		Body:   b.body,
	}, nil
}
