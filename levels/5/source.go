// START OMIT
/***********************
 * Mission: Sum square difference *
 ***********************
 *
 * The sum of the squares of the first ten natural numbers is,
 *
 * 12 + 22 + ... + 102 = 385
 * The square of the sum of the first ten natural numbers is,
 *
 * (1 + 2 + ... + 10)2 = 552 = 3025
 * Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
 *
 * Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
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
