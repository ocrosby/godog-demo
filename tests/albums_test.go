package main

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"github.com/ocrosby/godog-demo/pkg/steps"
	"testing"
)

type albumFeature struct {
	albums []*models.Album
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	albumFeature := &albumFeature{
		albums: []*models.Album{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		albumFeature.albums = []*models.Album{}
		return ctx, nil
	})

	steps.InitializeCommonSteps(ctx)
	steps.InitializeAlbumSteps(ctx)
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format: "pretty",
			Paths:  []string{"../features"},
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
