package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	filename := "input"
	inputFile, _ := os.Open(filename)
	defer inputFile.Close()

	inputScanner := bufio.NewScanner(inputFile)

	type game struct {
		OpponentPlays string
		IPlay         string
		Result        string
		Points        int
	}

	runningTotal := 0

	var gameRecords []game
	// Each line of file
	for inputScanner.Scan() {
		var currentGame game
		game := strings.Split(inputScanner.Text(), " ")
		currentGame.Points = 0
		//A for Rock, B for Paper, and C for Scissors.
		switch game[0] {
		case "A":
			currentGame.OpponentPlays = "Rock"

		case "B":
			currentGame.OpponentPlays = "Paper"

		case "C":
			currentGame.OpponentPlays = "Scissors"

		}
		//X for Rock, Y for Paper, and Z for Scissors
		switch game[1] {
		case "X":
			currentGame.IPlay = "Rock"
			currentGame.Points += 1
		case "Y":
			currentGame.IPlay = "Paper"
			currentGame.Points += 2
		case "Z":
			currentGame.IPlay = "Scissors"
			currentGame.Points += 3

		}

		// Get the game outcome
		currentGame.Result, currentGame.Points = findResult(currentGame.OpponentPlays, currentGame.IPlay, currentGame.Points)
		gameRecords = append(gameRecords, currentGame)
		runningTotal += currentGame.Points
	}
	fmt.Printf("%v \n", gameRecords)
	fmt.Printf("%d \n", runningTotal)

}

func findResult(opp string, my string, currentPoints int) (outcome string, addpoints int) {

	handOptions := []string{"Rock", "Paper", "Scissors"}

	// rotate array so my hand is in the middle
	for my != handOptions[1] {
		storeFirst := handOptions[0]
		handOptions = handOptions[1:]
		handOptions = append(handOptions, storeFirst)
	}

	switch {
	case indexOf(my, handOptions) > indexOf(opp, handOptions):
		outcome = "win"
		addpoints = currentPoints + 6
	case indexOf(my, handOptions) < indexOf(opp, handOptions):
		outcome = "lose"
		addpoints = currentPoints
	default:
		outcome = "draw"
		addpoints = currentPoints + 3
	}
	return outcome, addpoints

}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
