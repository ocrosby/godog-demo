package steps

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
)

type CommentFeature struct {
	comments []*models.Comment
}

func (c *CommentFeature) thereShouldBeCommentsInTheResponseBody(expectedCommentCount int) error {
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

	if err := json.Unmarshal(body, &c.comments); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(c.comments) != expectedCommentCount {
		return fmt.Errorf("expected %d comments, but got %d", expectedCommentCount, len(c.comments))
	}

	return nil
}

func InitializeCommentSteps(ctx *godog.ScenarioContext) {
	commentFeature := &CommentFeature{
		comments: []*models.Comment{},
	}

	ctx.Step(`^there should be (\d+) comments in the response body$`, commentFeature.thereShouldBeCommentsInTheResponseBody)
}

func InitializeCommentTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})

	ctx.AfterSuite(func() {
	})
}

func InitializeCommentScenario(ctx *godog.ScenarioContext) {
	//commentFeature := &CommentFeature{
	//	comments: []*models.Comment{},
	//}

	//ctx.Before(func(*godog.Scenario) {
	//	commentFeature.comments = []*models.Comment{}
	//})
	//
	//ctx.After(func(*godog.Scenario, error) {
	//	commentFeature.comments = nil
	//})

	InitializeCommonSteps(ctx)
	InitializeCommentSteps(ctx)
}
