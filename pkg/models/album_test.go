package models

import (
	"encoding/json"
	"testing"
)

func TestAlbum_MarshalJSON(t *testing.T) {
	a := Album{UserID: 1, ID: 42, Title: "Test Album"}
	b, err := json.Marshal(a)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("re-unmarshal: %v", err)
	}

	if got["userId"].(float64) != 1 {
		t.Errorf("userId = %v, want 1", got["userId"])
	}
	if got["id"].(float64) != 42 {
		t.Errorf("id = %v, want 42", got["id"])
	}
	if got["title"].(string) != "Test Album" {
		t.Errorf("title = %v, want %q", got["title"], "Test Album")
	}
}

func TestAlbum_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Album
		wantErr bool
	}{
		{
			name:  "full object",
			input: `{"userId":1,"id":42,"title":"Test Album"}`,
			want:  Album{UserID: 1, ID: 42, Title: "Test Album"},
		},
		{
			name:  "missing fields default to zero values",
			input: `{}`,
			want:  Album{},
		},
		{
			name:  "extra fields ignored",
			input: `{"userId":1,"id":42,"title":"Test","extra":"ignored"}`,
			want:  Album{UserID: 1, ID: 42, Title: "Test"},
		},
		{
			name:  "unicode title",
			input: `{"userId":1,"id":1,"title":"こんにちは"}`,
			want:  Album{UserID: 1, ID: 1, Title: "こんにちは"},
		},
		{
			name:  "zero ids",
			input: `{"userId":0,"id":0,"title":""}`,
			want:  Album{},
		},
		{
			name:    "invalid JSON",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Album
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

func TestAlbum_RoundTrip(t *testing.T) {
	original := Album{UserID: 5, ID: 99, Title: "Round Trip Album"}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored Album
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != original {
		t.Errorf("round-trip: got %+v, want %+v", restored, original)
	}
}
