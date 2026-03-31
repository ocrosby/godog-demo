// Package helpers provides HTTP utilities for constructing and executing
// requests against the JSONPlaceholder REST API used by the BDD test suite.
package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// baseURL is the root of the JSONPlaceholder API targeted by all scenarios.
const baseURL = "https://jsonplaceholder.typicode.com"

// ResolveUrl prepends the JSONPlaceholder base URL to resource and returns the
// absolute URL string. resource should begin with a slash, e.g. "/albums/1".
func ResolveUrl(resource string) string {
	return baseURL + resource
}

// SendRequest creates an HTTP request with the given method, url, and optional
// body, executes it with a default client, and returns the response.
//
// body may be nil for requests that carry no payload (GET, DELETE).
// The caller is responsible for closing the returned response body.
// It returns a non-nil error if the request cannot be constructed or the
// network call fails; HTTP error status codes are not treated as errors here.
func SendRequest(method, url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ReadResponseBody reads the entire body of resp into a string and returns it.
// It does NOT close the response body; that responsibility belongs to the caller.
// It returns an error if the underlying read fails.
func ReadResponseBody(resp *http.Response) (string, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// HandlePostResponse reads and closes response.Body, unmarshals the JSON payload
// into target, and extracts the integer "id" field from the response.
//
// It is intended for POST endpoints that echo the created resource back with a
// server-assigned ID (e.g. JSONPlaceholder's POST /albums).
// It returns the extracted ID and a nil error on success, or (0, err) when:
//   - the body cannot be read,
//   - the JSON cannot be unmarshalled into target,
//   - the response does not contain a non-zero "id" field.
func HandlePostResponse(response *http.Response, target interface{}) (int, error) {
	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return 0, fmt.Errorf("failed to read response body: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return 0, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	// Extract the server-assigned id using a typed struct to avoid the
	// map[string]interface{} + float64 assertion that the previous implementation used.
	var idHolder struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(body, &idHolder); err != nil {
		return 0, fmt.Errorf("extracting id from response: %w", err)
	}
	if idHolder.ID == 0 {
		return 0, fmt.Errorf("response does not contain an id property")
	}

	return idHolder.ID, nil
}
