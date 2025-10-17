package main

import (
	"crypto/sha1"
)

func hashSHA1(s string) []byte {
	h := sha1.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}
