// START OMIT
/***********************
 * Mission: Largest palindrome product *
 ***********************
 *
 * A palindromic number reads the same both ways.
 * The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 *
 */
package main

import (
"fmt"
)

func main() {
  h := 0
  for i := 999; i>99; i-- {
    for j := 100; j < 1000; j++ {
      // EDITABLE OMIT

      // Your code

    }
  }
  fmt.Printf("The largest palindrome of 3 digits is: %d",h)
}

func reverse(n int) (r int) {
  for {

    // Your code

    // UNEDITABLE OMIT
  }
  return r
}

func isPalindrome(r,v int) bool{
  return r == v
}

// END OMIT
