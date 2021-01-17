package utils

import (
	"Reactive-Welfare-Housing-System/src/storage"
	"sync"
)

type Index struct {
	Idx int
	Mu  sync.Mutex
}

func RemoveReside(s []storage.Reside, i int) ([]storage.Reside, storage.Reside) {
	val := s[i]
	s[i] = s[len(s)-1]
	return s[:len(s)-1], val
}

type Resides []storage.Reside

func (r Resides) Len() int {
	return len(r)
}

func (r Resides) Less(i, j int) bool {
	return r[i].HouseID < r[j].HouseID
}

func (r Resides) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
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
