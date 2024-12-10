package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
)

func thereShouldBeUsersInTheResponseBody(expectedUserCount int) error {
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

	var users []models.User
	if err := json.Unmarshal(body, &users); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(users) != expectedUserCount {
		return fmt.Errorf("expected %d albums, but got %d", expectedUserCount, len(users))
	}

	return nil
}

// InitializeUserTestSuite initializes the user test suite
func InitializeUserTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This code will run before the test suite starts
	})

	ctx.AfterSuite(func() {
		// This code will run after the test suite finishes
	})
}

type userFeature struct {
	users []*models.User
}

// InitializeUserScenario initializes the user scenario
func InitializeUserScenario(ctx *godog.ScenarioContext) {
	userFeature := &userFeature{
		users: []*models.User{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		userFeature.users = []*models.User{}
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeUserSteps(ctx)
}

// InitializeUserSteps initializes the user steps
func InitializeUserSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) users in the response body$`, thereShouldBeUsersInTheResponseBody)
}
