package util

import "testing"

func TestMin(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		want := 5
		got := Min(5, 6)
		if got != want {
			t.Errorf("Error got = %v and want = %v", got, want)
		}
	})
	t.Run("", func(t *testing.T) {
		want := 6
		got := Min(5, 6)
		if got == want {
			t.Errorf("Error got = %v and want = %v", got, want)
		}
	})
}
