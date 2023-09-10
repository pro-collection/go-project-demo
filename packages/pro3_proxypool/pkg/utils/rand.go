package utils

import (
	"math/rand"
)

func RandInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}

	x := rand.Intn(max-min+1) + min

	return x
}
