package main

import (
	"fmt"
)

func main() {

	board := [3][3]string{}
	fmt.Println("WELCOME TO XOX GAME")
	fmt.Println("___________________")
	fmt.Println()
	fmt.Println("GAME STARTS NOW")
	fmt.Println("------------------")
	player := "x"

	for {
		var row int
		var column int

		fmt.Println("Enter the row number")
		fmt.Scanln(&row)

		if row < 0 || row > 2 {
			fmt.Println("Invalid number")
			continue
		}

		fmt.Println("Enter the column number")
		fmt.Scanln(&column)

		if column < 0 || column > 2 {
			fmt.Println("Invalid number")
			continue
		}

		if board[row][column] == "" {
			board[row][column] = player
		} else {
			fmt.Println("Already box is filled")
			continue
		}
		fmt.Println(board[0])
		fmt.Println(board[1])
		fmt.Println(board[2])

		for i := 0; i < 3; i++ {
			if (board[0][i] == board[1][i] && board[1][i] == board[2][i] && board[2][i] == player) || board[i][0] == board[i][1] && board[i][1] == board[i][2] && board[i][2] == player {
				fmt.Println("Game over, player won", player)
				return
			}
		}
		for i := 0; i < 3; i++ {
			if (board[0][0] == board[1][1] && board[1][1] == board[2][2] && board[2][2] == player) || board[0][2] == board[1][1] && board[1][1] == board[2][0] && board[2][0] == player {
				fmt.Println("Game Over Player won", player)
				return
			}
		}
		for i := 0; i < 3; i++ {
			if board[0][0] != "" && board[0][1] != "" && board[0][2] != "" && board[1][0] != "" && board[1][1] != "" && board[1][2] != "" && board[2][0] != "" && board[2][1] != "" && board[2][2] != "" {
				fmt.Println("Nobody wins, it's DRAW")
				return
			}

			if player == "x" {
				player = "o"
			} else {
				player = "x"
			}
		}

	}
}
