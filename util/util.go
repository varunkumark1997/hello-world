package util

import "math/rand"

func RandomPos(min, max int) int {
	return rand.Intn(max-min) + min
}
