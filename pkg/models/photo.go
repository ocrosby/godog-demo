package models

// Photo represents a photo as returned by the JSONPlaceholder /photos endpoint.
type Photo struct {
	// AlbumID is the identifier of the album this photo belongs to.
	AlbumID int `json:"albumId"`

	// ID is the unique identifier of the photo.
	ID int `json:"id"`

	// Title is the human-readable name of the photo.
	Title string `json:"title"`

	// URL is the absolute URL of the full-size photo.
	URL string `json:"url"`

	// ThumbnailURL is the absolute URL of the thumbnail version of the photo.
	ThumbnailURL string `json:"thumbnailUrl"`
}
