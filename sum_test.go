package main

import "testing"

func TestSum(t *testing.T){
	got := sum(4, 4)
	want := 8

	if got != want {
		t.Errorf("Abc(\"4\",\"4\") =%q; want %q", got, want)
	}
}