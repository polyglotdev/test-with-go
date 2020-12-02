package math

import "testing"

func TestMath(t *testing.T) {
	want := 7
	got := 7
	if got != want {
		t.Errorf("got: %q; want: %q", got, want)
	}
}
