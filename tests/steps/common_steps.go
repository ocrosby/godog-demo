package steps

import (
	"github.com/cucumber/godog"
)

func CommonSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^the response status code should be (\d+)$`, responseStatusCodeShouldBe)
}

func responseStatusCodeShouldBe(expectedStatusCode int) error {
	// Implement the step definition logic here
	return nil
}
