package steps

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
)

// thereShouldBePhotosInTheResponseBody asserts that the most recent API response
// body contains exactly expected Photo objects. It delegates to the shared
// assertResponseBodyCount generic helper.
func thereShouldBePhotosInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.Photo](expected)
}

// InitializePhotoTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for photo scenarios.
func InitializePhotoTestSuite(_ *godog.TestSuiteContext) {}

// InitializePhotoScenario registers the common HTTP steps and the photo-specific
// step definition for the photo feature scenarios.
func InitializePhotoScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializePhotoSteps(ctx)
}

// InitializePhotoSteps registers the photo-specific step definition.
// It is separated from InitializePhotoScenario so it can be reused by other
// scenario contexts that need photo assertions without the full scenario setup.
func InitializePhotoSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) photos in the response body$`, thereShouldBePhotosInTheResponseBody)
}
