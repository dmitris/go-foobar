package foobar

import "testing"

func TestFoobar(t *testing.T) {
	want := "foobar"
	if s := Foobar(); s != want {
		t.Errorf("Wrong result: got %s, want %s", s, want)
	}
}
