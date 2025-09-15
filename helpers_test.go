package main

import (
	"go-server/app"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		a              int
		b              int
		expectedResult int
	}{
		{1, 1, 2},
		{2, 2, 3},
	}

	for _, tt := range tests {
		result := app.Add(tt.a, tt.b)
		if result != tt.expectedResult {
			t.Errorf("expectedResult: %d, actualResult: %d", tt.expectedResult, result)
		}

	}
}
