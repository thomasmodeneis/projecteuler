package main

import (
	"testing"
)

func TestLevel1(t *testing.T) {
	total := xgetMultiples(0,1000)
	if total != 233168 {
		t.Errorf("Wrong multiple found -> %d",total)
	}
}

func xgetMultiples(i, m int64) (t int64) {
	for ; i<m;i++ {
		if i % 3 == 0 || i % 5 == 0 {
			t = t + i
		}
	}
	return t
}

