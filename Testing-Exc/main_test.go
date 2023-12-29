package main

import "testing"

func TestSum(t *testing.T) {
	got := add(2, 5)
	want := 7
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want:%d.", got, want)
	}
}

func TestSub(t *testing.T) {
	got := sub(5, 2)
	want := 3
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want:%d.", got, want)
	}
}

func TestDoMath(t *testing.T) {
	got := doMath(5, 2, sub)
	want := 3
	if got != want {
		t.Errorf("Sum was incorrect, got: %d, want:%d.", got, want)
	}

}
