package steps

import (
	"github.com/cucumber/godog"
)

func PostSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^a post with title "([^"]*)"$`, aPostWithTitle)
}

func aPostWithTitle(title string) error {
	// Implement the step definition logic here
	return nil
}
