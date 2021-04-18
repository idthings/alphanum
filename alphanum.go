package alphanum

import (
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

/*
   Originally taken from a great post on SO: https://stackoverflow.com/a/31832326
*/
func New(length int) string {

	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, length)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

func IsValidAlphaNum(input string, length int) bool {

	if len(input) != length {
		return false
	}

	for _, c := range input {
		if c < 48 || c > 122 { // less than 0
			return false
		} else if c > 57 && c < 65 { // between 9 and A
			return false
		} else if c > 90 && c < 97 { // between Z and a
			return false
		}
	}
	return true
}
