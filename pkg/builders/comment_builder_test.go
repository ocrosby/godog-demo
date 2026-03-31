package builders_test

import (
	"testing"

	"github.com/ocrosby/godog-demo/pkg/builders"
)

func TestCommentBuilder_Build_RequiresPostID(t *testing.T) {
	_, err := builders.NewCommentBuilder().Build()
	if err == nil {
		t.Fatal("expected error when PostID is zero, got nil")
	}
}

func TestCommentBuilder_Build_Success(t *testing.T) {
	comment, err := builders.NewCommentBuilder().
		WithPostID(1).
		WithID(501).
		WithName("Alice").
		WithEmail("alice@example.com").
		WithBody("Hello").
		Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if comment.PostID != 1 {
		t.Errorf("PostID: got %d, want 1", comment.PostID)
	}
	if comment.ID != 501 {
		t.Errorf("ID: got %d, want 501", comment.ID)
	}
	if comment.Name != "Alice" {
		t.Errorf("Name: got %q, want %q", comment.Name, "Alice")
	}
	if comment.Email != "alice@example.com" {
		t.Errorf("Email: got %q, want %q", comment.Email, "alice@example.com")
	}
	if comment.Body != "Hello" {
		t.Errorf("Body: got %q, want %q", comment.Body, "Hello")
	}
}

func TestCommentBuilder_Build_OptionalFields(t *testing.T) {
	// Only PostID is required; all other fields have zero-value defaults.
	comment, err := builders.NewCommentBuilder().WithPostID(3).Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if comment.PostID != 3 {
		t.Errorf("PostID: got %d, want 3", comment.PostID)
	}
	if comment.Name != "" {
		t.Errorf("Name: got %q, want empty string", comment.Name)
	}
}

func TestCommentBuilder_Chaining(t *testing.T) {
	b := builders.NewCommentBuilder()
	returned := b.WithPostID(1)
	if returned != b {
		t.Error("WithPostID should return the same builder instance for chaining")
	}
}

func TestCommentBuilder_Build_MultipleCallsIndependent(t *testing.T) {
	b := builders.NewCommentBuilder().WithPostID(1).WithName("Alice")

	c1, err := b.Build()
	if err != nil {
		t.Fatalf("first Build: %v", err)
	}

	b.WithName("Bob")
	c2, err := b.Build()
	if err != nil {
		t.Fatalf("second Build: %v", err)
	}

	if c1.Name == c2.Name {
		t.Error("second Build should reflect updated name")
	}
}
