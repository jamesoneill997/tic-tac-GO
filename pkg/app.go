package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//create a player from stdin
func createPlayer() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Player 1 name: ")
	playerName, _ := reader.ReadString('\n')

	return playerName
}

func intro(player1 string, player2 string) {
	fmt.Printf("Player 1: %s", player1)
	fmt.Printf("Player 2: %s", player2)

	fmt.Println("Welcome to tic-tac-GO. To play, simply enter the square number that you would like to place your mark on. Get 3 in a vertical or horizontal row to win!")
}

func drawBoard(tiles [9]string) {
	for i, j := 0, 0; i <= 13; i++ {
		if i == 4 || i == 9 {
			fmt.Println(strings.Repeat("-", 25))
		} else if i%5 == 1 {
			fmt.Println(strings.Repeat(" ", 2), tiles[j], " ", "|", strings.Repeat(" ", 2), tiles[j+1], " ", "|", strings.Repeat(" ", 2), tiles[j+2], " ")
			j += 3
		} else {
			fmt.Println(strings.Repeat(" ", 6), "|", strings.Repeat(" ", 6), "|")
		}

	}
}

func makeMove(currPlayer int, tiles [9]string) [9]string {
	var move int
	fmt.Scan(&move)

	if tiles[move] == "X" || tiles[move] == "O" {
		fmt.Println("Invalid move, space already taken. Please try again.")

		makeMove(currPlayer, tiles)
	}

	if currPlayer == 0 {
		tiles[move] = "X"
	} else {
		tiles[move] = "O"
	}

	return tiles
}

func checkWinner(tiles [9]string) bool {
	return (tiles[0] == tiles[1] && tiles[1] == tiles[2]) || (tiles[3] == tiles[4] && tiles[4] == tiles[5]) || (tiles[6] == tiles[7] && tiles[7] == tiles[8]) || (tiles[1] == tiles[3] && tiles[3] == tiles[6]) || (tiles[4] == tiles[4] && tiles[4] == tiles[7]) || (tiles[2] == tiles[5] && tiles[5] == tiles[8]) || (tiles[0] == tiles[4] && tiles[4] == tiles[8]) || (tiles[2] == tiles[4] && tiles[4] == tiles[6])
}

//main func runs game logic
func main() {
	players := []string{createPlayer(), createPlayer()}

	player1 := players[0]
	player2 := players[1]

	tiles := [9]string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}

	drawBoard(tiles)

	intro(player1, player2)

	currMove := 0

	for currMove <= 9 && !checkWinner(tiles) {
		drawBoard(tiles)
		if currMove%2 != 0 {
			tiles = makeMove(0, tiles)
			if checkWinner(tiles) {
				fmt.Println(player1 + " is the winner!")
			}
		} else {
			tiles = makeMove(1, tiles)
			if checkWinner(tiles) {
				fmt.Println(player2 + " is the winner!")
			}
		}

		currMove++
	}

}
