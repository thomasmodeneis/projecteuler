// START OMIT
/***********************
 * Mission: Smallest multiple *
 ***********************
 *
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 *
 * Tip: 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
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
