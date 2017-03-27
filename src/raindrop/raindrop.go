/* Convert a number to a string, the contents of which depend on the number's factors.

If the number has 3 as a factor, output 'Pling'.
If the number has 5 as a factor, output 'Plang'.
If the number has 7 as a factor, output 'Plong'.
If the number does not have 3, 5, or 7 as a factor, just pass the number's digits straight through.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	arg := int(os.Args[1])
	s := ""
	for i := 1; i <= arg; i++ {
		if arg%i == 0 {
			if i == 3 {
				s += "Pling "
			}
			if i == 7 {
				s += "Plang "
			}
			if i == 9 {
				s += "Plong "
			}
		}
	}
	if s == "" {
		fmt.Println("The raindrop string for the number input is :", arg)
		exit()
	}
	fmt.Println("The raindrop string for the number input is :", s)
}
