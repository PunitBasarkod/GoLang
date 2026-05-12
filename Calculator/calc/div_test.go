package calc_test

import (
	"testing"

	"github.com/PunitBasarkod/Calculator/calc"
)

func TestDiv(t *testing.T) {
	tests := []struct {
		a, b    int
		want    int
		hasErr  bool
	}{
		{10, 2, 5, false},
		{9, 3, 3, false},
		{7, 0, 0, true}, // division by zero should give error
	}

	for _, tt := range tests {
		got, err := calc.Div(tt.a, tt.b)

		if tt.hasErr {
			// Test case expects an error
			if err == nil {
				t.Errorf("Div(%d, %d) expected error, got none", tt.a, tt.b)
			}
			// Don't check 'got' value when error is expected
		} else {
			// Test case does NOT expect an error
			if err != nil {
				t.Errorf("Div(%d, %d) unexpected error: %v", tt.a, tt.b, err)
			}
			if got != tt.want {
				t.Errorf("Div(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		}
	}
}