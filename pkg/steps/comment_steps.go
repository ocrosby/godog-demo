package steps

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
	"net/http"
)

type commentFeature struct {
	response   *http.Response
	body       []byte
	url        string
	resource   string
	lastError  error
	newComment *models.Comment
	comment    *models.Comment
	comments   []*models.Comment
}

func newCommentFeature() *commentFeature {
	return &commentFeature{
		comments: nil,
	}
}

func (c *commentFeature) thereShouldBeCommentsInTheResponseBody(expectedCommentCount int) error {
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

func InitializeCommentTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})

	ctx.AfterSuite(func() {
	})
}

func InitializeCommentScenario(ctx *godog.ScenarioContext) {
	feature := &commentFeature{
		comments: []*models.Comment{},
	}

	ctx.Step(`^there should be (\d+) comments in the response body$`, feature.thereShouldBeCommentsInTheResponseBody)
}
