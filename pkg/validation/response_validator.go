// Package validation provides the ResponseValidator Strategy interface and its
// concrete implementations for asserting HTTP response status codes in BDD step
// definitions.
//
// Using the Strategy pattern here eliminates the duplicate 2xx-check logic that
// previously appeared independently in both common_steps.go and album_steps.go.
// All status-code assertions now share one canonical implementation.
package validation

import (
	"fmt"
	"net/http"
)

// ResponseValidator defines the contract for asserting that an HTTP response
// meets a specific status-code condition. Concrete implementations are chosen
// at the call site, making the validation rule an explicit, swappable strategy.
type ResponseValidator interface {
	Validate(resp *http.Response) error
}

// SuccessValidator is a ResponseValidator that accepts any 2xx HTTP status code.
// It is the canonical implementation of "the response should be successful".
type SuccessValidator struct{}

// Validate returns nil when resp carries a 2xx status code and a descriptive
// error otherwise.
func (SuccessValidator) Validate(resp *http.Response) error {
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("expected successful status code, got %d", resp.StatusCode)
	}
	return nil
}

// FailureValidator is a ResponseValidator that accepts any non-2xx HTTP status
// code. It is the canonical implementation of "the response should be
// unsuccessful".
type FailureValidator struct{}

// Validate returns nil when resp carries a non-2xx status code and a
// descriptive error when resp is unexpectedly successful.
func (FailureValidator) Validate(resp *http.Response) error {
	if resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices {
		return fmt.Errorf("expected unsuccessful status code, got %d", resp.StatusCode)
	}
	return nil
}

// ExactStatusValidator is a ResponseValidator that accepts exactly one HTTP
// status code. Use it when a step must assert a specific code such as 201 or 404.
type ExactStatusValidator struct {
	// Expected is the HTTP status code the response must carry.
	Expected int
}

// Validate returns nil when resp.StatusCode equals v.Expected and a descriptive
// error otherwise.
func (v ExactStatusValidator) Validate(resp *http.Response) error {
	if resp.StatusCode != v.Expected {
		return fmt.Errorf("expected status code %d, got %d", v.Expected, resp.StatusCode)
	}
	return nil
}
