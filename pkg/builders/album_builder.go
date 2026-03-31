package builders

import (
	"errors"

	"github.com/ocrosby/godog-demo/pkg/models"
)

// AlbumBuilder accumulates field values for an Album before creating it.
// Use NewAlbumBuilder to obtain an instance, call the With* methods to set
// fields, and call Build to obtain a validated *models.Album.
type AlbumBuilder struct {
	id     int
	userID int
	title  string
}

// NewAlbumBuilder returns a new, empty AlbumBuilder.
func NewAlbumBuilder() *AlbumBuilder {
	return &AlbumBuilder{}
}

// WithID sets the ID field and returns the builder for method chaining.
func (b *AlbumBuilder) WithID(id int) *AlbumBuilder {
	b.id = id
	return b
}

// WithUserID sets the UserID field and returns the builder for method chaining.
func (b *AlbumBuilder) WithUserID(userID int) *AlbumBuilder {
	b.userID = userID
	return b
}

// WithTitle sets the Title field and returns the builder for method chaining.
func (b *AlbumBuilder) WithTitle(title string) *AlbumBuilder {
	b.title = title
	return b
}

// Build validates the accumulated state and returns a new *models.Album.
// It returns an error if UserID is zero, as UserID is required by
// JSONPlaceholder's POST /albums endpoint.
func (b *AlbumBuilder) Build() (*models.Album, error) {
	if b.userID == 0 {
		return nil, errors.New("album requires a non-zero UserID")
	}
	return &models.Album{
		ID:     b.id,
		UserID: b.userID,
		Title:  b.title,
	}, nil
}
