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

// commentKey is the context key used to store and retrieve the current Comment
// being built or inspected within a scenario.
const commentKey contextKey = "comment"

// withComment stores comment in ctx under commentKey and returns the updated context.
func withComment(ctx context.Context, comment *models.Comment) context.Context {
	return context.WithValue(ctx, commentKey, comment)
}

// getComment retrieves the Comment stored by withComment from ctx.
// It returns (nil, false) if no comment has been stored or the stored value is
// not a *models.Comment.
func getComment(ctx context.Context) (*models.Comment, bool) {
	c, ok := ctx.Value(commentKey).(*models.Comment)
	return c, ok
}

// aNewComment initialises a blank Comment in the scenario context, ready for
// subsequent "the new comment has a …" steps to populate.
func aNewComment(ctx context.Context) (context.Context, error) {
	return withComment(ctx, &models.Comment{}), nil
}

// theNewCommentHasAPostIdOf sets the PostID on the comment being built in ctx.
// It returns an error when no comment has been initialised in the context.
func theNewCommentHasAPostIdOf(ctx context.Context, postId int) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.PostID = postId
	return withComment(ctx, comment), nil
}

// theNewCommentHasAnIdOf sets the ID on the comment being built in ctx.
// It returns an error when no comment has been initialised in the context.
func theNewCommentHasAnIdOf(ctx context.Context, id int) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.ID = id
	return withComment(ctx, comment), nil
}

// theNewCommentHasANameOf sets the Name on the comment being built in ctx.
// It returns an error when no comment has been initialised in the context.
func theNewCommentHasANameOf(ctx context.Context, value string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Name = value
	return withComment(ctx, comment), nil
}

// theNewCommentHasAnEmailOf sets the Email on the comment being built in ctx.
// It returns an error when no comment has been initialised in the context.
func theNewCommentHasAnEmailOf(ctx context.Context, value string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Email = value
	return withComment(ctx, comment), nil
}

// theNewCommentHasABodyOf sets the Body on the comment being built in ctx.
// It returns an error when no comment has been initialised in the context.
func theNewCommentHasABodyOf(ctx context.Context, value string) (context.Context, error) {
	comment, ok := getComment(ctx)
	if !ok {
		return ctx, fmt.Errorf("comment not found in context")
	}
	comment.Body = value
	return withComment(ctx, comment), nil
}

// postComment marshals comment to JSON, POSTs it to POST /comments, stores the
// response in lastResponse, checks that the API returned HTTP 201 Created, and
// returns the server-assigned ID extracted from the response body.
//
// It is extracted from iCreateTheNewComment to keep that function within its
// complexity budget. It returns (0, err) on any failure.
func postComment(comment *models.Comment) (int, error) {
	payload, err := json.Marshal(comment)
	if err != nil {
		return 0, fmt.Errorf("marshalling comment: %w", err)
	}

	resp, err := helpers.SendRequest("POST", helpers.ResolveUrl("/comments"), payload)
	if err != nil {
		return 0, fmt.Errorf("sending POST /comments: %w", err)
	}
	lastResponse = resp

	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		return 0, fmt.Errorf("expected status %d, got %d", http.StatusCreated, resp.StatusCode)
	}

	return helpers.HandlePostResponse(resp, comment)
}

// iCreateTheNewComment POSTs the comment built by prior steps, sets the
// server-assigned ID on the comment, and stores the updated comment back in ctx.
// Any error is both returned and stored in ctx so that assertion steps can
// detect it with "the result should be an error".
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

// theCommentShouldHaveAnIdOf returns an error when the comment stored in ctx
// does not have the expected ID.
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

// iRequestComment fetches the comment identified by commentId from
// GET /comments/{id}, stores the response in lastResponse for assertion steps,
// and stores the unmarshalled comment in ctx.
// It returns a wrapped error on any transport, read, or unmarshal failure.
func iRequestComment(ctx context.Context, commentId int) (context.Context, error) {
	url := helpers.ResolveUrl(fmt.Sprintf("/comments/%d", commentId))

	resp, err := helpers.SendRequest("GET", url, nil)
	if err != nil {
		return withError(ctx, err), fmt.Errorf("sending GET /comments/%d: %w", commentId, err)
	}
	lastResponse = resp

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return withError(ctx, err), fmt.Errorf("reading response body: %w", err)
	}

	var comment models.Comment
	if err := json.Unmarshal(body, &comment); err != nil {
		return withError(ctx, err), fmt.Errorf("unmarshalling comment: %w", err)
	}

	return withComment(ctx, &comment), nil
}

// theCommentShouldHaveAPostIdOf returns an error when the comment in ctx does
// not have the expected PostID.
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

// theCommentShouldHaveANameOf returns an error when the comment in ctx does not
// have the expected Name.
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

// theCommentShouldHaveAnEmailOf returns an error when the comment in ctx does
// not have the expected Email.
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

// theCommentShouldHaveABodyOf returns an error when the comment in ctx does not
// have the expected Body.
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

// iDeleteACommentWithId sends DELETE /comments/{commentId}, stores the response
// in lastResponse, and returns an error when the API does not respond with HTTP 200 OK.
func iDeleteACommentWithId(ctx context.Context, commentId int) (context.Context, error) {
	url := helpers.ResolveUrl(fmt.Sprintf("/comments/%d", commentId))

	resp, err := helpers.SendRequest("DELETE", url, nil)
	if err != nil {
		return ctx, fmt.Errorf("sending DELETE /comments/%d: %w", commentId, err)
	}
	lastResponse = resp
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ctx, fmt.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	return ctx, nil
}

// thereShouldBeCommentsInTheResponseBody asserts that the most recent API
// response body contains exactly expected Comment objects. It delegates to the
// shared assertResponseBodyCount generic helper.
func thereShouldBeCommentsInTheResponseBody(expected int) error {
	return assertResponseBodyCount[models.Comment](expected)
}

// InitializeCommentTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for comment scenarios.
func InitializeCommentTestSuite(_ *godog.TestSuiteContext) {}

// InitializeCommentScenario wires all comment step definitions to their Gherkin
// patterns.
func InitializeCommentScenario(ctx *godog.ScenarioContext) {
	InitializeCommonSteps(ctx)

	ctx.Step(`^there should be (\d+) comments in the response body$`, thereShouldBeCommentsInTheResponseBody)
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
