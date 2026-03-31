package models

import (
	"encoding/json"
	"testing"
)

func TestComment_MarshalJSON(t *testing.T) {
	c := Comment{PostID: 1, ID: 501, Name: "Alice", Email: "alice@example.com", Body: "Great post!"}
	b, err := json.Marshal(c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("re-unmarshal: %v", err)
	}

	if got["postId"].(float64) != 1 {
		t.Errorf("postId = %v, want 1", got["postId"])
	}
	if got["id"].(float64) != 501 {
		t.Errorf("id = %v, want 501", got["id"])
	}
	if got["name"].(string) != "Alice" {
		t.Errorf("name = %v, want Alice", got["name"])
	}
	if got["email"].(string) != "alice@example.com" {
		t.Errorf("email = %v, want alice@example.com", got["email"])
	}
	if got["body"].(string) != "Great post!" {
		t.Errorf("body = %v, want 'Great post!'", got["body"])
	}
}

func TestComment_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Comment
		wantErr bool
	}{
		{
			name:  "full object",
			input: `{"postId":1,"id":501,"name":"Alice","email":"alice@example.com","body":"Great post!"}`,
			want:  Comment{PostID: 1, ID: 501, Name: "Alice", Email: "alice@example.com", Body: "Great post!"},
		},
		{
			name:  "missing fields default to zero values",
			input: `{}`,
			want:  Comment{},
		},
		{
			name:  "extra fields ignored",
			input: `{"postId":1,"id":1,"name":"Bob","email":"b@b.com","body":"hi","extra":"ignored"}`,
			want:  Comment{PostID: 1, ID: 1, Name: "Bob", Email: "b@b.com", Body: "hi"},
		},
		{
			name:  "unicode in name and body",
			input: `{"postId":2,"id":10,"name":"山田太郎","email":"t@t.com","body":"日本語のコメント"}`,
			want:  Comment{PostID: 2, ID: 10, Name: "山田太郎", Email: "t@t.com", Body: "日本語のコメント"},
		},
		{
			name:    "invalid JSON",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Comment
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

func TestComment_RoundTrip(t *testing.T) {
	original := Comment{PostID: 3, ID: 55, Name: "Bob", Email: "bob@example.com", Body: "Nice"}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored Comment
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != original {
		t.Errorf("round-trip: got %+v, want %+v", restored, original)
	}
}
