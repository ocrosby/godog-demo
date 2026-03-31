// Package steps contains the GoDog step definitions and shared helpers for all
// BDD feature scenarios in this project. Each domain (albums, comments, photos,
// posts, todos, users, calculator) has its own file; this file holds the
// cross-cutting infrastructure that every domain depends on.
package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/helpers"
)

// contextKey is an unexported string type used as the key type for values stored
// in a context.Context. Using a named type instead of a raw string prevents
// accidental collisions with keys defined by other packages.
type contextKey string

// errorKey is the context key under which a step-level error is stored so that
// subsequent steps in the same scenario can inspect it (e.g. "the result should
// be an error").
const errorKey contextKey = "error"

// lastResponse holds the most recent HTTP response received by iSendRequestTo.
// It is package-level so that assertion steps registered through InitializeCommonSteps
// can access it without passing it through the scenario context.
var lastResponse *http.Response

// withError stores err in ctx under errorKey and returns the updated context.
// It is used by operation steps to signal that an expected error occurred so
// that a later "the result should be an error" step can verify it.
func withError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey, err)
}

// getError retrieves the error stored by withError from ctx.
// It returns nil if no error has been stored or the stored value is not an error.
func getError(ctx context.Context) error {
	if err, ok := ctx.Value(errorKey).(error); ok {
		return err
	}
	return nil
}

// assertResponseBodyCount reads lastResponse.Body, unmarshals it into a []T,
// and returns an error when the item count does not equal expected.
//
// This generic helper eliminates the read/unmarshal/count pattern that would
// otherwise be duplicated across every resource-listing step (photos, todos,
// users, etc.). It closes the response body on return.
func assertResponseBodyCount[T any](expected int) error {
	defer lastResponse.Body.Close()

	body, err := io.ReadAll(lastResponse.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	var items []T
	if err := json.Unmarshal(body, &items); err != nil {
		return fmt.Errorf("unmarshalling response body: %w", err)
	}

	if len(items) != expected {
		return fmt.Errorf("expected %d items in response body, got %d", expected, len(items))
	}

	return nil
}

// iSendRequestTo sends an HTTP request using method (e.g. "GET") to the
// JSONPlaceholder resource path (e.g. "/posts/1") and stores the response in
// lastResponse for subsequent assertion steps.
// It returns any transport-level error; HTTP error status codes are not treated
// as errors here.
func iSendRequestTo(method, resource string) error {
	var err error
	lastResponse, err = helpers.SendRequest(method, helpers.ResolveUrl(resource), nil)
	return err
}

// responseShouldBeSuccessful returns an error when lastResponse carries a
// non-2xx status code, indicating an unexpected failure from the API.
func responseShouldBeSuccessful() error {
	if lastResponse.StatusCode < http.StatusOK || lastResponse.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("expected successful status code, got %d", lastResponse.StatusCode)
	}
	return nil
}

// responseStatusCodeShouldBe returns an error when lastResponse does not carry
// exactly the expected HTTP status code.
func responseStatusCodeShouldBe(expected int) error {
	if lastResponse.StatusCode != expected {
		return fmt.Errorf("expected status code %d, got %d", expected, lastResponse.StatusCode)
	}
	return nil
}

// thereShouldBeNoErrors returns an error if a step error has been stored in ctx
// via withError, allowing scenarios to assert a clean execution path.
func thereShouldBeNoErrors(ctx context.Context) error {
	if err := getError(ctx); err != nil {
		return fmt.Errorf("expected no errors, got %v", err)
	}
	return nil
}

// InitializeCommonSteps registers the HTTP request and response assertion steps
// that are shared across all domain scenario contexts.
func InitializeCommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response should be successful$`, responseShouldBeSuccessful)
	ctx.Step(`^the response status code should be (\d+)$`, responseStatusCodeShouldBe)
	ctx.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
}
