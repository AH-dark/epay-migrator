package utils

import "math/rand"

// RandString generates a random string of length n
func RandString(n int) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(b)
}

// RandInt generates a random integer between min and max
func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}
