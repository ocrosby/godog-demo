package validation_test

import (
	"net/http"
	"testing"

	"github.com/ocrosby/godog-demo/pkg/validation"
)

func fakeResponse(statusCode int) *http.Response {
	return &http.Response{StatusCode: statusCode}
}

func TestSuccessValidator(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		wantErr    bool
	}{
		{name: "200 OK", statusCode: http.StatusOK, wantErr: false},
		{name: "201 Created", statusCode: http.StatusCreated, wantErr: false},
		{name: "204 No Content", statusCode: http.StatusNoContent, wantErr: false},
		{name: "299 boundary", statusCode: 299, wantErr: false},
		{name: "300 Multiple Choices", statusCode: http.StatusMultipleChoices, wantErr: true},
		{name: "400 Bad Request", statusCode: http.StatusBadRequest, wantErr: true},
		{name: "404 Not Found", statusCode: http.StatusNotFound, wantErr: true},
		{name: "500 Internal Server Error", statusCode: http.StatusInternalServerError, wantErr: true},
		{name: "199 below range", statusCode: 199, wantErr: true},
	}

	v := validation.SuccessValidator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(fakeResponse(tt.statusCode))
			if (err != nil) != tt.wantErr {
				t.Errorf("SuccessValidator.Validate(%d) error = %v, wantErr %v", tt.statusCode, err, tt.wantErr)
			}
		})
	}
}

func TestFailureValidator(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		wantErr    bool
	}{
		{name: "400 Bad Request", statusCode: http.StatusBadRequest, wantErr: false},
		{name: "404 Not Found", statusCode: http.StatusNotFound, wantErr: false},
		{name: "500 Internal Server Error", statusCode: http.StatusInternalServerError, wantErr: false},
		{name: "199 below 2xx", statusCode: 199, wantErr: false},
		{name: "200 OK", statusCode: http.StatusOK, wantErr: true},
		{name: "201 Created", statusCode: http.StatusCreated, wantErr: true},
		{name: "204 No Content", statusCode: http.StatusNoContent, wantErr: true},
		{name: "299 boundary", statusCode: 299, wantErr: true},
	}

	v := validation.FailureValidator{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(fakeResponse(tt.statusCode))
			if (err != nil) != tt.wantErr {
				t.Errorf("FailureValidator.Validate(%d) error = %v, wantErr %v", tt.statusCode, err, tt.wantErr)
			}
		})
	}
}

func TestExactStatusValidator(t *testing.T) {
	tests := []struct {
		name       string
		expected   int
		statusCode int
		wantErr    bool
	}{
		{name: "exact match 200", expected: 200, statusCode: 200, wantErr: false},
		{name: "exact match 201", expected: 201, statusCode: 201, wantErr: false},
		{name: "exact match 404", expected: 404, statusCode: 404, wantErr: false},
		{name: "mismatch 200 vs 201", expected: 200, statusCode: 201, wantErr: true},
		{name: "mismatch 404 vs 200", expected: 404, statusCode: 200, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := validation.ExactStatusValidator{Expected: tt.expected}
			err := v.Validate(fakeResponse(tt.statusCode))
			if (err != nil) != tt.wantErr {
				t.Errorf("ExactStatusValidator{%d}.Validate(%d) error = %v, wantErr %v",
					tt.expected, tt.statusCode, err, tt.wantErr)
			}
		})
	}
}

func TestResponseValidatorInterface(t *testing.T) {
	// Verify all three types satisfy the interface at compile time.
	var _ validation.ResponseValidator = validation.SuccessValidator{}
	var _ validation.ResponseValidator = validation.FailureValidator{}
	var _ validation.ResponseValidator = validation.ExactStatusValidator{}
}
