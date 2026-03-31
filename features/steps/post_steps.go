package steps

import (
	"fmt"

	"github.com/cucumber/godog"
)

// iDeleteAPostWithId sends DELETE /posts/{postId} using the shared iSendRequestTo
// helper and returns any transport-level error.
func iDeleteAPostWithId(postId int) error {
	return iSendRequestTo("DELETE", fmt.Sprintf("/posts/%d", postId))
}

// InitializePostTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for post scenarios.
func InitializePostTestSuite(_ *godog.TestSuiteContext) {}

// InitializePostScenario wires the post step definitions and the shared
// response-assertion steps to their Gherkin patterns.
func InitializePostScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I delete a post with id (\d+)$`, iDeleteAPostWithId)
	ctx.Step(`^the response should be successful$`, responseShouldBeSuccessful)
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
}
