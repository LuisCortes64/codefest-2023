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

func getOptions(board [][]int, x, y, ampl int) [][]int {
	totalOptions := [][]int{
		{x, y},
		{x + ampl, y + ampl},
		{x + ampl, y - ampl},
		{x - ampl, y + ampl},
		{x - ampl, y - ampl},
		{x, y + ampl},
		{x, y - ampl},
		{x + ampl, y},
		{x - ampl, y},
	}

	for i := 1; i <= ampl; i++ {
		totalOptions = append(totalOptions, []int{x + ampl, y + i})
		totalOptions = append(totalOptions, []int{x + ampl, y - i})

		totalOptions = append(totalOptions, []int{x + ampl, y - i})
		totalOptions = append(totalOptions, []int{x + ampl, y - i})

		totalOptions = append(totalOptions, []int{x - ampl, y + i})
		totalOptions = append(totalOptions, []int{x + ampl, y - i})

		totalOptions = append(totalOptions, []int{x - ampl, y - i})
		totalOptions = append(totalOptions, []int{x + ampl, y - i})

    
    totalOptions = append(totalOptions, []int{x + i, y + ampl})
    totalOptions = append(totalOptions, []int{x + i, y - ampl})

    totalOptions = append(totalOptions, []int{x + i, y - ampl})
    totalOptions = append(totalOptions, []int{x + i, y - ampl})

    totalOptions = append(totalOptions, []int{x - i, y + ampl})
    totalOptions = append(totalOptions, []int{x + i, y - ampl})

    totalOptions = append(totalOptions, []int{x - i, y - ampl})
    totalOptions = append(totalOptions, []int{x + i, y - ampl})
	}

	options := make([][]int, 0, len(totalOptions))

	for _, op := range totalOptions {
		if op[0] < 0 || op[1] < 0 || op[0] >= len(board) || op[1] >= len(board[0]) {
			continue
		}
		if board[op[1]][op[0]] > 0 {
			continue
		}
		options = append(options, op)
	}
	return options
}

func shoot(board [][]int, x, y, p int) [][]int {
	if p == 0 {
		return board
	}

	board[y][x] = p
	p2 := p
	p--
	if p == 1 {
		return board
	}

	ops := [][]int{}

	for i := p; i > 0; i-- {
		ops = getOptions(board, x, y, i)
		for _, op := range ops {
			board[op[1]][op[0]] = p2 - p
		}
		p--
	}
	return board
}

func main() {
	size, x, y, p := 0, 0, 0, 0
	fmt.Print("Ingrese el tamaño del tablero >> ")
	fmt.Scanln(&size)

	fmt.Print("Ingrese la posicion en X de la flecha >> ")
	fmt.Scanln(&x)

	fmt.Print("Ingrese la posicion en Y de la flecha >> ")
	fmt.Scanln(&y)

	fmt.Print("Ingrese la Potencia de la flecha >> ")
	fmt.Scanln(&p)

	board := createBoard(size)
	shoot(board, x, y, p)
	showBoard(board)

}
