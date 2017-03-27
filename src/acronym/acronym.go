// Prints the acronym for a given input
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := ""
	for _, arg := range os.Args[1:] {
		s += string(arg[0])
	}
	s = strings.ToUpper(s)
	fmt.Println("The acronym for the input string is ", s)
}
