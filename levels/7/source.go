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
  "fmt"
  "math"
)

func main() {
      divisorMax := 20
      math.Log(0)
      //implement generatePrimes
      p := generatePrimes(divisorMax)
      result := float64(1)

      // EDITABLE OMIT
      for i := 0; i < len(p); i++ {
        // Your code here
        // use math.Log &  math.Pow
      }

      fmt.Println(int64(result))
}


func generatePrimes(upperLimit int) (primes []int) {
  // Your code here
  return primes
}
// UNEDITABLE OMIT


// END OMIT
