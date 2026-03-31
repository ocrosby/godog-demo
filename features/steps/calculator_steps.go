package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg"
)

// calculatorKey is the context key used to store and retrieve the Calculator
// instance shared across steps within a single scenario.
const calculatorKey contextKey = "calculator"

// withCalculator stores calculator in ctx under calculatorKey and returns the
// updated context.
func withCalculator(ctx context.Context, calculator *pkg.Calculator) context.Context {
	return context.WithValue(ctx, calculatorKey, calculator)
}

// getCalculator retrieves the Calculator stored by withCalculator from ctx.
// It returns (nil, false) if no calculator has been stored or the stored value
// is not a *pkg.Calculator.
func getCalculator(ctx context.Context) (*pkg.Calculator, bool) {
	c, ok := ctx.Value(calculatorKey).(*pkg.Calculator)
	return c, ok
}

// iAddAnd calls Add(arg1, arg2) on the calculator in ctx and returns an error
// if no calculator has been initialised for this scenario.
func iAddAnd(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	calculator.Add(arg1, arg2)
	return ctx, nil
}

// iDivideBy calls Divide(arg1, arg2) on the calculator in ctx. If division
// fails (e.g. divide by zero) the error is stored in ctx via withError so that
// a subsequent "the result should be an error" step can verify it; the step
// itself returns nil so the scenario continues.
func iDivideBy(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	_, err := calculator.Divide(arg1, arg2)
	if err != nil {
		return withError(ctx, err), nil
	}
	return ctx, nil
}

// iHaveANewCalculator creates a fresh Calculator and stores it in ctx, making
// it available to all subsequent arithmetic steps in this scenario.
func iHaveANewCalculator(ctx context.Context) (context.Context, error) {
	return withCalculator(ctx, pkg.NewCalculator()), nil
}

// iMultiplyBy calls Multiply(arg1, arg2) on the calculator in ctx and returns
// an error if no calculator has been initialised for this scenario.
func iMultiplyBy(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	calculator.Multiply(arg1, arg2)
	return ctx, nil
}

// iSubtractFrom calls Subtract(arg2, arg1) on the calculator in ctx, mapping
// the Gherkin phrase "subtract X from Y" to Subtract(Y, X). It returns an
// error if no calculator has been initialised for this scenario.
func iSubtractFrom(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	calculator.Subtract(arg2, arg1)
	return ctx, nil
}

// theResultShouldBe asserts that the calculator's Accumulator equals expected.
// It returns an error if the values differ or if no calculator is in ctx.
func theResultShouldBe(ctx context.Context, expected int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	actual := calculator.GetAccumulator()
	if expected != actual {
		return ctx, fmt.Errorf("expected result %d, got %d", expected, actual)
	}
	return ctx, nil
}

// theResultShouldBeAnError asserts that a prior step stored an error in ctx via
// withError. It returns an error if no error is present, indicating an
// unexpected success.
func theResultShouldBeAnError(ctx context.Context) (context.Context, error) {
	if getError(ctx) == nil {
		return ctx, fmt.Errorf("expected an error, got nil")
	}
	return ctx, nil
}

// InitializeCalculatorTestSuite satisfies the godog.TestSuiteInitializer signature.
// No suite-level setup is required for calculator scenarios.
func InitializeCalculatorTestSuite(_ *godog.TestSuiteContext) {}

// InitializeCalculatorScenario wires all calculator step definitions to their
// Gherkin patterns.
func InitializeCalculatorScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I add (\d+) and (\d+)$`, iAddAnd)
	ctx.Step(`^I divide (\d+) by (\d+)$`, iDivideBy)
	ctx.Step(`^I have a new calculator$`, iHaveANewCalculator)
	ctx.Step(`^I multiply (\d+) by (\d+)$`, iMultiplyBy)
	ctx.Step(`^I subtract (\d+) from (\d+)$`, iSubtractFrom)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
	ctx.Step(`^the result should be an error$`, theResultShouldBeAnError)
}
