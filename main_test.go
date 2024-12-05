package godog_demo

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/ocrosby/godog-demo/pkg/steps"
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
		"features/albums.feature":   {steps.InitializeAlbumScenario, nil},
		"features/comments.feature": {steps.InitializeCommentScenario, steps.InitializeCommentTestSuite},
		"features/posts.feature":    {steps.InitializePostScenario, steps.InitializePostTestSuite},
	}

	for featurePath, initializers := range featureTests {
		t.Run(featurePath, func(t *testing.T) {
			runFeatureTests(t, featurePath, initializers.scenarioInitializer, initializers.testSuiteInitializer)
		})
	}
}
