package service

import (
	"crypto/md5"
	"math/rand"
	"strconv"

	"github.com/kalpit-sharma-dev/url-shortner/src/constants"
)

func randStringBytesRmndr(n int, data string) string {
	b := make([]byte, n)
	v := md5.Sum([]byte(data))
	randomNumber := strconv.Itoa(rand.Intn(100))
	for i := range b {
		b[i] = constants.LetterBytes[rand.Int63()%int64(len(constants.LetterBytes))]
	}
	return string(b) + string(v[1]) + randomNumber
}
