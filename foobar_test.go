package foobar

import "testing"

func TestDemo(t *testing.T) {
	want := "foobar"
	if s := Demo(); s != want {
		t.Errorf("Wrong result: got %s, want %s", s, want)
	}
}
