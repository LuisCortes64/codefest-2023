package main

import (
	"fmt"
)

func showBoard(board [][]int) {
	for _, row := range board {
		for _, v := range row {
			if v == 0 {
				fmt.Print("00 ")
			} else if v > 0 && v <= 9 {
				fmt.Print("0", v, " ")
			} else {
				fmt.Print(v, " ")
			}
		}
		fmt.Println()
	}
}

func createBoard(size int) [][]int {
	board := make([][]int, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board[i] = append(board[i], 0)
		}
	}

	return board
}

func main() {
	size := 0
	fmt.Print("Ingrese el tamaño del tablero >> ")
	fmt.Scanln(&size)

	board := createBoard(size)
	board = [][]int{
		{1, 2, 3},
		{8,9,4},
		{7,6,5},
	}
	showBoard(board)
}
