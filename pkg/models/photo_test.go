package models

import (
	"encoding/json"
	"testing"
)

func TestPhoto_MarshalJSON(t *testing.T) {
	p := Photo{
		AlbumID:      2,
		ID:           1,
		Title:        "Sunset",
		URL:          "https://example.com/photo.jpg",
		ThumbnailURL: "https://example.com/thumb.jpg",
	}
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("re-unmarshal: %v", err)
	}

	if got["albumId"].(float64) != 2 {
		t.Errorf("albumId = %v, want 2", got["albumId"])
	}
	if got["id"].(float64) != 1 {
		t.Errorf("id = %v, want 1", got["id"])
	}
	if got["title"].(string) != "Sunset" {
		t.Errorf("title = %v, want Sunset", got["title"])
	}
	if got["url"].(string) != "https://example.com/photo.jpg" {
		t.Errorf("url = %v, want https://example.com/photo.jpg", got["url"])
	}
	if got["thumbnailUrl"].(string) != "https://example.com/thumb.jpg" {
		t.Errorf("thumbnailUrl = %v, want https://example.com/thumb.jpg", got["thumbnailUrl"])
	}
}

func TestPhoto_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Photo
		wantErr bool
	}{
		{
			name:  "full object",
			input: `{"albumId":2,"id":1,"title":"Sunset","url":"https://example.com/photo.jpg","thumbnailUrl":"https://example.com/thumb.jpg"}`,
			want: Photo{
				AlbumID:      2,
				ID:           1,
				Title:        "Sunset",
				URL:          "https://example.com/photo.jpg",
				ThumbnailURL: "https://example.com/thumb.jpg",
			},
		},
		{
			name:  "missing fields default to zero values",
			input: `{}`,
			want:  Photo{},
		},
		{
			name:  "extra fields ignored",
			input: `{"albumId":1,"id":1,"title":"x","url":"u","thumbnailUrl":"t","extra":"ignored"}`,
			want:  Photo{AlbumID: 1, ID: 1, Title: "x", URL: "u", ThumbnailURL: "t"},
		},
		{
			name:  "empty strings for URL fields",
			input: `{"albumId":1,"id":1,"title":"no urls","url":"","thumbnailUrl":""}`,
			want:  Photo{AlbumID: 1, ID: 1, Title: "no urls", URL: "", ThumbnailURL: ""},
		},
		{
			name:    "invalid JSON",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Photo
			err := json.Unmarshal([]byte(tt.input), &got)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestPhoto_RoundTrip(t *testing.T) {
	original := Photo{AlbumID: 3, ID: 7, Title: "Landscape", URL: "https://u.com/p.jpg", ThumbnailURL: "https://u.com/t.jpg"}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored Photo
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != original {
		t.Errorf("round-trip: got %+v, want %+v", restored, original)
	}
}
