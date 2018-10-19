package utils

import "math/rand"

func RandomIntN(rng *rand.Rand, n int) int {
	if rng != nil {
		return rng.Intn(n)
	} else {
		return rand.Intn(n)
	}
}
