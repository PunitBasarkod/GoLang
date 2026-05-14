package calc_test

import (
	"testing"

	"github.com/PunitBasarkod/Calculator/calc"
)

func TestAdd(t *testing.T ) {
	tests := []struct {
		name	string
		a, b	int
		want	int
	}{
		{"positive numbers", 3, 3, 6},
		{"zero numbers", 0, 5, 5},
		{"negative numbers", -2, 7, 5},
		{"both negative", -4, -6, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := calc.Add(tt.a, tt.b)

			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

