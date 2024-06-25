package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/killuox/nhl-sim/data"
)

func main() {
	if len(os.Args) != 5 {
		fmt.Println("Usage: go run main.go sim <team1> vs <team2>")
		return
	}

	command := os.Args[1]
	operator := os.Args[3]

	if command != "sim" || operator != "vs" {
		fmt.Println("Unknown command. Usage: go run main.go sim <team1> vs <team2>")
		return
	}

	team1 := strings.ToLower(os.Args[2])
	team2 := strings.ToLower(os.Args[4])

	if team1 == team2 {
		fmt.Println("A team cannot play against each other")
		return
	}

	team1Data := findTeam(team1)
	team2Data := findTeam(team2)

	// Validate if the teams provided exist in the data
	if team1Data != nil && team2Data != nil {
		startGame(team1Data, team1Data)
	} else {
		fmt.Println("Please provide a valid NHL Team")
	}

}

func findTeam(key string) *data.Team {
	for _, team := range data.Teams {
		if key == team.Key {
			return &team
		}
	}

	return nil
}

func startGame(team1, team2 *data.Team) {
	fmt.Println("The game is starting!")

}
