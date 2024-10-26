package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

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
	fmt.Println("---------------")

	isOvertime := false
	team1Score := 0
	team2Score := 0

	// FIRST
	team1Score += g.simPeriod(g.Team1.Rank.Off, g.Team2.Rank.Def)
	team2Score += g.simPeriod(g.Team2.Rank.Off, g.Team1.Rank.Def)
	g.showResult(team1Score, team2Score, "First")

	// SECOND
	team1Score += g.simPeriod(g.Team1.Rank.Off, g.Team2.Rank.Def)
	team2Score += g.simPeriod(g.Team2.Rank.Off, g.Team1.Rank.Def)
	g.showResult(team1Score, team2Score, "Second")

	// THIRD
	team1Score += g.simPeriod(g.Team1.Rank.Off, g.Team2.Rank.Def)
	team2Score += g.simPeriod(g.Team2.Rank.Off, g.Team1.Rank.Def)
	g.showResult(team1Score, team2Score, "Third")

	// OVERTIME
	if team1Score == team2Score {
		team1Result, team2Result := g.overtime(g.Team1, g.Team2)

		team1Score += int(team1Result)
		team2Score += int(team2Result)

		isOvertime = true
	}
	g.endGame(team1Score, team2Score, isOvertime)
}

func (g *Game) endGame(team1Score int, team2Score int, isOvertime bool) {
	green := "\033[32m"
	red := "\033[31m"
	reset := "\033[0m"

	fmt.Println("Game ended")

	if team1Score > team2Score {
		// Team1 wins, print their score in green and Team2's in red
		fmt.Printf("%s: %s%v%s\n", g.Team1.Name, green, team1Score, reset)
		fmt.Printf("%s: %s%v%s\n", g.Team2.Name, red, team2Score, reset)
	} else if team2Score > team1Score {
		// Team2 wins, print their score in green and Team1's in red
		fmt.Printf("%s: %s%v%s\n", g.Team1.Name, red, team1Score, reset)
		fmt.Printf("%s: %s%v%s\n", g.Team2.Name, green, team2Score, reset)
	}
	if isOvertime {
		fmt.Println("OT")
	}
	fmt.Println("---------------")
}

func (g *Game) overtime(team1 data.Team, team2 data.Team) (float64, float64) {
	fmt.Println("OVERTIME!")
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

func (g *Game) showResult(team1Score int, team2Score int, period string) {
	fmt.Printf("%s period results:\n", period)
	fmt.Printf("%s: %v\n", g.Team1.Name, team1Score)
	fmt.Printf("%s: %v\n", g.Team2.Name, team2Score)
	fmt.Println("---------------")
	time.Sleep(1 * time.Second)
}

func (g *Game) simPeriod(off float64, def float64) int {
	// Calculate goals based on the teams' ranks off vs def
	return int(getRandomGoal(data.GoalsOddsPerPeriod, off, def))
}
