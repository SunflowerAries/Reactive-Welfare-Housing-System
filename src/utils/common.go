package utils

import (
	"sync"
)

type Index struct {
	Idx int
	Mu  sync.Mutex
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
