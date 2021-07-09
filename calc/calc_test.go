package main

import "testing"

func TestCalc(t *testing.T) {
	var want float64 = 5
	got := calc(2, 3, "+")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	want = 1
	got = calc(3, 2, "-")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	want = 2
	got = calc(4, 2, "/")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	want = 6
	got = calc(3, 2, "*")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
	want = 9
	got = calc(3, 2, "^")
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
