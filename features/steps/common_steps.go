package steps

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/helpers"
	"net/http"
)

var lastResponse *http.Response

const errorKey contextKey = "error"

func withError(ctx context.Context, err error) context.Context {
	return context.WithValue(ctx, errorKey, err)
}

func getError(ctx context.Context) error {
	if err, ok := ctx.Value(errorKey).(error); ok {
		return err
	}
	
	return nil
}

// iSendRequestTo sends a request to the specified resource
func iSendRequestTo(method, resource string) error {
	var err error

	url := "https://jsonplaceholder.typicode.com" + resource

	lastResponse, err = helpers.SendRequest(method, url, nil)

	return err
}

// ResponseShouldBeSuccessful checks if the response is successful
func ResponseShouldBeSuccessful() error {
	if lastResponse.StatusCode < 200 || lastResponse.StatusCode >= 300 {
		return fmt.Errorf("expected status code to be successful, but got %d", lastResponse.StatusCode)
	}

	return nil
}

// responseStatusCodeShouldBe checks if the response status code is as expected
func responseStatusCodeShouldBe(expectedStatusCode int) error {
	if lastResponse.StatusCode != expectedStatusCode {
		return fmt.Errorf("expected status code %d, but got %d", expectedStatusCode, lastResponse.StatusCode)
	}

	return nil
}

func thereShouldBeNoErrors(ctx context.Context) error {
	err := ctx.Value("error")
	if err != nil {
		return fmt.Errorf("expected no errors, but got %v", err)
	}

	return nil
}

// InitializeCommonSteps defines the common steps for the test suite
func InitializeCommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response should be successful$`, ResponseShouldBeSuccessful)
	ctx.Step(`^the response status code should be (\d+)$`, responseStatusCodeShouldBe)
	ctx.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
}
