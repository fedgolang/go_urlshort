package shortener

import (
	"bytes"

	"golang.org/x/exp/rand"
)

func RandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result bytes.Buffer
	for i := 0; i < length; i++ {
		result.WriteByte(letterBytes[rand.Intn(len(letterBytes))])
	}
	return result.String()
}
