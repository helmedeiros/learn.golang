package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	printBoard(board)

	// The players take turns.
	board[0][0] = "X"
	printBoard(board)

	board[2][2] = "O"
	printBoard(board)

	board[1][2] = "X"
	printBoard(board)

	board[1][0] = "O"
	printBoard(board)

	board[0][2] = "X"
	printBoard(board)
}

func printBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	fmt.Println()
}
