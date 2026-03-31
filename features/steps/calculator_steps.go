package steps

import (
	"context"
	"fmt"

	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg"
)

const calculatorKey contextKey = "calculator"

func withCalculator(ctx context.Context, calculator *pkg.Calculator) context.Context {
	return context.WithValue(ctx, calculatorKey, calculator)
}

func getCalculator(ctx context.Context) (*pkg.Calculator, bool) {
	c, ok := ctx.Value(calculatorKey).(*pkg.Calculator)
	return c, ok
}

func iAddAnd(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	calculator.Add(arg1, arg2)
	return ctx, nil
}

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

func iHaveANewCalculator(ctx context.Context) (context.Context, error) {
	return withCalculator(ctx, pkg.NewCalculator()), nil
}

func iMultiplyBy(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	calculator.Multiply(arg1, arg2)
	return ctx, nil
}

func iSubtractFrom(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator, ok := getCalculator(ctx)
	if !ok {
		return ctx, fmt.Errorf("calculator not found in context")
	}
	calculator.Subtract(arg2, arg1)
	return ctx, nil
}

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

func theResultShouldBeAnError(ctx context.Context) (context.Context, error) {
	if getError(ctx) == nil {
		return ctx, fmt.Errorf("expected an error, got nil")
	}
	return ctx, nil
}

func InitializeCalculatorTestSuite(_ *godog.TestSuiteContext) {}

func InitializeCalculatorScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I add (\d+) and (\d+)$`, iAddAnd)
	ctx.Step(`^I divide (\d+) by (\d+)$`, iDivideBy)
	ctx.Step(`^I have a new calculator$`, iHaveANewCalculator)
	ctx.Step(`^I multiply (\d+) by (\d+)$`, iMultiplyBy)
	ctx.Step(`^I subtract (\d+) from (\d+)$`, iSubtractFrom)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
	ctx.Step(`^the result should be an error$`, theResultShouldBeAnError)
}
