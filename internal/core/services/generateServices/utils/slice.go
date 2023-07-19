package utils

import "math/rand"

func ShuffleSliceOfString(s []string) {
	rand.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}
