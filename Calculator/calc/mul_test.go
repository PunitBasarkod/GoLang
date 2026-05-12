package calc_test

import (
	"testing"
	"github.com/PunitBasarkod/Calculator/calc"
)

func TestMul(t *testing.T){
	got := calc.Mul(3,3)
	want := 9

	if got != want {
		t.Errorf("Sub(\"a\",\"b\") =%q; want %q", got, want)
	}
}
