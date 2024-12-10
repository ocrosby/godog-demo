package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/helpers"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
	"net/http"
)

type contextKey string

const commentKey contextKey = "comment"

func withComment(ctx context.Context, comment *models.Comment) context.Context {
	return context.WithValue(ctx, commentKey, comment)
}

func getComment(ctx context.Context) *models.Comment {
	return ctx.Value(commentKey).(*models.Comment)
}

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
	ctx.Step(`^a new comment$`, aNewComment)
	ctx.Step(`^the new comment has a post id of (\d+)$`, theNewCommentHasAPostIdOf1)
	ctx.Step(`^the new comment has an id of (\d+)$`, theNewCommentHasAnIdOf501)
	ctx.Step(`^the new comment has a name of "([^"]*)"$`, theNewCommentHasANameOf)
	ctx.Step(`^the new comment has an email of "([^"]*)"$`, theNewCommentHasAnEmailOf)
	ctx.Step(`^the new comment has a body of "([^"]*)"$`, theNewCommentHasABodyOf)
	ctx.Step(`^I create the new comment$`, iCreateTheNewComment)
	ctx.Step(`^the comment should have an id of (\d+)$`, theCommentShouldHaveAnIdOf)
	ctx.Step(`^there should be (\d+) comments in the response body$`, thereShouldBeCommentsInTheResponseBody)
	ctx.Step(`^I request comment (\d+)$`, iRequestComment)
	ctx.Step(`^the comment should have a post id of (\d+)$`, theCommentShouldHaveAPostIdOf1)
	ctx.Step(`^the comment should have a name of "([^"]*)"$`, theCommentShouldHaveANameOfDolorumUtInVoluptas)
	ctx.Step(`^the comment should have an email of "([^"]*)"$`, theCommentShouldHaveAnEmailOfSomething)
	ctx.Step(`^the comment should have a body of "([^"]*)"$`, theCommentShouldHaveABodyOfQuidemMolestiaeEnim)
	ctx.Step(`^I delete a comment with id (\d+)$`, iDeleteACommentWithId1)
}

func aNewComment(ctx context.Context, ) (context.Context, error) {
	comment := &models.Comment{}
	return withComment(ctx, comment), nil
}

func theNewCommentHasAPostIdOf1(ctx context.Context, ) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.PostID = 1
	return withComment(ctx, comment), nil
}

func theNewCommentHasAnIdOf501(ctx context.Context, ) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.ID = 501
	return withComment(ctx, comment), nil
}

func theNewCommentHasANameOf(ctx context.Context, value string) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Name = value
	return withComment(ctx, comment), nil
}

func theNewCommentHasAnEmailOf(ctx context.Context, value string) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Email = value
	return withComment(ctx, comment), nil
}

func theNewCommentHasABodyOf(ctx context.Context, value string) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Body = value
	return withComment(ctx, comment), nil
}

func iCreateTheNewComment(ctx context.Context) (context.Context, error) {
	var (
		body     []byte
		err      error
		comment  *models.Comment
		response *http.Response
	)

	if comment = getComment(ctx); comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	body, err = json.Marshal(comment)
	if err != nil {
		return ctx, fmt.Errorf("failed to marshal comment: %w", err)
	}

	url := helpers.ResolveUrl("/comments")
	if response, err = helpers.SendRequest("POST", url, body); err != nil {
		return ctx, fmt.Errorf("failed to send request: %w", err)
	} else {
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				fmt.Println("failed to close response body:", err)
			}
		}(response.Body)
		if response.StatusCode != http.StatusCreated {
			return ctx, fmt.Errorf("expected status code %d, but got %d", http.StatusCreated, response.StatusCode)
		}
	}

	var responseBody map[string]interface{}
	if err = json.Unmarshal(body, &responseBody); err != nil {
		return ctx, fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if id, ok := responseBody["id"].(float64); ok {
		comment.ID = int(id)
	} else {
		return ctx, fmt.Errorf("failed to parse comment ID from response body")
	}

	return withComment(ctx, comment), nil
}

func theCommentShouldHaveAnIdOf(ctx context.Context, value int) (context.Context, error) {
	var comment *models.Comment

	if comment = getComment(ctx); comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	if comment.ID != value {
		return ctx, fmt.Errorf("expected comment ID %d, but got %d", value, comment.ID)
	}

	return ctx, nil
}

func thereShouldBeCommentsInTheResponseBody(ctx context.Context, expectedCount int) (context.Context, error) {
	if comment := getComment(ctx); comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	if len(comment) != expectedCount {

	}

}
func iRequestComment(ctx context.Context, ) (context.Context, error) {

	return ctx, godog.ErrPending
}
func theCommentShouldHaveAPostIdOf1(ctx context.Context, ) (context.Context, error) {
	return ctx, godog.ErrPending
}
func theCommentShouldHaveANameOfDolorumUtInVoluptas(ctx context.Context, ) (context.Context, error) {
	return ctx, godog.ErrPending
}
func theCommentShouldHaveAnEmailOfSomething(ctx context.Context, ) (context.Context, error) {
	return ctx, godog.ErrPending
}
func theCommentShouldHaveABodyOfQuidemMolestiaeEnim(ctx context.Context, ) (context.Context, error) {
	return ctx, godog.ErrPending
}
func iDeleteACommentWithId1(ctx context.Context, ) (context.Context, error) {
	return ctx, godog.ErrPending
}
