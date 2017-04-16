/*Package queen ...
Given the position of two queens on a chess board, indicate whether or not they are positioned
so that they can attack each other.
In the game of chess, a queen can attack pieces which are on the same row, column, or diagonal.
A chessboard can be represented by an 8 by 8 array.
_ _ _ _ _ _ _ _
_ _ _ _ _ _ _ _
_ _ _ ♛ _ _ _ _
_ _ _ _ _ _ _ _
_ _ _ _ _ _ _ _
_ _ _ _ _ _ ♛ _
_ _ _ _ _ _ _ _
_ _ _ _ _ _ _ _
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	whiteRow, _ := strconv.Atoi(os.Args[1])
	whiteCol, _ := strconv.Atoi(os.Args[2])
	blackRow, _ := strconv.Atoi(os.Args[3])
	blackCol, _ := strconv.Atoi(os.Args[4])
	var possible int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if (i == whiteRow && j == whiteCol) || (i == blackRow && j == blackCol) {
				fmt.Print("\u265b ")
			} else {
				fmt.Print("_ ")
			}
		}
		fmt.Println()
	}
	if whiteRow == blackRow || whiteCol == blackCol {
		possible = 1
	}
	if Abs(whiteRow-blackRow) == Abs(whiteCol-blackCol) {
		possible = 1
	}
	if possible == 1 {
		fmt.Println("The attack is possible")
	}
	if possible == 0 {
		fmt.Println("The attack is not possible")
	}
}

//Abs returns the absolute value of an integer
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
