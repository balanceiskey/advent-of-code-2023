package main

import (
	"example.com/aoc-2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Game struct {
	Id         int
	Rounds     []RGB
	Constraint *RGB
}

type RGB struct {
	Blue  int
	Green int
	Red   int
}

func (g *Game) FitsConstraint() bool {
	c := g.Constraint
	for _, round := range g.Rounds {
		if round.Red > c.Red || round.Green > c.Green || round.Blue > c.Blue {
			return false
		}
	}

	return true
}

func (g *Game) GetMaxColor() RGB {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, r := range g.Rounds {
		if r.Red > maxRed {
			maxRed = r.Red
		}

		if r.Green > maxGreen {
			maxGreen = r.Green
		}

		if r.Blue > maxBlue {
			maxBlue = r.Blue
		}
	}

	return RGB{
		Red:   maxRed,
		Green: maxGreen,
		Blue:  maxBlue,
	}
}

func (g *Game) GetPower() int {
	maxColor := g.GetMaxColor()
	return maxColor.Red * maxColor.Green * maxColor.Blue
}

func NewRound(roundStr string) (*RGB, error) {
	divs := strings.Split(roundStr, ",")
	r := &RGB{}

	for _, d := range divs {
		countColor := strings.Split(strings.TrimSpace(d), " ")
		count, err := strconv.Atoi(countColor[0])

		if err != nil {
			return nil, err
		}

		color := countColor[1]

		if color == "red" {
			r.Red = count
		}

		if color == "blue" {
			r.Blue = count
		}

		if color == "green" {
			r.Green = count
		}
	}

	return r, nil
}

func NewGame(entry string, c *RGB) (*Game, error) {
	// Game 1: 9 red, 5 blue, 6 green; 6 red, 13 blue; 2 blue, 7 green, 5 red
	colonSpl := strings.Split(entry, ":")
	// Game 1
	gamePart := colonSpl[0]

	// 9 red, ...
	roundsPart := colonSpl[1]

	// 1
	gameId, err := strconv.Atoi(strings.Split(gamePart, " ")[1])

	if err != nil {
		return nil, err
	}

	roundStrings := strings.Split(roundsPart, ";")

	rounds := []RGB{}

	for _, roundStr := range roundStrings {
		round, _ := NewRound(roundStr)
		rounds = append(rounds, *round)
	}

	return &Game{
		Id:         gameId,
		Rounds:     rounds,
		Constraint: c,
	}, nil
}

func main() {
	var (
		lines, err = utils.ReadAndSplit("2.1.txt")
	)

	if err != nil {
		log.Fatalf("readLlines %s", err)
	}

	c1 := &RGB{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	games := []Game{}
	powers := []int{}

	for _, line := range lines {
		g, _ := NewGame(line, c1)
		if g.FitsConstraint() {
			games = append(games, *g)
		}

		powers = append(powers, g.GetPower())
	}

	idSum := 0

	for _, g := range games {
		idSum += g.Id
	}

	powersSum := 0

	for _, p := range powers {
		powersSum += p
	}

	fmt.Printf("Total sum of games that do not fit constraint: %d\n", idSum)
	fmt.Printf("Total sum of powers: %d\n", powersSum)
}
