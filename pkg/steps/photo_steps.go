package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
)

func AndThereShouldBePhotosInTheResponseBody(expectedPhotoCount int) error {
	body, err := io.ReadAll(lastResponse.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close response body:", err)
		}
	}(lastResponse.Body)

	var photos []models.Photo
	if err := json.Unmarshal(body, &photos); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(photos) != expectedPhotoCount {
		return fmt.Errorf("expected %d albums, but got %d", expectedPhotoCount, len(photos))
	}

	return nil
}

type photoFeature struct {
	photos []*models.Photo
}

func InitializePhotoTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This code will run before the test suite starts
	})

	ctx.AfterSuite(func() {
		// This code will run after the test suite finishes
	})
}

func InitializePhotoScenario(ctx *godog.ScenarioContext) {
	photoFeature := &photoFeature{
		photos: []*models.Photo{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		photoFeature.photos = []*models.Photo{}
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializePhotoSteps(ctx)
}

func InitializePhotoSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) photos in the response body$`, AndThereShouldBePhotosInTheResponseBody)
}
