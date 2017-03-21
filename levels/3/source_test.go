package main

import (
	"testing"
)

func TestLevel3(t *testing.T) {

	var numm = int64(600851475143)
	var largestFact int64
	factors := [2]int64{};

	for i := int64(2); i*i < numm; i++ {
		if numm%i == 0 { // It is a divisor
			factors[0] = i;
			factors[1] = numm/i;

			for k := int64(0); k < 2; k++ {
				isPrime := true
				for j := int64(2); j*j < factors[k]; j++ {
					if factors[k]%j == 0 {
						isPrime = false
						break
					}
				}
				if isPrime && int64(factors[k]) > largestFact {
					largestFact = int64(factors[k])
				}
			}
		}
	}
	if largestFact != 6857 {
		t.Errorf("Error: The largest primefactor of %d should not be: %d",numm,largestFact)
	}
}
