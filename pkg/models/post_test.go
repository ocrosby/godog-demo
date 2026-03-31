package models

import (
	"encoding/json"
	"testing"
)

func TestPost_MarshalJSON(t *testing.T) {
	p := Post{UserID: 1, ID: 10, Title: "Hello World", Body: "Post body text."}
	b, err := json.Marshal(p)
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
	if got["id"].(float64) != 10 {
		t.Errorf("id = %v, want 10", got["id"])
	}
	if got["title"].(string) != "Hello World" {
		t.Errorf("title = %v, want 'Hello World'", got["title"])
	}
	if got["body"].(string) != "Post body text." {
		t.Errorf("body = %v, want 'Post body text.'", got["body"])
	}
}

func TestPost_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Post
		wantErr bool
	}{
		{
			name:  "full object",
			input: `{"userId":1,"id":10,"title":"Hello World","body":"Post body text."}`,
			want:  Post{UserID: 1, ID: 10, Title: "Hello World", Body: "Post body text."},
		},
		{
			name:  "missing fields default to zero values",
			input: `{}`,
			want:  Post{},
		},
		{
			name:  "extra fields ignored",
			input: `{"userId":1,"id":1,"title":"T","body":"B","extra":"ignored"}`,
			want:  Post{UserID: 1, ID: 1, Title: "T", Body: "B"},
		},
		{
			name:  "unicode in title and body",
			input: `{"userId":1,"id":1,"title":"タイトル","body":"本文"}`,
			want:  Post{UserID: 1, ID: 1, Title: "タイトル", Body: "本文"},
		},
		{
			name:  "empty strings",
			input: `{"userId":0,"id":0,"title":"","body":""}`,
			want:  Post{},
		},
		{
			name:    "invalid JSON",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Post
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

func TestPost_RoundTrip(t *testing.T) {
	original := Post{UserID: 2, ID: 20, Title: "Round Trip", Body: "body content"}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored Post
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != original {
		t.Errorf("round-trip: got %+v, want %+v", restored, original)
	}
}
