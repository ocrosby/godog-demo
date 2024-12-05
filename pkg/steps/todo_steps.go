package steps

import "github.com/cucumber/godog"

// InitializeTodoTestSuite initializes the todo test suite
func InitializeTodoTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This code will run before the test suite starts
	})

	ctx.AfterSuite(func() {
		// This code will run after the test suite finishes
	})
}

// InitializeTodoScenario initializes the todo scenario
func InitializeTodoScenario(ctx *godog.ScenarioContext) {
}
