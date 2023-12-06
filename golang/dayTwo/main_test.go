package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConstraint(t *testing.T) {
	c := RGB{
		Blue:  1,
		Red:   4,
		Green: 3,
	}
	g := Game{
		Rounds:     []RGB{{Red: 1, Green: 4, Blue: 1}},
		Constraint: &c,
	}
	assert.Equal(t, false, g.FitsConstraint())

	g = Game{
		Rounds: []RGB{
			{
				Red: 1, Green: 1, Blue: 1,
			},
			{

				Red: 2, Green: 2, Blue: 2,
			},
			{

				Red: 3, Green: 3, Blue: 3,
			},
		},
		Constraint: &RGB{
			Blue:  3,
			Red:   3,
			Green: 3,
		},
	}

	assert.Equal(t, true, g.FitsConstraint())
}

func TestNewGame(t *testing.T) {
	g, _ := NewGame("Game 1: 9 red, 5 blue, 6 green; 6 red, 13 blue; 2 blue, 7 green, 5 red\n", nil)

	assert.Equal(t, 1, g.Id)
	assert.Equal(t, 3, len(g.Rounds))
}

func TestTrialInput(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	c := &RGB{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	lines := strings.Split(input, "\n")

	games := []Game{}

	for _, line := range lines {
		g, _ := NewGame(line, c)
		if g.FitsConstraint() {
			games = append(games, *g)
		}
	}

	idSum := 0

	for _, g := range games {
		idSum += g.Id
	}

	assert.Equal(t, 8, idSum)

}

func TestMaxColor(t *testing.T) {
	g, _ := NewGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", nil)
	maxColor := g.GetMaxColor()

	assert.Equal(t, 4, maxColor.Red)
	assert.Equal(t, 2, maxColor.Green)
	assert.Equal(t, 6, maxColor.Blue)

	power := g.GetPower()

	assert.Equal(t, 48, power)

	g, _ = NewGame("Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", nil)
	maxColor = g.GetMaxColor()

	assert.Equal(t, 1, maxColor.Red)
	assert.Equal(t, 3, maxColor.Green)
	assert.Equal(t, 4, maxColor.Blue)

	power = g.GetPower()

	assert.Equal(t, 12, power)
}
