package service

import (
	"testing"
)

func TestRandStringBytesRmndr(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		want := `N81`
		got := randStringBytesRmndr(1, "Hello")

		if want != got {
			t.Errorf("Error Occured Got = %v, Want = %v", got, want)
		}
	})
}
