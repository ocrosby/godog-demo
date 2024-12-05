package main

import (
	"context"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/steps"
)

func InitializePostScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// Initialize any necessary state here
		return ctx, nil
	})

	// Register post steps
	steps.InitializePostSteps(ctx)
}
