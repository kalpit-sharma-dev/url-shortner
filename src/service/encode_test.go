package service

import (
	"testing"
)

func TestRandStringBytesRmndr(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		want := "O"
		got := randStringBytesRmndr(1)

		if want != got {
			t.Errorf("Error Occured Got = %v, Want = %v", got, want)
		}
	})
}
