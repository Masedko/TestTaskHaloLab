package utils

import "math/rand"

func GenerateIntInInterval(start int, end int) int {
	return rand.Intn(end-start+1) + start
}
