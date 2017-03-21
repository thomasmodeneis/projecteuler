// START OMIT
package main

import (
	"testing"
)

type testpair struct {
	limit    int
	expected int
}

var tests = []testpair{
	{100, 25164150},
}

func TestLevel6(t *testing.T) {

	for _, test := range tests {
		diff := xSumSquareDiff(test.limit)
		if diff != test.expected {
			t.Errorf("Error: SumSquareDiff(%d) = %d, WANT: %d\n", test.limit, diff, test.expected)
		} else {
			t.Logf("Good job, Sum Square %d is equals to Diff %d\n", test.limit, test.expected)
		}
	}


}

func xsumSquare(limit int) (sum int) {
	// Your code here
	// EDITABLE OMIT
	for i := 1; i <= limit; i++ {
		sum = sum + (i * i)
	}
	// END OMIT
	return sum
}

func xsquareSum(limit int) (sum int) {
	// Your code here
	// EDITABLE OMIT
	for i := 1; i <= limit; i++ {
		sum = sum + i
	}
	// END OMIT
	return sum * sum
}

// UNEDITABLE OMIT

// SumSquareDiff
func xSumSquareDiff(limit int) int {
	return xsquareSum(limit) - xsumSquare(limit)
}

// END OMIT
