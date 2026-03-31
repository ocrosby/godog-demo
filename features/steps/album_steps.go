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

type albumFeature struct {
	response  *http.Response
	body      []byte
	url       string
	resource  string
	lastError error
	newAlbum  *models.Album
	album     *models.Album
	albums    []*models.Album
}

func newAlbumFeature() *albumFeature {
	return &albumFeature{}
}

func (f *albumFeature) aNewAlbum() error {
	f.newAlbum = &models.Album{}
	return nil
}

func (f *albumFeature) theNewAlbumHasId(id int) error {
	f.newAlbum.ID = id
	return nil
}

func (f *albumFeature) theNewAlbumHasUserId(userId int) error {
	f.newAlbum.UserID = userId
	return nil
}

func (f *albumFeature) theNewAlbumHasTitle(title string) error {
	f.newAlbum.Title = title
	return nil
}

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

func (f *albumFeature) unmarshalResponseBody(target interface{}) error {
	if err := json.Unmarshal(f.body, target); err != nil {
		f.lastError = err
		return fmt.Errorf("unmarshalling response body: %w", err)
	}
	return nil
}

func (f *albumFeature) iCreateANewAlbum() error {
	body, err := json.Marshal(f.newAlbum)
	if err != nil {
		return fmt.Errorf("marshalling new album: %w", err)
	}

	f.response, f.lastError = helpers.SendRequest("POST", helpers.ResolveUrl("/albums"), body)
	if f.lastError != nil {
		return f.lastError
	}

	id, err := helpers.HandlePostResponse(f.response, &f.newAlbum)
	if err != nil {
		return fmt.Errorf("handling POST response: %w", err)
	}

	f.album = &models.Album{ID: id}
	return nil
}

func (f *albumFeature) iRequestAllAlbums() error {
	if err := f.sendRequest("GET", "/albums"); err != nil {
		return err
	}
	return f.unmarshalResponseBody(&f.albums)
}

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

func (f *albumFeature) iDeleteAlbum(albumId int) error {
	return f.sendRequest("DELETE", fmt.Sprintf("/albums/%d", albumId))
}

func (f *albumFeature) thereShouldBeNoErrors() error {
	if f.lastError != nil {
		return fmt.Errorf("expected no errors, got %v", f.lastError)
	}
	return nil
}

func (f *albumFeature) theResponseShouldBeSuccessful() error {
	if f.response.StatusCode < http.StatusOK || f.response.StatusCode >= http.StatusMultipleChoices {
		return fmt.Errorf("expected successful status code, got %d", f.response.StatusCode)
	}
	return nil
}

func (f *albumFeature) theResponseShouldBeUnsuccessful() error {
	if f.response.StatusCode < http.StatusOK || f.response.StatusCode >= http.StatusMultipleChoices {
		return nil
	}
	return fmt.Errorf("expected unsuccessful status code, got %d", f.response.StatusCode)
}

func (f *albumFeature) theAlbumShouldHaveATitleOf(expected string) error {
	if f.album.Title != expected {
		return fmt.Errorf("expected album title %q, got %q", expected, f.album.Title)
	}
	return nil
}

func (f *albumFeature) theAlbumShouldHaveAUserIdOf(expected int) error {
	if f.album.UserID != expected {
		return fmt.Errorf("expected album userId %d, got %d", expected, f.album.UserID)
	}
	return nil
}

func (f *albumFeature) theAlbumShouldHaveAnIdOf(expected int) error {
	if f.album.ID != expected {
		return fmt.Errorf("expected album ID %d, got %d", expected, f.album.ID)
	}
	return nil
}

func (f *albumFeature) thereShouldBeAlbumsInTheResponseBody(expected int) error {
	if len(f.albums) != expected {
		return fmt.Errorf("expected %d albums in response body, got %d", expected, len(f.albums))
	}
	return nil
}

func InitializeAlbumTestSuite(_ *godog.TestSuiteContext) {}

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
