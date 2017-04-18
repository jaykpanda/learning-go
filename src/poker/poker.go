// The program is to decide on the best poker hand.
// Author - Jay Panda

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the no of players in the game: ")
	text, _ := reader.ReadString('\n')

}
