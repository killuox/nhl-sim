package main

import (
	"fmt"
	"math/rand"
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
		startGame(team1Data, team2Data)
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

func getRandomGoal(goalsOdds map[float64]float64, off float64, def float64) float64 {
	var totalOdds float64
	adversaryDiff := off - def

	for _, odds := range goalsOdds {
		totalOdds += odds
	}
	random := (rand.Float64() * totalOdds) + adversaryDiff

	var cumulative float64

	for goal, odds := range goalsOdds {
		cumulative += odds
		if random < cumulative {
			return goal
		}
	}

	return 0 // Fallback in case something goes wrong
}

func decideTieBreaker(team1, team2 data.Team) (float64, float64) {
	team1Chances := team1.Rank.Ovr - team2.Rank.Ovr
	team2Chances := team2.Rank.Ovr - team1.Rank.Ovr

	// Decide the winner with a random number
	team1Random := rand.Float64() * team1Chances
	team2Random := rand.Float64() * team2Chances

	if team1Random > team2Random {
		return 1, 0
	} else {
		return 0, 1

	}
}

func startGame(team1, team2 *data.Team) {
	fmt.Println("The game is starting!")
	isOvertime := false
	team1Score := getRandomGoal(data.GoalsOdds, team1.Rank.Off, team2.Rank.Def)
	team2Score := getRandomGoal(data.GoalsOdds, team2.Rank.Off, team1.Rank.Def)

	if team1Score == team2Score {
		team1Result, team2Result := decideTieBreaker(*team1, *team2)

		team1Score += team1Result
		team2Score += team2Result

		isOvertime = true
	}

	fmt.Println("---------------")
	fmt.Printf("%s: %v\n", team1.Name, team1Score)
	fmt.Printf("%s: %v\n", team2.Name, team2Score)
	if isOvertime {
		fmt.Println("OT")
	}
	fmt.Println("---------------")
}
