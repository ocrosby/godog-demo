package models

import (
	"encoding/json"
	"testing"
)

func TestTodo_MarshalJSON(t *testing.T) {
	tests := []struct {
		name          string
		todo          Todo
		wantCompleted bool
	}{
		{"completed true", Todo{UserID: 1, ID: 1, Title: "Buy milk", Completed: true}, true},
		{"completed false", Todo{UserID: 1, ID: 2, Title: "Buy eggs", Completed: false}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(tt.todo)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			var got map[string]interface{}
			if err := json.Unmarshal(b, &got); err != nil {
				t.Fatalf("re-unmarshal: %v", err)
			}
			if got["completed"].(bool) != tt.wantCompleted {
				t.Errorf("completed = %v, want %v", got["completed"], tt.wantCompleted)
			}
		})
	}
}

func TestTodo_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Todo
		wantErr bool
	}{
		{
			name:  "completed true",
			input: `{"userId":1,"id":1,"title":"Buy milk","completed":true}`,
			want:  Todo{UserID: 1, ID: 1, Title: "Buy milk", Completed: true},
		},
		{
			name:  "completed false",
			input: `{"userId":1,"id":2,"title":"Buy eggs","completed":false}`,
			want:  Todo{UserID: 1, ID: 2, Title: "Buy eggs", Completed: false},
		},
		{
			name:  "missing fields default to zero values",
			input: `{}`,
			want:  Todo{},
		},
		{
			name:  "extra fields ignored",
			input: `{"userId":1,"id":1,"title":"T","completed":false,"extra":"ignored"}`,
			want:  Todo{UserID: 1, ID: 1, Title: "T", Completed: false},
		},
		{
			name:  "unicode title",
			input: `{"userId":1,"id":1,"title":"やること","completed":true}`,
			want:  Todo{UserID: 1, ID: 1, Title: "やること", Completed: true},
		},
		{
			name:    "invalid JSON",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Todo
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

func TestTodo_RoundTrip(t *testing.T) {
	tests := []Todo{
		{UserID: 1, ID: 1, Title: "Task A", Completed: true},
		{UserID: 2, ID: 2, Title: "Task B", Completed: false},
	}
	for _, original := range tests {
		b, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		var restored Todo
		if err := json.Unmarshal(b, &restored); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if restored != original {
			t.Errorf("round-trip: got %+v, want %+v", restored, original)
		}
	}
}
