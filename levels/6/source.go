// START OMIT
/***********************
 * Mission: Sum square difference *
 ***********************
 *
 * Your job is to find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 * hint: The sum of the squares of the first ten natural numbers is -> 12 + 22 + ... + 102 = 385
 * The square of the sum of the first ten natural numbers is -> (1 + 2 + ... + 10)2 = 552 = 3025
 * Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640
 *
 */
package main

import (
	"fmt"
)

func main() {
	limit := 10
	expected := 2640

	diff := SumSquareDiff(limit)
	if diff != expected {
		fmt.Printf("Error: SumSquareDiff(%d) = %d, WANT: %d", limit, diff, expected)
	}else {
		fmt.Printf("Good job, Sum Square %d is equals to Diff %d", limit, expected)
	}
}

func sumSquare(limit int) (sum int) {
	// Your code here
	// EDITABLE OMIT

	// END OMIT
	return sum
}

func squareSum(limit int) (sum int) {
	// Your code here
	// EDITABLE OMIT

	// END OMIT
	return sum
}

// UNEDITABLE OMIT

// SumSquareDiff
func SumSquareDiff(limit int) int {
	return squareSum(limit) - sumSquare(limit)
}

// END OMIT
