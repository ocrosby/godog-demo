package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
)

// thereShouldBeUsersInTheResponseBody asserts that the most recent API response
// body contains exactly expected User objects. It delegates to the shared
// assertResponseBodyCount generic helper.
func thereShouldBeUsersInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.User](expected)
}

// InitializeUserTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for user scenarios.
func InitializeUserTestSuite(_ *godog.TestSuiteContext) {}

// InitializeUserScenario registers the common HTTP steps and the user-specific
// step definition for the user feature scenarios.
func InitializeUserScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeUserSteps(ctx)
}

// InitializeUserSteps registers the user-specific step definition.
// It is separated from InitializeUserScenario so it can be reused by other
// scenario contexts that need user assertions without the full scenario setup.
func InitializeUserSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) users in the response body$`, thereShouldBeUsersInTheResponseBody)
}
