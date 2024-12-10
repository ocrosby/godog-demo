package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
)

func thereShouldBeTodosInTheResponseBody(expectedTodoCount int) error {
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

	var todos []models.Todo
	if err := json.Unmarshal(body, &todos); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(todos) != expectedTodoCount {
		return fmt.Errorf("expected %d todos, but got %d", expectedTodoCount, len(todos))
	}

	return nil
}

type todoFeature struct {
	todos []*models.Todo
}

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
	todoFeature := &todoFeature{
		todos: []*models.Todo{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		todoFeature.todos = []*models.Todo{}
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeTodoSteps(ctx)
}

func InitializeTodoSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) todos in the response body$`, thereShouldBeTodosInTheResponseBody)
}
