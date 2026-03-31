// Package models defines the domain types that mirror the JSONPlaceholder API
// resources used by the BDD acceptance-test suite.
package models

// Album represents a photo album as returned by the JSONPlaceholder /albums
// endpoint. All fields are serialised using the camelCase JSON keys expected
// by the API.
type Album struct {
	// UserID is the identifier of the user who owns the album.
	UserID int `json:"userId"`

	// ID is the unique identifier of the album.
	ID int `json:"id"`

	// Title is the human-readable name of the album.
	Title string `json:"title"`
}
