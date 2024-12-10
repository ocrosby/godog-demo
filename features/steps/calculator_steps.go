package steps

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/ocrosby/godog-demo/pkg"
)

const calculatorKey string = "calculator"

func withCalculator(ctx context.Context, calculator *pkg.Calculator) context.Context {
	return context.WithValue(ctx, calculatorKey, calculator)
}

func getCalculator(ctx context.Context) *pkg.Calculator {
	return ctx.Value(calculatorKey).(*pkg.Calculator)
}

func iAddAnd(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator := getCalculator(ctx)
	if calculator == nil {
		return ctx, fmt.Errorf("calculator not found in context")
	}

	calculator.Add(arg1, arg2)

	return ctx, nil
}

func iDivideBy(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator := getCalculator(ctx)
	if calculator == nil {
		return ctx, fmt.Errorf("calculator not found in context")
	}

	_, err := calculator.Divide(arg1, arg2)
	if err != nil {
		return withError(ctx, err), nil
	}

	return ctx, nil
}

func iHaveANewCalculator(ctx context.Context) (context.Context, error) {
	calculator := pkg.NewCalculator()
	return withCalculator(ctx, calculator), nil
}

func iMultiplyBy(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator := getCalculator(ctx)
	if calculator == nil {
		return ctx, fmt.Errorf("calculator not found in context")
	}

	calculator.Multiply(arg1, arg2)

	return ctx, nil
}

func iSubtractFrom(ctx context.Context, arg1, arg2 int) (context.Context, error) {
	calculator := getCalculator(ctx)
	if calculator == nil {
		return ctx, fmt.Errorf("calculator not found in context")
	}

	calculator.Subtract(arg2, arg1)

	return ctx, nil
}

func theResultShouldBe(ctx context.Context, arg1 int) (context.Context, error) {
	calculator := getCalculator(ctx)
	if calculator == nil {
		return ctx, fmt.Errorf("calculator not found in context")
	}

	actual := calculator.GetAccumulator()

	if arg1 != actual {
		return ctx, fmt.Errorf("expected result %d, but got %d", arg1, actual)
	}

	return ctx, nil
}

func theResultShouldBeAnError(ctx context.Context) (context.Context, error) {
	err := getError(ctx)
	if err == nil {
		return ctx, fmt.Errorf("expected an error, but got nil")
	}

	return ctx, nil
}

func InitializeCalculatorTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// This code will run before the test suite starts
	})

	ctx.AfterSuite(func() {
		// This code will run after the test suite finishes
	})
}

func InitializeCalculatorScenario(ctx *godog.ScenarioContext) {
	//calculator := pkg.NewCalculator()

	ctx.Step(`^I add (\d+) and (\d+)$`, iAddAnd)
	ctx.Step(`^I divide (\d+) by (\d+)$`, iDivideBy)
	ctx.Step(`^I have a new calculator$`, iHaveANewCalculator)
	ctx.Step(`^I multiply (\d+) by (\d+)$`, iMultiplyBy)
	ctx.Step(`^I subtract (\d+) from (\d+)$`, iSubtractFrom)
	ctx.Step(`^the result should be (\d+)$`, theResultShouldBe)
	ctx.Step(`^the result should be an error$`, theResultShouldBeAnError)
}
