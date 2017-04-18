/*Use the Sieve of Eratosthenes to find all the primes from 2 up to a given number.
The Sieve of Eratosthenes is a simple, ancient algorithm for finding all prime numbers up to any given limit.
It does so by iteratively marking as composite (i.e. not prime) the multiples of each prime, starting with the multiples of 2. */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	no, _ := strconv.Atoi(os.Args[1])
	div := 0
	sieve := make([]int, no+1)
	for i := 2; i <= no; i++ {
		sieve[i] = i
	}
	for j := 2; j <= no; j++ {
		if sieve[j] != 0 {
			for k := 2; k < j; k++ {
				if (sieve[j] % k) == 0 {
					div = div + 1
				}
			}
		}
		if div == 0 {
			for n := 2 * j; n <= no; n = n + j {
				sieve[n] = 0
			}
		}
		div = 0
	}
	fmt.Println("All the prime nos upto", no, "are:")
	for i := 2; i <= no; i++ {
		if sieve[i] != 0 {
			fmt.Print(i, " ")
		}
	}
}
