package main

import (
	"fmt"
	"math/rand"
	"time"
)

var board = [][]string{
	{"_", "_", "_"},
	{"_", "_", "_"},
	{"_", "_", "_"},
}

func printBoard() {
	fmt.Println("Current Baord:")
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			fmt.Print(board[i][j] + " ")
		}
		fmt.Print("\n")
	}
}

func checkValidMove(row int, col int) bool {
	if board[row][col] == "_" {
		return true
	} else {
		return false
	}
}

/*
*         Row Win Example
*            _ _ _
*            X X X
*            _ _ _
 */
func isRowWin(piece string) bool {
	for i := 0; i < len(board); i++ {
		var count = 0
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == piece {
				count++
			}
		}
		if count == 3 {
			fmt.Println("row win")
			return true
		}
	}
	return false
}

/*
*      Column Win Example
*            O _ _
*            O _ _
*            O _ _
 */
func isColWin(piece string) bool {
	for i := 0; i < len(board[0]); i++ {
		var count = 0
		for j := 0; j < len(board); j++ {
			if board[j][i] == piece {
				count++
			}
		}
		if count == 3 {
			fmt.Println("col win")
			return true
		}
	}
	return false
}

/*
*Diagnal Right Win Example
*             X _ _
*             _ X _
*             _ _ X
 */
func isDagRightWin(piece string) bool {
	if board[0][0] == piece && board[1][1] == piece && board[2][2] == piece {
		fmt.Println("dag right win")
		return true
	} else {
		return false
	}
}

/*
*Diagnal Left Win Example
*            _ _ O
*            _ O _
*            O _ _
 */
func isDagLeftWin(piece string) bool {
	if board[0][2] == piece && board[1][1] == piece && board[2][0] == piece {
		fmt.Println("dag left win")
		return true
	} else {
		return false
	}
}

func isFullBoard() bool {
	var count = 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "X" || board[i][j] == "O" {
				count++
			}
		}
	}
	if count == 9 {
		fmt.Println("Tie win")
		return true
	} else {
		return false
	}
}

func isWinner() bool {
	if isRowWin("X") || isColWin("X") || isDagRightWin("X") || isDagLeftWin("X") {
		fmt.Println("You win!")
		return true
	} else if isRowWin("O") || isColWin("O") || isDagRightWin("O") || isDagLeftWin("O") {
		fmt.Println("Comp win!")
		return true
	} else if isFullBoard() {
		fmt.Println("Tie win!")
		return true
	} else {
		return false
	}
}

/*
Game assumes that you will work within the counds for now. Will not account for wrong numbers
*/
func main() {
	var user_row, user_col, comp_row, comp_col int
	var possibleMoves = 9
	var currentMoves = 0
	var isUserMoveValid, isCompMoveValid bool

	fmt.Println("Let's play Tic-Tac-Toe. \nYou are X")
	printBoard()

	for currentMoves < possibleMoves {
		isUserMoveValid = false
		isCompMoveValid = false
		for !isUserMoveValid {
			fmt.Println()
			fmt.Print("Pick a row (0-2): ")
			fmt.Scanln(&user_row)
			fmt.Print("Pick a row (0-2): ")
			fmt.Scanln(&user_col)
			fmt.Println()
			if checkValidMove(user_row, user_col) {
				board[user_row][user_col] = "X"
				isUserMoveValid = true
			}
		}
		printBoard()
		fmt.Println()
		if isWinner() {
			break
		}

		for !isCompMoveValid {
			rand.Seed(time.Now().UnixNano())
			comp_row = rand.Intn(3)
			comp_col = rand.Intn(3)
			if checkValidMove(comp_row, comp_col) {
				board[comp_row][comp_col] = "O"
				isCompMoveValid = true
			}
		}

		printBoard()
		fmt.Println("")
		if isWinner() {
			break
		}
		currentMoves++
	}
}
