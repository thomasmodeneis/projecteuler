// START OMIT
/***********************
 * Mission: Find the 10.001st prime number *
 ***********************
 *
 * By listing the first six prime numbers:
 * 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 *
 * Find the 10.001st prime number.
 *
 *
 *
 */
package main

import (
	"testing"
)

func TestLevel6(t *testing.T) {
	prime := xgeneratePrimes(0)
	if 104743 != prime {
		t.Errorf("The 10.001st prime number is not: %d \n",prime)
	}else {
		t.Logf("The 10.001st prime number is: %d \n", prime)
	}
}

func xgeneratePrimes(k int) (prime int) {
	count, prime := 2, 3
	arr := make([]bool, 105000)
	arr[0], arr[1] = true, true
	// Your code here
	for {
		for k = 2 * prime; k < len(arr); k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < len(arr) && arr[k]; k += 2 {
		}
		if k < len(arr) {
			prime = k
			count++
			if count == 10001 {
				break
			}
		} else {
			break
		}
	}
	return prime
}
// UNEDITABLE OMIT


// END OMIT
