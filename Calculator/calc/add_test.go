package calc_test

import (
	"testing"

	"github.com/PunitBasarkod/Calculator/calc"
)

func TestAdd(t *testing.T ) {
	got := calc.Add(3,3)
	want := 6

	if got != want {
		t.Errorf("Add(\"a\",\"b\") =%q; want %q", got, want)
	}
}