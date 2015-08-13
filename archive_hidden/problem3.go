package main

import (
	"fmt"
)

func main() {

	numm := 600851475143
	largestFact := 0
	factors := [2]int{};

	for i := 2; i*i < numm; i++ {
		if numm%i == 0 { // It is a divisor
			factors[0] = i;
			factors[1] = numm/i;

			for k := 0; k < 2; k++ {
				isPrime := true
				for j := 2; j*j < factors[k]; j++ {
					if factors[k]%j == 0 {
						isPrime = false
						break
					}
				}
				if isPrime && factors[k] > largestFact {
					largestFact = factors[k]
				}
			}
		}
	}
	fmt.Printf("The largest primefactor of %d is: %d",numm,largestFact)
}
