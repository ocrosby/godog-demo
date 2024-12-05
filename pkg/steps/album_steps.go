package steps

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg/models"
	"io"
)

func thereShouldBeAlbumsInTheResponseBody(expectedAlbumCount int) error {
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

	var albums []models.Album
	if err := json.Unmarshal(body, &albums); err != nil {
		return fmt.Errorf("failed to unmarshal response body: %w", err)
	}

	if len(albums) != expectedAlbumCount {
		return fmt.Errorf("expected %d albums, but got %d", expectedAlbumCount, len(albums))
	}

	return nil
}

type albumFeature struct {
	albums []*models.Album
}

func InitializeAlbumScenario(ctx *godog.ScenarioContext) {
	albumFeature := &albumFeature{
		albums: []*models.Album{},
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		albumFeature.albums = []*models.Album{}
		return ctx, nil
	})

	InitializeCommonSteps(ctx)
	InitializeAlbumSteps(ctx)
}

func InitializeAlbumSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^there should be (\d+) albums in the response body$`, thereShouldBeAlbumsInTheResponseBody)
}
