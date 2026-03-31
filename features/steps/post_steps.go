package steps

import (
	"fmt"

	"github.com/cucumber/godog"
)

func iDeleteAPostWithId(postId int) error {
	return iSendRequestTo("DELETE", fmt.Sprintf("/posts/%d", postId))
}

func InitializePostTestSuite(_ *godog.TestSuiteContext) {}

func InitializePostScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I delete a post with id (\d+)$`, iDeleteAPostWithId)
	ctx.Step(`^the response should be successful$`, responseShouldBeSuccessful)
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
}
