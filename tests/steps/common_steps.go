package steps

import (
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/helpers"
	"net/http"
)

var lastResponse *http.Response

// CommonSteps defines the common steps for the test suite
func CommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response status code should be (\d+)$`, responseStatusCodeShouldBe)
	ctx.Step(`^I send a "([^"]*)" request to "([^"]*)"$`, iSendRequestTo)
}

// iSendRequestTo sends a request to the specified resource
func iSendRequestTo(method, resource string) error {
	var err error

	lastResponse, err = helpers.SendRequest(method, resource, nil)

	return err
}

// responseStatusCodeShouldBe checks if the response status code is as expected
func responseStatusCodeShouldBe(expectedStatusCode int) error {
	if lastResponse.StatusCode != expectedStatusCode {
		return fmt.Errorf("expected status code %d, but got %d", expectedStatusCode, lastResponse.StatusCode)
	}

	return nil
}
