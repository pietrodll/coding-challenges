package day21

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInput(t *testing.T) {
	parsed := parseInput(`Player 1 starting position: 4
Player 2 starting position: 8`)

	assert.Equal(t, []int{4, 8}, parsed)
}

func TestDeterministicGame(t *testing.T) {
	dice := InitDeterministicDice()
	game := InitGame([]int{4, 8}, 1000)

	game.PlayRound(&dice)
	assert.Equal(t, 10, game.positions[0])
	assert.Equal(t, 10, game.scores[0])

	game.PlayRound(&dice)
	assert.Equal(t, 3, game.positions[1])
	assert.Equal(t, 3, game.scores[1])

	game.PlayRound(&dice)
	assert.Equal(t, 4, game.positions[0])
	assert.Equal(t, 14, game.scores[0])

	game.PlayRound(&dice)
	assert.Equal(t, 6, game.positions[1])
	assert.Equal(t, 9, game.scores[1])
}

func TestSimulateDeterministicGame(t *testing.T) {
	assert.Equal(t, 739785, simulateDeterministicGame([]int{4, 8}))
}

func TestSimulateDiracGame(t *testing.T) {
	assert.Equal(t, 444356092776315, simulateDiracGame([]int{4, 8}))
}
