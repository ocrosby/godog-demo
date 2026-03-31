package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/helpers"
	"github.com/ocrosby/godog-demo/pkg/models"
)

const commentKey contextKey = "comment"

func withComment(ctx context.Context, comment *models.Comment) context.Context {
	return context.WithValue(ctx, commentKey, comment)
}

func getComment(ctx context.Context) (*models.Comment, bool) {
	c, ok := ctx.Value(commentKey).(*models.Comment)
	return c, ok
}

func aNewComment(ctx context.Context) (context.Context, error) {
	return withComment(ctx, &models.Comment{}), nil
}

func theNewCommentHasAPostIdOf(ctx context.Context, postId int) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.PostID = postId
	return withComment(ctx, comment), nil
}

func theNewCommentHasAnIdOf(ctx context.Context, id int) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.ID = id
	return withComment(ctx, comment), nil
}

func theNewCommentHasANameOf(ctx context.Context, value string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Name = value
	return withComment(ctx, comment), nil
}

func theNewCommentHasAnEmailOf(ctx context.Context, value string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Email = value
	return withComment(ctx, comment), nil
}

func theNewCommentHasABodyOf(ctx context.Context, value string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Body = value
	return withComment(ctx, comment), nil
}

// postComment marshals comment, POSTs it to /comments, and returns the server-assigned ID.
// Extracted to keep iCreateTheNewComment below its complexity budget.
func postComment(comment *models.Comment) (int, error) {
	payload, err := json.Marshal(comment)
	if err != nil {
		return 0, fmt.Errorf("marshalling comment: %w", err)
	}

	resp, err := helpers.SendRequest("POST", helpers.ResolveUrl("/comments"), payload)
	if err != nil {
		return 0, fmt.Errorf("sending POST /comments: %w", err)
	}

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return 0, fmt.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	return helpers.HandlePostResponse(resp, comment)
}

func iCreateTheNewComment(ctx context.Context) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}

	id, err := postComment(comment)
	if err != nil {
		return withError(ctx, err), err
	}

	comment.ID = id
	return withComment(ctx, comment), nil
}

func theCommentShouldHaveAnIdOf(ctx context.Context, expected int) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	if comment.ID != expected {
		return ctx, fmt.Errorf("expected comment ID %d, got %d", expected, comment.ID)
	}
	return ctx, nil
}

func iRequestComment(ctx context.Context, commentId int) (context.Context, error) {
	url := helpers.ResolveUrl(fmt.Sprintf("/comments/%d", commentId))

	resp, err := helpers.SendRequest("GET", url, nil)
	if err != nil {
		return withError(ctx, err), fmt.Errorf("sending GET /comments/%d: %w", commentId, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return withError(ctx, err), fmt.Errorf("reading response body: %w", err)
	}

	var comment models.Comment
	if err := json.Unmarshal(body, &comment); err != nil {
		return withError(ctx, err), fmt.Errorf("unmarshalling comment: %w", err)
	}

	return withComment(ctx, &comment), nil
}

func theCommentShouldHaveAPostIdOf(ctx context.Context, expected int) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	if comment.PostID != expected {
		return ctx, fmt.Errorf("expected post ID %d, got %d", expected, comment.PostID)
	}
	return ctx, nil
}

func theCommentShouldHaveANameOf(ctx context.Context, expected string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	if comment.Name != expected {
		return ctx, fmt.Errorf("expected name %q, got %q", expected, comment.Name)
	}
	return ctx, nil
}

func theCommentShouldHaveAnEmailOf(ctx context.Context, expected string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	if comment.Email != expected {
		return ctx, fmt.Errorf("expected email %q, got %q", expected, comment.Email)
	}
	return ctx, nil
}

func theCommentShouldHaveABodyOf(ctx context.Context, expected string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	if comment.Body != expected {
		return ctx, fmt.Errorf("expected body %q, got %q", expected, comment.Body)
	}
	return ctx, nil
}

func iDeleteACommentWithId(ctx context.Context, commentId int) (context.Context, error) {
	url := helpers.ResolveUrl(fmt.Sprintf("/comments/%d", commentId))

	resp, err := helpers.SendRequest("DELETE", url, nil)
	if err != nil {
		return ctx, fmt.Errorf("sending DELETE /comments/%d: %w", commentId, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ctx, fmt.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	return ctx, nil
}

func InitializeCommentTestSuite(_ *godog.TestSuiteContext) {}

func InitializeCommentScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a new comment$`, aNewComment)
	ctx.Step(`^the new comment has a post id of (\d+)$`, theNewCommentHasAPostIdOf)
	ctx.Step(`^the new comment has an id of (\d+)$`, theNewCommentHasAnIdOf)
	ctx.Step(`^the new comment has a name of "([^"]*)"$`, theNewCommentHasANameOf)
	ctx.Step(`^the new comment has an email of "([^"]*)"$`, theNewCommentHasAnEmailOf)
	ctx.Step(`^the new comment has a body of "([^"]*)"$`, theNewCommentHasABodyOf)
	ctx.Step(`^I create the new comment$`, iCreateTheNewComment)
	ctx.Step(`^the comment should have an id of (\d+)$`, theCommentShouldHaveAnIdOf)
	ctx.Step(`^there should be no errors$`, thereShouldBeNoErrors)
	ctx.Step(`^I request comment (\d+)$`, iRequestComment)
	ctx.Step(`^the comment should have a post id of (\d+)$`, theCommentShouldHaveAPostIdOf)
	ctx.Step(`^the comment should have a name of "([^"]*)"$`, theCommentShouldHaveANameOf)
	ctx.Step(`^the comment should have an email of "([^"]*)"$`, theCommentShouldHaveAnEmailOf)
	ctx.Step(`^the comment should have a body of "([^"]*)"$`, theCommentShouldHaveABodyOf)
	ctx.Step(`^I delete a comment with id (\d+)$`, iDeleteACommentWithId)
}
