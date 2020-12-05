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

type CheckFunc func(int) bool
type NextTryFunc func (int) int

func Crack(check CheckFunc, nextTry NextTryFunc) int {
	guess := 0
	for {
		if check(guess) {
			return guess
		}
		guess = nextTry(guess)
	}
}