/*Compute Pascal's triangle up to a given number of rows.

In Pascal's Triangle each number is computed by adding the numbers to the right and left
of the current position in the previous row.
    1
   1 1
  1 2 1
 1 3 3 1
1 4 6 4 1
# ... etc............*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	height, _ := strconv.Atoi(os.Args[1])
	var val [30][30]int
	val[0][1] = 1
	for i := 0; i < height; i++ {
		fmt.Print(" ")
	}
	fmt.Println(val[0][1])
	for i := 1; i < height; i++ {
		for k := 0; k < height-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i+1; j++ {
			val[i][j] = val[i-1][j-1] + val[i-1][j]
			fmt.Print(val[i][j], " ")
		}
		fmt.Println("")
	}
}
