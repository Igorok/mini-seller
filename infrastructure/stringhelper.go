package infrastructure

import (
	"crypto/sha1"
	"fmt"
	"regexp"
)

// ReplaceWhitespace - replace whitespace by selected string like _
func ReplaceWhitespace(str, r string) string {
	space := regexp.MustCompile(`\s+`)
	str = space.ReplaceAllString(str, r)
	return str
}

// HashString - hash string by sha1
func HashString(rawString, hashSalt string) string {
	hash := sha1.New()
	hash.Write([]byte(rawString))
	hash.Write([]byte(hashSalt))
	rawString = fmt.Sprintf("%x", hash.Sum(nil))

	return rawString
}
