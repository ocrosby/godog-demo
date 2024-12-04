package main

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/models"
)

type albumFeature struct {
	albums []*models.Album
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	calcFeature := &calculatorFeature{
		calc: &Calculator{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		calcFeature.calc.Clear()
		return ctx, nil
	})

	ctx.Step(`^calculator is cleared$`, calcFeature.calculatorIsCleared)
	ctx.Step(`^I add (\d+)$`, calcFeature.iAdd)
	ctx.Step(`^I press (\d+)$`, calcFeature.iPress)
	ctx.Step(`^I subtract (\d+)$`, calcFeature.iSubtract)
	ctx.Step(`^the result should be (\d+)$`, calcFeature.theResultShouldBe)
}
