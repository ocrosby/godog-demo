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
	return &albumFeature{
		response:  nil,
		body:      nil,
		url:       "",
		resource:  "",
		lastError: nil,
		newAlbum:  nil,
		album:     nil,
		albums:    nil,
	}
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

	f.body, f.lastError = io.ReadAll(f.response.Body)
	if f.lastError != nil {
		return fmt.Errorf("failed to read response body: %w", f.lastError)
	}
	defer f.response.Body.Close()

	return nil
}

func (f *albumFeature) unmarshalResponseBody(target interface{}) error {
	if f.lastError = json.Unmarshal(f.body, target); f.lastError != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", f.lastError)
	}
	return nil
}

func (f *albumFeature) iCreateANewAlbum() error {
	body, err := json.Marshal(f.newAlbum)
	if err != nil {
		return fmt.Errorf("failed to marshal new album: %w", err)
	}

	f.url = helpers.ResolveUrl("/albums")

	f.response, f.lastError = helpers.SendRequest("POST", f.url, body)
	if f.lastError != nil {
		return f.lastError
	}

	var responseBodyStr string
	var responseBody map[string]interface{}

	responseBodyStr, f.lastError = helpers.ReadResponseBody(f.response)
	if f.lastError != nil {
		return f.lastError
	}

	if err := json.Unmarshal([]byte(responseBodyStr), &responseBody); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if id, ok := responseBody["id"].(float64); ok {
		fmt.Println("Response ID:", int(id))
	} else {
		return fmt.Errorf("response does not contain an id property")
	}

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
		return fmt.Errorf("expected no errors, but got %v", f.lastError)
	}
	return nil
}

func (f *albumFeature) theResponseShouldBeSuccessful() error {
	if f.response.StatusCode < 200 || f.response.StatusCode >= 300 {
		return fmt.Errorf("expected status code to be successful, but got %d", f.response.StatusCode)
	}
	return nil
}

func (f *albumFeature) theResponseShouldBeUnsuccessful() error {
	if f.response.StatusCode < 200 || f.response.StatusCode >= 300 {
		return nil
	}

	return fmt.Errorf("expected status code to be unsuccessful, but got %d", f.response.StatusCode)
}

func (f *albumFeature) theAlbumShouldHaveATitleOf(expectedTitle string) error {
	if f.album.Title != expectedTitle {
		return fmt.Errorf("expected album title %q, but got %q", expectedTitle, f.album.Title)
	}
	return nil
}

func (f *albumFeature) theAlbumShouldHaveAUserIdOf(expectedUserId int) error {
	if f.album.UserID != expectedUserId {
		return fmt.Errorf("expected album userId %d, but got %d", expectedUserId, f.album.UserID)
	}
	return nil
}

func (f *albumFeature) theAlbumShouldHaveAnIdOf(expectedId int) error {
	if f.album.ID != expectedId {
		return fmt.Errorf("expected album ID %d, but got %d", expectedId, f.album.ID)
	}
	return nil
}

func (f *albumFeature) thereShouldBeAlbumsInTheResponseBody(expectedAlbumCount int) error {
	if len(f.albums) != expectedAlbumCount {
		return fmt.Errorf("expected %d albums in the response body, but got %d", expectedAlbumCount, len(f.albums))
	}
	return nil
}

func InitializeAlbumTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
	})

	ctx.AfterSuite(func() {
	})
}

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
