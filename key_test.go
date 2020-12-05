package key_test

import (
	"key"
	"testing"
)

func TestSpace(t *testing.T) {
	input := "he"
	want := uint64(10_000)
	got := key.Space(input)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestMedianCrackTime(t *testing.T) {
	input := "he"
	tryRate := uint64(1)
	want := uint64(5_000)
	got := key.MedianCrackTime(input, tryRate)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestStrongEnough(t *testing.T) {
	input := "hello world more"
	tryRate := uint64(1e12)
	minCrackTime := uint64(100 * 365 * 24 * 60 * 60)
	want := false
	got := key.StrongEnough(input, tryRate, minCrackTime)
	if want != got {
		t.Errorf("want %t, got %t", want, got)
	}
}

func TestCrack(t *testing.T) {
	want := 99
	check := func(guess int) bool {
		return guess == want
	}
	nextTry := func(guess int) int {
		return guess + 1
	}
	got := key.Crack(check, nextTry)
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}