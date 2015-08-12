/**

 */
package main

import (
	"fmt"
)

func main(){
	fib := []int{2,0}
	i := 0
	summed := 0;

	for fib[i] < 4000000 {
		summed += fib[i];
		i = (i + 1) % 2;
		fib[i] = 4 * fib[(i + 1) % 2] + fib[i];
	}

	fmt.Println(summed)

}

func f(i,summed int, fib []int) (int) {
	return summed
}

