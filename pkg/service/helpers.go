package service

import (
	"bytes"
	"math/rand"
)

const genArray = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const lenLink = 10

func generShortLink() string {
	var buffer bytes.Buffer
	underscore := rand.Intn(lenLink)
	lenGenArray := len([]rune(genArray))
	for i := 0; i < lenLink; i++ {
		if i == underscore {
			buffer.WriteString("_")
		} else {
			buffer.WriteString(string(genArray[rand.Intn(lenGenArray)]))
		}

	}
	return buffer.String()
}
