package calc_test

import (
	"testing"
	"github.com/PunitBasarkod/Calculator/calc"
)

func TestSub(t *testing.T){
	got := calc.Sub(3,3)
	want := 0

	if got != want {
		t.Errorf("Sub(\"a\",\"b\") =%q; want %q", got, want)
	}
}
