package service

import (
	"math/rand"

	"github.com/kalpit-sharma-dev/url-shortner/src/constants"
)

func randStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = constants.LetterBytes[rand.Int63()%int64(len(constants.LetterBytes))]
	}
	return string(b)
}
