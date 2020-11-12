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

//main func runs game logic
func main() {
	players := []string{createPlayer(), createPlayer()}

	player1 := players[0]
	player2 := players[1]

	intro(player1, player2)

	tiles := [9]string{"0", "1", "2", "3", "4", "5", "6", "7", "8"}

	drawBoard(tiles)
}
