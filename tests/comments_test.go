package main

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/ocrosby/godog-demo/pkg/steps"
	"os"
	"testing"
	"time"
)

func TestCommentFeatures(t *testing.T) {
	opts := godog.Options{
		Format:    "pretty",
		Paths:     []string{"../features/comments.feature"}, // This includes only the feature file for comments
		Randomize: time.Now().UTC().UnixNano(),              // Optional: randomizes scenario execution order
		Output:    colors.Colored(os.Stdout),                // Output is colored by default
	}

	status := godog.TestSuite{
		Name:                 "Comment Features",
		Options:              &opts,
		TestSuiteInitializer: steps.InitializeCommentTestSuite,
		ScenarioInitializer:  steps.InitializeCommentScenario,
	}.Run()

	if status != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
