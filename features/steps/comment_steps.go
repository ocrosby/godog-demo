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

// const commentsKey contextKey = "comments"
// const errorsKey contextKey = "errors"

func withComment(ctx context.Context, comment *models.Comment) context.Context {
	return context.WithValue(ctx, commentKey, comment)
}

//	func withComments(ctx context.Context, comments []*models.Comment) context.Context {
//		return context.WithValue(ctx, commentsKey, comments)
//	}
//
//	func withErrors(ctx context.Context, errors []error) context.Context {
//		return context.WithValue(ctx, errorsKey, errors)
//	}

func getComment(ctx context.Context) *models.Comment {
	return ctx.Value(commentKey).(*models.Comment)
}

func getComments(ctx context.Context) []*models.Comment {
	return ctx.Value(commentKey).([]*models.Comment)
}

//func getErrors(ctx context.Context) []error {
//	return ctx.Value(errorsKey).([]error)
//}

func aNewComment(ctx context.Context) (context.Context, error) {
	comment := &models.Comment{}
	return withComment(ctx, comment), nil
}

func theNewCommentHasAPostIdOf1(ctx context.Context) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.PostID = 1
	return withComment(ctx, comment), nil
}

func theNewCommentHasAnIdOf(ctx context.Context) (context.Context, error) {
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
		err = fmt.Errorf("comment not found in context")
		return withError(ctx, err), err
	}

	body, err = json.Marshal(comment)
	if err != nil {
		err = fmt.Errorf("failed to marshal comment: %w", err)
		return withError(ctx, err), err
	}

	url := helpers.ResolveUrl("/comments")
	if response, err = helpers.SendRequest("POST", url, body); err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return withError(ctx, err), err
	} else {
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				withError(ctx, err)
				fmt.Println("failed to close response body:", err)
			}
		}(response.Body)
		if response.StatusCode != http.StatusCreated {
			err = fmt.Errorf("expected status code %d, but got %d", http.StatusCreated, response.StatusCode)
			return withError(ctx, err), err
		}
	}

	var responseBody map[string]interface{}
	if err = json.Unmarshal(body, &responseBody); err != nil {
		err = fmt.Errorf("failed to unmarshal response body: %w", err)
		return withError(ctx, err), err
	}

	if id, ok := responseBody["id"].(float64); ok {
		comment.ID = int(id)
	} else {
		err = fmt.Errorf("failed to parse comment ID from response body")
		return withError(ctx, err), err
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
	var comments []*models.Comment

	if comments = getComments(ctx); comments == nil {
		return ctx, fmt.Errorf("comments not found in context")
	}

	if len(comments) != expectedCount {
		return ctx, fmt.Errorf("expected %d comments, but got %d", expectedCount, len(comments))
	}

	return ctx, nil
}

func iRequestComment(ctx context.Context, commentId int) (context.Context, error) {
	var err error

	resource := fmt.Sprintf("/comments/%d", commentId)
	url := helpers.ResolveUrl(resource)

	response, err := helpers.SendRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("failed to send request: %w", err)
		return withError(ctx, err), err
	}

	var body []byte
	if body, err = io.ReadAll(response.Body); err != nil {
		err = fmt.Errorf("failed to read response body: %w", err)
		return withError(ctx, err), err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			withError(ctx, err)
			fmt.Println("failed to close response body:", err)
		}
	}(response.Body)

	var comment models.Comment
	if err = json.Unmarshal(body, &comment); err != nil {
		return withError(ctx, err), err
	}

	return withComment(ctx, &comment), nil
}

func theCommentShouldHaveAPostIdOf(ctx context.Context, expectedPostId int) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	if comment.PostID != expectedPostId {
		return ctx, fmt.Errorf("expected post ID %d, but got %d", expectedPostId, comment.PostID)
	}

	return ctx, nil
}

func theCommentShouldHaveANameOf(ctx context.Context, expectedName string) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	if comment.Name != expectedName {
		return ctx, fmt.Errorf("expected name %s, but got %s", expectedName, comment.Name)
	}

	return ctx, nil
}

func theCommentShouldHaveAnEmailOf(ctx context.Context, expectedEmail string) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	if comment.Email != expectedEmail {
		return ctx, fmt.Errorf("expected email %s, but got %s", expectedEmail, comment.Email)
	}

	return ctx, nil
}

func theCommentShouldHaveABodyOf(ctx context.Context, expectedBody string) (context.Context, error) {
	comment := getComment(ctx)
	if comment == nil {
		return ctx, fmt.Errorf("comment not found in context")
	}

	if comment.Body != expectedBody {
		return ctx, fmt.Errorf("expected body '%s', but got '%s'", expectedBody, comment.Body)
	}

	return ctx, nil
}

func iDeleteACommentWithId(ctx context.Context, commentId int) (context.Context, error) {
	resource := fmt.Sprintf("/comments/%d", commentId)
	url := helpers.ResolveUrl(resource)

	response, err := helpers.SendRequest("DELETE", url, nil)
	if err != nil {
		return ctx, fmt.Errorf("failed to send request: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return ctx, fmt.Errorf("expected status code %d, but got %d", http.StatusOK, response.StatusCode)
	}

	return ctx, nil
}

func iSendARequestTo(ctx context.Context, arg1, arg2 string) error {
	return godog.ErrPending
}

func theResponseShouldBeSuccessful(ctx context.Context) error {
	return godog.ErrPending
}

func InitializeCommentTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})

	ctx.AfterSuite(func() {
	})
}

func InitializeCommentScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) comments in the response body$`, thereShouldBeCommentsInTheResponseBody)
	ctx.Step(`^a new comment$`, aNewComment)
	ctx.Step(`^the new comment has a post id of (\d+)$`, theNewCommentHasAPostIdOf1)
	ctx.Step(`^the new comment has an id of (\d+)$`, theNewCommentHasAnIdOf)
	ctx.Step(`^the new comment has a name of "([^"]*)"$`, theNewCommentHasANameOf)
	ctx.Step(`^the new comment has an email of "([^"]*)"$`, theNewCommentHasAnEmailOf)
	ctx.Step(`^the new comment has a body of "([^"]*)"$`, theNewCommentHasABodyOf)
	ctx.Step(`^I create the new comment$`, iCreateTheNewComment)
	ctx.Step(`^the comment should have an id of (\d+)$`, theCommentShouldHaveAnIdOf)
	ctx.Step(`^there should be (\d+) comments in the response body$`, thereShouldBeCommentsInTheResponseBody)
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
	ctx.Step(`^I request comment (\d+)$`, iRequestComment)
	ctx.Step(`^the comment should have a post id of (\d+)$`, theCommentShouldHaveAPostIdOf)
	ctx.Step(`^the comment should have a name of "([^"]*)"$`, theCommentShouldHaveANameOf)
	ctx.Step(`^the comment should have an email of "([^"]*)"$`, theCommentShouldHaveAnEmailOf)
	ctx.Step(`^the comment should have a body of "([^"]*)"$`, theCommentShouldHaveABodyOf)
	ctx.Step(`^I delete a comment with id (\d+)$`, iDeleteACommentWithId)
}
