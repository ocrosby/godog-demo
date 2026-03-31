package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
)

func thereShouldBeUsersInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.User](expected)
}

func InitializeUserTestSuite(_ *godog.TestSuiteContext) {}

func InitializeUserScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeUserSteps(ctx)
}

func InitializeUserSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) users in the response body$`, thereShouldBeUsersInTheResponseBody)
}
