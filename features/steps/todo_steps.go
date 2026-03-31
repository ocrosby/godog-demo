package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
)

func thereShouldBeTodosInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.Todo](expected)
}

func InitializeTodoTestSuite(_ *godog.TestSuiteContext) {}

func InitializeTodoScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeTodoSteps(ctx)
}

func InitializeTodoSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) todos in the response body$`, thereShouldBeTodosInTheResponseBody)
}
