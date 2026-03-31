package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/builders"
	"github.com/ocrosby/godog-demo/pkg/helpers"
	"github.com/ocrosby/godog-demo/pkg/models"
	"github.com/ocrosby/godog-demo/pkg/validation"
)

// albumFeature holds the per-scenario state for the album BDD steps.
// A fresh instance is created by newAlbumFeature and bound to every scenario
// via InitializeAlbumScenario, so no state leaks between scenarios.
type albumFeature struct {
	response     *http.Response
	body         []byte
	url          string
	resource     string
	lastError    error
	albumBuilder *builders.AlbumBuilder
	album        *models.Album
	albums       []*models.Album
}

// newAlbumFeature returns an albumFeature with all fields at their zero values.
func newAlbumFeature() *albumFeature {
	return &albumFeature{}
}

// aNewAlbum initialises a fresh AlbumBuilder ready for the scenario to populate
// via the theNewAlbumHas* steps.
func (f *albumFeature) aNewAlbum() error {
	f.albumBuilder = builders.NewAlbumBuilder()
	return nil
}

// theNewAlbumHasId sets the ID on the album being built by the current scenario.
func (f *albumFeature) theNewAlbumHasId(id int) error {
	f.albumBuilder.WithID(id)
	return nil
}

// theNewAlbumHasUserId sets the UserID on the album being built by the current scenario.
func (f *albumFeature) theNewAlbumHasUserId(userId int) error {
	f.albumBuilder.WithUserID(userId)
	return nil
}

// theNewAlbumHasTitle sets the Title on the album being built by the current scenario.
func (f *albumFeature) theNewAlbumHasTitle(title string) error {
	f.albumBuilder.WithTitle(title)
	return nil
}

// sendRequest sends an HTTP request using method to resource (e.g. "/albums/1"),
// stores the response and its body on the receiver, and returns any error.
// The response body is read and closed before this function returns.
func (f *albumFeature) sendRequest(method, resource string) error {
	f.resource = resource
	f.url = helpers.ResolveUrl(f.resource)

	f.response, f.lastError = helpers.SendRequest(method, f.url, nil)
	if f.lastError != nil {
		return f.lastError
	}
	defer f.response.Body.Close()

	f.body, f.lastError = io.ReadAll(f.response.Body)
	if f.lastError != nil {
		return fmt.Errorf("reading response body: %w", f.lastError)
	}

	return nil
}

// unmarshalResponseBody decodes the previously-read response body JSON into target.
// It returns a wrapped error if unmarshalling fails.
func (f *albumFeature) unmarshalResponseBody(target interface{}) error {
	if err := json.Unmarshal(f.body, target); err != nil {
		f.lastError = err
		return fmt.Errorf("unmarshalling response body: %w", err)
	}
	return nil
}

// iCreateANewAlbum builds the album from f.albumBuilder, POSTs it to /albums,
// and stores the server-assigned ID on f.album. It returns an error if the
// builder validation, marshalling, the HTTP request, or the response parsing
// fails.
func (f *albumFeature) iCreateANewAlbum() error {
	album, err := f.albumBuilder.Build()
	if err != nil {
		return fmt.Errorf("building album: %w", err)
	}

	body, err := json.Marshal(album)
	if err != nil {
		return fmt.Errorf("marshalling new album: %w", err)
	}

	f.response, f.lastError = helpers.SendRequest("POST", helpers.ResolveUrl("/albums"), body)
	if f.lastError != nil {
		return f.lastError
	}

	id, err := helpers.HandlePostResponse(f.response, album)
	if err != nil {
		return fmt.Errorf("handling POST response: %w", err)
	}

	f.album = &models.Album{ID: id}
	return nil
}

// iRequestAllAlbums fetches the full list of albums from GET /albums and
// deserialises them into f.albums.
func (f *albumFeature) iRequestAllAlbums() error {
	if err := f.sendRequest("GET", "/albums"); err != nil {
		return err
	}
	return f.unmarshalResponseBody(&f.albums)
}

// iRequestAlbum fetches the album identified by albumId from GET /albums/{id},
// appends it to f.albums, and sets f.album for subsequent assertion steps.
func (f *albumFeature) iRequestAlbum(albumId int) error {
	if err := f.sendRequest("GET", fmt.Sprintf("/albums/%d", albumId)); err != nil {
		return err
	}
	var album models.Album
	if err := f.unmarshalResponseBody(&album); err != nil {
		return err
	}
	f.album = &album
	f.albums = append(f.albums, &album)
	return nil
}

// iDeleteAlbum sends DELETE /albums/{albumId} and returns any transport error.
func (f *albumFeature) iDeleteAlbum(albumId int) error {
	return f.sendRequest("DELETE", fmt.Sprintf("/albums/%d", albumId))
}

// thereShouldBeNoErrors returns an error when a previous step stored one in
// f.lastError, allowing scenarios to assert a clean execution path.
func (f *albumFeature) thereShouldBeNoErrors() error {
	if f.lastError != nil {
		return fmt.Errorf("expected no errors, got %v", f.lastError)
	}
	return nil
}

// theResponseShouldBeSuccessful returns an error when f.response carries a
// non-2xx HTTP status code. It delegates to validation.SuccessValidator.
func (f *albumFeature) theResponseShouldBeSuccessful() error {
	return validation.SuccessValidator{}.Validate(f.response)
}

// theResponseShouldBeUnsuccessful returns an error when f.response carries a
// 2xx HTTP status code, i.e. when the response was unexpectedly successful.
// It delegates to validation.FailureValidator.
func (f *albumFeature) theResponseShouldBeUnsuccessful() error {
	return validation.FailureValidator{}.Validate(f.response)
}

// theAlbumShouldHaveATitleOf returns an error when f.album.Title does not equal expected.
func (f *albumFeature) theAlbumShouldHaveATitleOf(expected string) error {
	if f.album.Title != expected {
		return fmt.Errorf("expected album title %q, got %q", expected, f.album.Title)
	}
	return nil
}

// theAlbumShouldHaveAUserIdOf returns an error when f.album.UserID does not equal expected.
func (f *albumFeature) theAlbumShouldHaveAUserIdOf(expected int) error {
	if f.album.UserID != expected {
		return fmt.Errorf("expected album userId %d, got %d", expected, f.album.UserID)
	}
	return nil
}

// theAlbumShouldHaveAnIdOf returns an error when f.album.ID does not equal expected.
func (f *albumFeature) theAlbumShouldHaveAnIdOf(expected int) error {
	if f.album.ID != expected {
		return fmt.Errorf("expected album ID %d, got %d", expected, f.album.ID)
	}
	return nil
}

// thereShouldBeAlbumsInTheResponseBody returns an error when the number of albums
// fetched into f.albums does not equal expected.
func (f *albumFeature) thereShouldBeAlbumsInTheResponseBody(expected int) error {
	if len(f.albums) != expected {
		return fmt.Errorf("expected %d albums in response body, got %d", expected, len(f.albums))
	}
	return nil
}

// InitializeAlbumTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for album scenarios.
func InitializeAlbumTestSuite(_ *godog.TestSuiteContext) {}

// InitializeAlbumScenario wires all album step definitions to their Gherkin
// patterns and resets per-scenario state before each scenario runs.
func InitializeAlbumScenario(ctx *godog.ScenarioContext) {
	feature := newAlbumFeature()

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		feature.albums = []*models.Album{}
		return ctx, nil
	})

	ctx.Step(`^a new album$`, feature.aNewAlbum)
	ctx.Step(`^the new album has an id of (\d+)$`, feature.theNewAlbumHasId)
	ctx.Step(`^the new album has a user id of (\d+)$`, feature.theNewAlbumHasUserId)
	ctx.Step(`^the new album has a title of "([^"]*)"$`, feature.theNewAlbumHasTitle)
	ctx.Step(`^I create the new album$`, feature.iCreateANewAlbum)
	ctx.Step(`^I request all albums$`, feature.iRequestAllAlbums)
	ctx.Step(`^I request album (\d+)$`, feature.iRequestAlbum)
	ctx.Step(`^I delete album (\d+)$`, feature.iDeleteAlbum)
	ctx.Step(`^there should be no errors$`, feature.thereShouldBeNoErrors)
	ctx.Step(`^the response should be successful$`, feature.theResponseShouldBeSuccessful)
	ctx.Step(`^the response should be unsuccessful$`, feature.theResponseShouldBeUnsuccessful)
	ctx.Step(`^the album should have a title of "([^"]*)"$`, feature.theAlbumShouldHaveATitleOf)
	ctx.Step(`^the album should have a user id of (\d+)$`, feature.theAlbumShouldHaveAUserIdOf)
	ctx.Step(`^the album should have an id of (\d+)$`, feature.theAlbumShouldHaveAnIdOf)
	ctx.Step(`^there should be (\d+) albums in the response body$`, feature.thereShouldBeAlbumsInTheResponseBody)
}
