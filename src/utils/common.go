package utils

import (
	"Reactive-Welfare-Housing-System/src/storage"
)

func RemoveReside(s []storage.Reside, i int) ([]storage.Reside, storage.Reside) {
	val := s[i]
	s[i] = s[len(s)-1]
	return s[:len(s)-1], val
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
