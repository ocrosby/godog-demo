package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
)

// thereShouldBeTodosInTheResponseBody asserts that the most recent API response
// body contains exactly expected Todo objects. It delegates to the shared
// assertResponseBodyCount generic helper.
func thereShouldBeTodosInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.Todo](expected)
}

// InitializeTodoTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for todo scenarios.
func InitializeTodoTestSuite(_ *godog.TestSuiteContext) {}

// InitializeTodoScenario registers the common HTTP steps and the todo-specific
// step definition for the todo feature scenarios.
func InitializeTodoScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeTodoSteps(ctx)
}

// InitializeTodoSteps registers the todo-specific step definition.
// It is separated from InitializeTodoScenario so it can be reused by other
// scenario contexts that need todo assertions without the full scenario setup.
func InitializeTodoSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) todos in the response body$`, thereShouldBeTodosInTheResponseBody)
}
