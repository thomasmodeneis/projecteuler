package main

import (
	"math"
	"testing"
)

func TestLevel5(t *testing.T) {
	divisorMax := 20
	math.Log(0)
	//implement generatePrimes
	p := xgeneratePrimes(divisorMax)
	result := float64(1)
	for i := 0; i < len(p); i++ {
		a := math.Floor(math.Log(float64(divisorMax)) / math.Log(float64(p[i])))
		result = result * (math.Pow(float64(p[i]), a))

	}
	if 232792560 != int64(result) {
		t.Errorf("Error: the smallest positive number is not %d", int64(result))
	}

}

func xgeneratePrimes(upperLimit int) (primes []int) {

	isPrime := false
	j := 0
	primes = append(primes, 2)
	for i := 3; i <= upperLimit; i += 2 {
		j = 0
		isPrime = true
		for primes[j] * primes[j] <= i {
			if i % primes[j] == 0 {
				isPrime = false
				break
			}
			j++
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}
