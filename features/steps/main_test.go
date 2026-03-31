package steps

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
	"time"
)

func runFeatureTests(t *testing.T, featurePath string, scenarioInitializer func(*godog.ScenarioContext), testSuiteInitializer func(*godog.TestSuiteContext)) {
	opts := godog.Options{
		Format:    "pretty",
		Paths:     []string{featurePath},
		Randomize: time.Now().UTC().UnixNano(),
		Output:    colors.Colored(os.Stdout),
	}

	status := godog.TestSuite{
		ScenarioInitializer:  scenarioInitializer,
		TestSuiteInitializer: testSuiteInitializer,
		Options:              &opts,
	}.Run()

	if status != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func TestFeatures(t *testing.T) {
	featureTests := map[string]struct {
		scenarioInitializer  func(*godog.ScenarioContext)
		testSuiteInitializer func(*godog.TestSuiteContext)
	}{
		"../albums.feature":     {InitializeAlbumScenario, InitializeAlbumTestSuite},
		"../comments.feature":   {InitializeCommentScenario, InitializeCommentTestSuite},
		"../photos.feature":     {InitializePhotoScenario, InitializePhotoTestSuite},
		"../posts.feature":      {InitializePostScenario, InitializePostTestSuite},
		"../todos.feature":      {InitializeTodoScenario, InitializeTodoTestSuite},
		"../users.feature":      {InitializeUserScenario, InitializeUserTestSuite},
		"../calculator.feature": {InitializeCalculatorScenario, InitializeCalculatorTestSuite},
	}

	for featurePath, initializers := range featureTests {
		t.Run(featurePath, func(t *testing.T) {
			runFeatureTests(t, featurePath, initializers.scenarioInitializer, initializers.testSuiteInitializer)
		})
	}
}

