//solution with brute force version

package main

import (
	"testing"
)

func TestLevel5_BruteForce(t *testing.T) {

	counter := 0
	n := 0

	for j :=2520; counter < 20; j++ {

		for i := 1; i < 21; i++ {

			if j%i == 0 {
				n = j
				counter = counter + 1
			}else {
				counter = 0
			}
		}
	}

	if 232792560 != n {
		t.Errorf("Error: The smallest positive number that is evenly divisible by all of the numbers from 1 to 20 is not %d", n)
	}
}
