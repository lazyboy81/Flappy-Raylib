package util

import "math/rand/v2"

func RandRangeFloat64(min, max int) float64 {
	return float64(rand.IntN(max-min) + min)
}

func RandRangeFloat32(min, max int) float32 {
	return float32(rand.IntN(max-min) + min)
}

func RandRangeInt(min, max int) int {
	return rand.IntN(max-min) + min
}
