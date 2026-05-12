package prac

import "testing"

func TestAbc(t *testing.T) {
	got := Abc("Hi", "Hello")
	want := "Hi Hello"

	if got != want {
		t.Errorf("Abc(\"Hi\",\"Hello\") =%q; want %q", got, want)
	}
}