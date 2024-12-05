package steps

import (
	"context"
	"github.com/cucumber/godog"
	"strconv"
)

var lastError error

// iDeleteAPostWithId deletes a post with the given ID
func iDeleteAPostWithId(postId int) error {
	resource := "/posts/" + strconv.Itoa(postId)

	lastError = iSendRequestTo("DELETE", resource)

	return lastError
}

func thereShouldBeNoErrors() error {
	if lastError != nil {
		return lastError
	}

	return nil
}

func InitializePostScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Initialize any necessary state here
		lastError = nil
		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		// Clean up any state here
		return ctx, nil
	})

	// Register post steps
	ctx.Step(`^I delete a post with id (\d+)$`, iDeleteAPostWithId)
	ctx.Step(`^the response should be successful$`, ResponseShouldBeSuccessful)
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
}

func InitializePostTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This code will run before the test suite starts
	})

	ctx.AfterSuite(func() {
		// This code will run after the test suite finishes
	})
}
