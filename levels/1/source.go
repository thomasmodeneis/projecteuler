// START OMIT
/***********************
 * Mission: Multiples of 3 and 5 *
 ***********************
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 * Find the sum of all the multiples of 3 or 5 below 1000.
 *
 */
package main

import (
"fmt"
)

func main() {
  fib := []int{2,0}
  i := 0
  summed := 0;

  total := f(i,summed,fib)

  fmt.Println(total)
}

func f(i,summed int, fib []int) (int) {
  // EDITABLE OMIT

  // Your code

  // UNEDITABLE OMIT
  return summed
}
// END OMIT



