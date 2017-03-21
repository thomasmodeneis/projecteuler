package main

import (
	"testing"
)

func TestLevel4(t *testing.T) {
	h := 0
	for i := 999; i > 99; i-- {
		for j := 100; j < 1000; j++ {
			p := i * j
			r := xreverse(p)
			if xisPalindrome(r, p) {
				if h < p {
					h = p
				}
			}
		}
	}
	if h != 906609 {
		t.Errorf("Error: largest palindrome should not be: %d ", h)
	}

}

func xisPalindrome(r, v int) bool {
	return r == v
}

func xreverse(n int) (r int) {
	for {
		if n > 0 {
			r = r * 10 + n % 10
			n = n / 10
		} else {
			break
		}
	}
	return r
}
