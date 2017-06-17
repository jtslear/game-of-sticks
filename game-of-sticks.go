package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func validateStickCount(input int, minimum int, maximum int) bool {
	if input >= minimum && input <= maximum {
		return true
	} else {
		fmt.Printf("Please enter a number between %d and %d\n", minimum, maximum)
		return false
	}
}

func startGame() int {
	var stickCount int

	fmt.Printf("*************************************************\n")
	fmt.Printf("*\tWelcome to the game of  sticks!\t\t*\n")
	fmt.Printf("*************************************************\n\n\n")

	for !validateStickCount(stickCount, 10, 100) {
		fmt.Printf("Number of sticks initially (10-100)? ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		stickCount, _ = strconv.Atoi(strings.TrimSpace(input))
	}
	return stickCount
}

func swapPlayer(input []string) []string {
	var output []string
	output = append(output, input[1])
	output = append(output, input[0])
	return output
}

func playAFriend(stickCount int) {
	fmt.Printf("*************************************************\n")
	fmt.Printf("*\tPlaying against a Friend\t\t\t*\n")
	fmt.Printf("*************************************************\n\n\n")
	players := []string{"Player1", "Player2"}
	for stickCount > 0 {
		player := players[0]
		grabCount := 0
		fmt.Printf("--> There are %d sticks on the board.\n", stickCount)
		for !validateStickCount(grabCount, 1, 3) {
			fmt.Printf("%s How many sticks do you take (1-3)? ", player)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			grabCount, _ = strconv.Atoi(strings.TrimSpace(input))
		}
		stickCount = stickCount - grabCount
		fmt.Printf("\n")
		players = swapPlayer(players)
	}
	fmt.Println(stickCount)
}

func playAComputer(stickCount int) {
	fmt.Println(stickCount)
}

func chooseGamePlay(stickCount int) {
	fmt.Printf("*************************************************\n")
	fmt.Printf("*\tGame Modes\t\t\t\t*\n")
	fmt.Printf("*************************************************\n\n\n")
	fmt.Printf("1. Play against a Friend\n")
	fmt.Printf("2. Play against an unintelligent Computer\n")
	fmt.Printf("Pick one: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)
	switch choice {
	case "1":
		playAFriend(stickCount)
	case "2":
		playAComputer(stickCount)
	default:
		fmt.Printf("%s Is not a valid choice try again\n", choice)
		chooseGamePlay(stickCount)
	}
}

func main() {
	var stickCount int
	stickCount = startGame()
	chooseGamePlay(stickCount)
}
