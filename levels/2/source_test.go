package main

import (
	"testing"
)

func TestLevel2(t *testing.T) {
	sum := 0;
	fib := []int{2,0}
	i := 0

	for fib[i] < 4000000 {
		sum += fib[i];
		i = (i + 1) % 2;
		fib[i] = 4 * fib[(i + 1) % 2] + fib[i];
	}

	if sum != 4613732 {
		t.Errorf("Error: Wrong fib found -> %d",fib)
	}
}