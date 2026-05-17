package main

import "testing"

func TestHello(t * testing.T) {
	got := Hello()
	want := "Hello World by Go By Example website"

	if got != want {
		t.Errorf("got %q want %q", got, want)
		//error
	}
}