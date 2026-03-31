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

// contextKey is a private type for context keys in this package, preventing
// collisions with keys from other packages.
type contextKey string

const errorKey contextKey = "error"

var lastResponse *http.Response

func withError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey, err)
}

func getError(ctx context.Context) error {
	if err, ok := ctx.Value(errorKey).(error); ok {
		return err
	}
	return nil
}

// assertResponseBodyCount reads lastResponse.Body, unmarshals it into a []T,
// and returns an error when the count does not match expected.
// Applying rule-of-three: photo, todo, and user steps all share this pattern.
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

func iSendRequestTo(method, resource string) error {
	var err error
	lastResponse, err = helpers.SendRequest(method, helpers.ResolveUrl(resource), nil)
	return err
}

func responseShouldBeSuccessful() error {
	if lastResponse.StatusCode < http.StatusOK || lastResponse.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("expected successful status code, got %d", lastResponse.StatusCode)
	}
	return nil
}

func responseStatusCodeShouldBe(expected int) error {
	if lastResponse.StatusCode != expected {
		return fmt.Errorf("expected status code %d, got %d", expected, lastResponse.StatusCode)
	}
	return nil
}

func thereShouldBeNoErrors(ctx context.Context) error {
	if err := getError(ctx); err != nil {
		return fmt.Errorf("expected no errors, got %v", err)
	}
	return nil
}

func InitializeCommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response should be successful$`, responseShouldBeSuccessful)
	ctx.Step(`^the response status code should be (\d+)$`, responseStatusCodeShouldBe)
	ctx.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
}
