package builders_test

import (
	"testing"

	"github.com/ocrosby/godog-demo/pkg/builders"
)

func TestAlbumBuilder_Build_RequiresUserID(t *testing.T) {
	_, err := builders.NewAlbumBuilder().Build()
	if err == nil {
		t.Fatal("expected error when UserID is zero, got nil")
	}
}

func TestAlbumBuilder_Build_Success(t *testing.T) {
	album, err := builders.NewAlbumBuilder().
		WithUserID(1).
		WithID(101).
		WithTitle("My Album").
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if album.UserID != 1 {
		t.Errorf("UserID: got %d, want 1", album.UserID)
	}
	if album.ID != 101 {
		t.Errorf("ID: got %d, want 101", album.ID)
	}
	if album.Title != "My Album" {
		t.Errorf("Title: got %q, want %q", album.Title, "My Album")
	}
}

func TestAlbumBuilder_Build_OptionalFields(t *testing.T) {
	// Only UserID is required; ID and Title have zero-value defaults.
	album, err := builders.NewAlbumBuilder().WithUserID(5).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if album.UserID != 5 {
		t.Errorf("UserID: got %d, want 5", album.UserID)
	}
	if album.Title != "" {
		t.Errorf("Title: got %q, want empty string", album.Title)
	}
}

func TestAlbumBuilder_Chaining(t *testing.T) {
	b := builders.NewAlbumBuilder()
	returned := b.WithUserID(1)
	if returned != b {
		t.Error("WithUserID should return the same builder instance for chaining")
	}
}
