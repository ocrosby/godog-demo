package helpers

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// --- ResolveUrl ---

func TestResolveUrl(t *testing.T) {
	tests := []struct {
		name     string
		resource string
		want     string
	}{
		{"root path", "/", "https://jsonplaceholder.typicode.com/"},
		{"resource path", "/albums", "https://jsonplaceholder.typicode.com/albums"},
		{"nested path", "/albums/1", "https://jsonplaceholder.typicode.com/albums/1"},
		{"path with query string", "/albums?userId=1", "https://jsonplaceholder.typicode.com/albums?userId=1"},
		{"empty resource", "", "https://jsonplaceholder.typicode.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ResolveUrl(tt.resource)
			if got != tt.want {
				t.Errorf("ResolveUrl(%q) = %q, want %q", tt.resource, got, tt.want)
			}
		})
	}
}

// --- SendRequest ---

func TestSendRequest_Success(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		responseStatus int
		responseBody   string
	}{
		{"GET 200", http.MethodGet, http.StatusOK, `{"id":1}`},
		{"POST 201", http.MethodPost, http.StatusCreated, `{"id":101}`},
		{"DELETE 200", http.MethodDelete, http.StatusOK, `{}`},
		{"PUT 200", http.MethodPut, http.StatusOK, `{"id":1,"title":"updated"}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tt.method {
					t.Errorf("expected method %s, got %s", tt.method, r.Method)
				}
				w.WriteHeader(tt.responseStatus)
				fmt.Fprint(w, tt.responseBody)
			}))
			defer srv.Close()

			resp, err := SendRequest(tt.method, srv.URL, nil)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.responseStatus {
				t.Errorf("expected status %d, got %d", tt.responseStatus, resp.StatusCode)
			}
		})
	}
}

func TestSendRequest_WithBody(t *testing.T) {
	payload := []byte(`{"title":"new album","userId":1}`)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("server failed to read body: %v", err)
		}
		if string(body) != string(payload) {
			t.Errorf("server received body %q, want %q", body, payload)
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, `{"id":101}`)
	}))
	defer srv.Close()

	resp, err := SendRequest(http.MethodPost, srv.URL, payload)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected status 201, got %d", resp.StatusCode)
	}
}

func TestSendRequest_InvalidMethod(t *testing.T) {
	// A method containing a space is invalid per RFC 7230 and causes http.NewRequest to fail.
	_, err := SendRequest("INVALID METHOD", "http://localhost", nil)
	if err == nil {
		t.Error("expected error for invalid HTTP method, got nil")
	}
}

func TestSendRequest_InvalidURL(t *testing.T) {
	_, err := SendRequest(http.MethodGet, "://bad-url", nil)
	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
}

func TestSendRequest_NetworkError(t *testing.T) {
	// Use a local address on a port that is not listening.
	_, err := SendRequest(http.MethodGet, "http://127.0.0.1:1", nil)
	if err == nil {
		t.Error("expected network error for unreachable address, got nil")
	}
}

// --- ReadResponseBody ---

func TestReadResponseBody_Success(t *testing.T) {
	tests := []struct {
		name string
		body string
	}{
		{"JSON body", `{"id":1,"title":"test"}`},
		{"empty body", ""},
		{"plain text", "hello world"},
		{"unicode body", "こんにちは"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Body: io.NopCloser(strings.NewReader(tt.body)),
			}
			got, err := ReadResponseBody(resp)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.body {
				t.Errorf("ReadResponseBody = %q, want %q", got, tt.body)
			}
		})
	}
}

func TestReadResponseBody_DoesNotCloseBody(t *testing.T) {
	// ReadResponseBody must not close the body — the caller owns it.
	closed := false
	body := &trackingCloser{
		Reader: strings.NewReader(`{"id":1}`),
		onClose: func() {
			closed = true
		},
	}
	resp := &http.Response{Body: body}

	_, err := ReadResponseBody(resp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if closed {
		t.Error("ReadResponseBody closed the response body — caller owns the body, not the callee")
	}
}

// --- HandlePostResponse ---

func TestHandlePostResponse_Success(t *testing.T) {
	type target struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	tests := []struct {
		name     string
		jsonBody string
		wantID   int
		wantErr  bool
	}{
		{"standard POST response", `{"id":101,"title":"new album"}`, 101, false},
		{"id of 1", `{"id":1}`, 1, false},
		{"id of 0", `{"id":0}`, 0, false},
		{"large id", `{"id":9999}`, 9999, false},
		{"missing id field", `{"title":"no id here"}`, 0, true},
		{"invalid JSON", `not json`, 0, true},
		{"empty object", `{}`, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := &http.Response{
				Body: io.NopCloser(strings.NewReader(tt.jsonBody)),
			}
			var out target
			gotID, err := HandlePostResponse(resp, &out)

			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil (id=%d)", gotID)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotID != tt.wantID {
				t.Errorf("returned id = %d, want %d", gotID, tt.wantID)
			}
			if out.ID != tt.wantID {
				t.Errorf("unmarshalled id = %d, want %d", out.ID, tt.wantID)
			}
		})
	}
}

func TestHandlePostResponse_ClosesBody(t *testing.T) {
	closed := false
	body := &trackingCloser{
		Reader: strings.NewReader(`{"id":1}`),
		onClose: func() {
			closed = true
		},
	}
	resp := &http.Response{Body: body}

	var out struct{ ID int `json:"id"` }
	if _, err := HandlePostResponse(resp, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !closed {
		t.Error("HandlePostResponse did not close the response body")
	}
}

func TestHandlePostResponse_UnmarshalsTarget(t *testing.T) {
	type album struct {
		UserID int    `json:"userId"`
		ID     int    `json:"id"`
		Title  string `json:"title"`
	}

	body := `{"userId":1,"id":101,"title":"Test Album"}`
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(body)),
	}

	var out album
	id, err := HandlePostResponse(resp, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != 101 {
		t.Errorf("id = %d, want 101", id)
	}
	if out.UserID != 1 {
		t.Errorf("UserID = %d, want 1", out.UserID)
	}
	if out.Title != "Test Album" {
		t.Errorf("Title = %q, want %q", out.Title, "Test Album")
	}
}

// --- helpers ---

// trackingCloser wraps a Reader and records whether Close was called.
type trackingCloser struct {
	io.Reader
	onClose func()
}

func (t *trackingCloser) Close() error {
	t.onClose()
	return nil
}

