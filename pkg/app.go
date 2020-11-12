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

//initBoard creates a blank board
func initBoard() {
	for i := 0; i <= 12; i++ {
		if i == 4 || i == 8 {
			fmt.Println(strings.Repeat("-", 25))
		} else {
			fmt.Println(strings.Repeat(" ", 6), "|", strings.Repeat(" ", 6), "|")
		}
	}
}

//main func runs game logic
func main() {
	players := []string{createPlayer(), createPlayer()}

	player1 := players[0]
	player2 := players[1]

	intro(player1, player2)

	initBoard()
}
