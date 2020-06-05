package main

import (
	"fmt"
)

func initialValues() [9]string {
	var board [9]string
	board[0] = "a"
	board[1] = "b"
	board[2] = "c"
	board[3] = "d"
	board[4] = "e"
	board[5] = "f"
	board[6] = "g"
	board[7] = "h"
	board[8] = "i"
	return board
}

func ticTacToe(board [9]string) [9]string {

	fmt.Println("Lets Play TicTackToe")
	fmt.Println(board[0] + " | " + board[1] + " | " + board[2])
	fmt.Println("--|---|---")
	fmt.Println(board[3] + " | " + board[4] + " | " + board[5])
	fmt.Println("--|---|---")
	fmt.Println(board[6] + " | " + board[7] + " | " + board[8])

	return board
}

func gameStart(board [9]string) {
	var result string
	fmt.Println("For Player 1 'O' and For Player 2 'X'")
	for i := 0; i < 4; i++ {
		fmt.Println("Round", i)
		board, result = playerOne(board, i)
		if result != "" {
			fmt.Println(result)
			break
		}
		//fmt.Println(board)
		board, result = playerSecond(board, i)
		if result != "" {
			fmt.Println(result)
			break
		}
	}
}

func checkForWin(i int, player string, b [9]string) string {
	var result string
	if i == 2 {
		if (b[0] == player && b[0] == b[3] && b[6] == player) || (b[1] == player && b[1] == b[4] && b[7] == player) || (b[2] == player && b[2] == b[5] && b[8] == player) || (b[0] == player && b[0] == b[1] && b[2] == player) || (b[3] == player && b[3] == b[4] && b[5] == player) || (b[6] == player && b[6] == b[7] && b[8] == player) || (b[0] == player && b[0] == b[4] && b[8] == player) || (b[2] == player && b[2] == b[4] && b[6] == player) {
			if player == "O" {
				result = "Player 1 Won"
			} else {
				//fmt.Println(player)
				result = "Player 2 Won"
			}
		}
	}
	return result
}

func playerOne(board [9]string, i int) ([9]string, string) {
	var temp string
	var player string = "O"
	fmt.Println("Player1")
	fmt.Scanln(&temp)
	board = changingValues(board, temp, player)
	result := checkForWin(i, player, board)
	return board, result
}

func playerSecond(board [9]string, i int) ([9]string, string) {
	var temp string
	var player string = "X"
	fmt.Println("Player2")
	fmt.Scanln(&temp)
	board = changingValues(board, temp, player)
	result := checkForWin(i, player, board)
	return board, result
}

func changingValues(board [9]string, temp string, player string) [9]string {
	switch temp {
	case "a":
		board[0] = player
	case "b":
		board[1] = player
	case "c":
		board[2] = player
	case "d":
		board[3] = player
	case "e":
		board[4] = player
	case "f":
		board[5] = player
	case "g":
		board[6] = player
	case "h":
		board[7] = player
	case "i":
		board[8] = player
	}
	ticTacToe(board)
	//fmt.Println("change ", board)
	return board
}

func restart() {
	var restartVar string
	fmt.Println("Wanna Play Again [y/n]")
	fmt.Scanln(&restartVar)
	if restartVar != "" {
		if restartVar == "y" {
			main()
		} else {
			fmt.Println("See you again :)")
		}
	} else {
		fmt.Println("Enter value again")
		restart()
	}
}

func main() {
	board := initialValues()
	ticTacToe(board)
	gameStart(board)
	restart()
}
