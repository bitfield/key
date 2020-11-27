package key

import "math"

const TypeableCharacters = 100

func Space(k string) uint64 {
	return uint64(math.Pow(TypeableCharacters, float64(len(k))))
}

func MedianCrackTime(k string, tryRate uint64) uint64 {
	return Space(k) / 2 / tryRate
}

func StrongEnough(k string, tryRate uint64, minCrackTime uint64) bool {
	return MedianCrackTime(k, tryRate) >= minCrackTime
}