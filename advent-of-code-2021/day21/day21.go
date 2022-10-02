package day21

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	pattern := regexp.MustCompile(`\d+$`)
	lines := strings.Split(input, "\n")
	positions := make([]int, len(lines))

	for i, line := range lines {
		val, err := strconv.Atoi(pattern.FindString(line))

		if err != nil {
			panic(err)
		}

		positions[i] = val
	}

	return positions
}

type Dice interface {
	Roll() int
}

type DeterministicDice struct {
	value int
}

func (d *DeterministicDice) Roll() int {
	d.value++

	if d.value > 100 {
		d.value = 1
	}

	return d.value
}

func InitDeterministicDice() DeterministicDice {
	return DeterministicDice{0}
}

func nextPosition(position, result int) int {
	return (position+result-1)%10 + 1
}

type Game struct {
	scores, positions []int
	player, limit     int
}

func InitGame(startingPositions []int, limit int) *Game {
	numPlayers := len(startingPositions)
	scores := make([]int, numPlayers)
	positions := make([]int, numPlayers)
	copy(positions, startingPositions)

	return &Game{scores: scores, positions: positions, player: 0, limit: limit}
}

func (g *Game) RegisterResult(result int) {
	g.positions[g.player] = nextPosition(g.positions[g.player], result)
	g.scores[g.player] += g.positions[g.player]
	g.player = (g.player + 1) % 2
}

func (g *Game) PlayRound(d Dice) {
	result := 0

	for i := 0; i < 3; i++ {
		result += d.Roll()
	}

	g.RegisterResult(result)
}

func (g *Game) Copy() *Game {
	positions := make([]int, len(g.positions))
	copy(positions, g.positions)
	scores := make([]int, len(g.scores))
	copy(scores, g.scores)

	return &Game{positions: positions, scores: scores, player: g.player, limit: g.limit}
}

func (g *Game) Winner() int {
	for i, score := range g.scores {
		if score >= g.limit {
			return i
		}
	}

	return -1
}

func (g *Game) GetScore(player int) int {
	return g.scores[player]
}

func simulateDeterministicGame(startingPositions []int) int {
	game := InitGame(startingPositions, 1000)
	rolls := 0
	dice := InitDeterministicDice()

	for game.Winner() == -1 {
		game.PlayRound(&dice)
		rolls += 3
	}

	return rolls * game.GetScore(1-game.Winner())
}

func generateAggregatedResults() map[int]int {
	agg := make(map[int]int)

	for r1 := 1; r1 <= 3; r1++ {
		for r2 := 1; r2 <= 3; r2++ {
			for r3 := 1; r3 <= 3; r3++ {
				agg[r1+r2+r3]++
			}
		}
	}

	return agg
}

func simulateDiracGameRec(game *Game, winningUniverseCount *[]int, aggResults *map[int]int, universeAgg int) {
	if winner := game.Winner(); winner >= 0 {
		(*winningUniverseCount)[winner] += universeAgg
		return
	}

	for result, universeCount := range *aggResults {
		cp := game.Copy()
		cp.RegisterResult(result)
		simulateDiracGameRec(cp, winningUniverseCount, aggResults, universeAgg*universeCount)
	}
}

func simulateDiracGame(startingPositions []int) int {
	winningUniverseCount := make([]int, 2)
	game := InitGame(startingPositions, 21)
	aggResults := generateAggregatedResults()

	simulateDiracGameRec(game, &winningUniverseCount, &aggResults, 1)

	if winningUniverseCount[0] > winningUniverseCount[1] {
		return winningUniverseCount[0]
	}

	return winningUniverseCount[1]
}

func Run(input string) {
	startingPositions := parseInput(input)

	fmt.Println("Deterministic game:", simulateDeterministicGame(startingPositions))
	fmt.Println("Quantum game:", simulateDiracGame(startingPositions))
}
