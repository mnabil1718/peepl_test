package main

import (
	"crypto/sha1"
	"fmt"
)

func hashSHA1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
