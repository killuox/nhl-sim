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
		game := Game{
			Team1: *team1Data,
			Team2: *team2Data,
		}
		game.start()
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

type Game struct {
	Team1 data.Team
	Team2 data.Team
}

func (g *Game) start() {
	fmt.Println("The game is starting!")
	isOvertime := false
	// Calculate goals based on the teams' ranks off vs def
	team1Score := getRandomGoal(data.GoalsOdds, g.Team1.Rank.Off, g.Team2.Rank.Def)
	team2Score := getRandomGoal(data.GoalsOdds, g.Team2.Rank.Off, g.Team1.Rank.Def)

	if team1Score == team2Score {
		team1Result, team2Result := g.overtime(g.Team1, g.Team2)

		team1Score += team1Result
		team2Score += team2Result

		isOvertime = true
	}

	fmt.Println("---------------")
	fmt.Printf("%s: %v\n", g.Team1.Name, team1Score)
	fmt.Printf("%s: %v\n", g.Team2.Name, team2Score)
	if isOvertime {
		fmt.Println("OT")
	}
	fmt.Println("---------------")
}

func (g *Game) overtime(team1 data.Team, team2 data.Team) (float64, float64) {
	team1Chances := team1.Rank.Ovr - team2.Rank.Ovr
	team2Chances := team2.Rank.Ovr - team1.Rank.Ovr

	rangeAdjustment := 100.0

	// Generate random numbers within adjusted ranges
	team1Random := rand.Float64() * (1 + float64(team1Chances)/rangeAdjustment)
	team2Random := rand.Float64() * (1 + float64(team2Chances)/rangeAdjustment)

	if team1Random > team2Random {
		return 1, 0 // Team1 wins
	} else {
		return 0, 1 // Team2 wins
	}
}
