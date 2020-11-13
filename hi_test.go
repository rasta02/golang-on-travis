package main

import "testing"

func TestHi(t *testing.T) {
	got := Hi()
	want := "hi HELLCOME"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
