package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func CreateMd5Hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	println(CreateMd5Hash("Hello World"))
}
