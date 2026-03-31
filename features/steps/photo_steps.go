package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
)

func thereShouldBePhotosInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.Photo](expected)
}

func InitializePhotoTestSuite(_ *godog.TestSuiteContext) {}

func InitializePhotoScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializePhotoSteps(ctx)
}

func InitializePhotoSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) photos in the response body$`, thereShouldBePhotosInTheResponseBody)
}
