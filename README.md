# godog-demo

A demo repository demonstrating the use of GoDog for behavioral testing.

## Table of Contents

- [Introduction](#introduction)
- [Installation](#installation)
- [Usage](#usage)
- [Running Tests](#running-tests)
- [Project Structure](#project-structure)
- [References](#references)

## Introduction

This repository provides a simple example of how to use GoDog for acceptance testing in a Go project. GoDog is a 
behavior-driven development (BDD) framework for Go.

## Installation

To get started, clone the repository and install the necessary dependencies:

```sh
git clone https://github.com/yourusername/godog-demo.git
cd godog-demo
make install
```

## Useful JetBrains IDE Plugins

- [Cucumber +](https://plugins.jetbrains.com/plugin/16289-cucumber-)
- [Cucumber Go](https://plugins.jetbrains.com/plugin/24323-cucumber-go)
- [Gherkin](https://plugins.jetbrains.com/plugin/9164-gherkin)
- [Gherkin Overview](https://plugins.jetbrains.com/plugin/16716-gherkin-overview)




## Usage

### Writing Features

Create feature files in the features directory.  Each feature file should describe the behavior you want to test 
using Gherkin syntax.

Example features/example.feature:

```gherkin
Feature: Example Feature

  Scenario: Example Scenario
    Given I have a new calculator
    When I add 2 and 3
    Then the result should be 5
```

### Implementing Steps

Implement the steps defined in your feature files in Go.  Place the step definitions in the **steps** directory.

Example steps/example_steps.go:

```go
package steps

import (
    "github.com/cucumber/godog"
    "github.com/yourusername/godog-demo/pkg/calculator"
)

func (c *CalculatorFeatureContext) iHaveANewCalculator() error {
    c.calculator = calculator.NewCalculator()
    return nil
}

func (c *CalculatorFeatureContext) iAddAnd(arg1, arg2 int) error {
    c.result = c.calculator.Add(arg1, arg2)
    return nil
}

func (c *CalculatorFeatureContext) theResultShouldBe(expected int) error {
    if c.result != expected {
        return godog.ErrPending
    }
    return nil
}

func FeatureContext(s *godog.Suite) {
    c := &CalculatorFeatureContext{}
    s.Step(`^I have a new calculator$`, c.iHaveANewCalculator)
    s.Step(`^I add (\d+) and (\d+)$`, c.iAddAnd)
    s.Step(`^the result should be (\d+)$`, c.theResultShouldBe)
}
```


## References

- [GoDog - Repo](https://github.com/cucumber/godog/)
- [Examples](https://github.com/cucumber/godog/tree/main/_examples)
- [GOLANG & API Testing with GoDog](https://medium.com/propertyfinder-engineering/golang-api-testing-with-godog-2de8944d2511)
