package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Board struct {
	grid     [][]int
	selected [][]bool
}

func parseBoard(block string) Board {
	lines := strings.Split(block, "\n")

	sep := regexp.MustCompile(`\s+`)

	grid := make([][]int, len(lines))
	selected := make([][]bool, len(lines))

	for i, line := range lines {
		strVals := sep.Split(strings.Trim(line, " "), -1)

		vals := make([]int, len(strVals))
		selectedLine := make([]bool, len(strVals))

		for j, strVal := range strVals {
			val, err := strconv.Atoi(strVal)

			if err != nil {
				panic(err)
			}

			vals[j] = val
			selectedLine[j] = false
		}

		grid[i] = vals
		selected[i] = selectedLine
	}

	return Board{grid, selected}
}

func parseInput(input string) ([]int, []Board) {
	blocks := strings.Split(input, "\n\n")

	strNumbers := strings.Split(blocks[0], ",")
	numbers := make([]int, len(strNumbers))

	for i, strVal := range strNumbers {
		val, err := strconv.Atoi(strVal)

		if err != nil {
			panic(err)
		}

		numbers[i] = val
	}

	boards := make([]Board, len(blocks)-1)

	for i, block := range blocks[1:] {
		boards[i] = parseBoard(block)
	}

	return numbers, boards
}

func (board *Board) drawNumber(number int) {
	for i, line := range board.grid {
		for j, val := range line {
			if val == number {
				board.selected[i][j] = true
			}
		}
	}
}

func (board *Board) checkWin() bool {
	// check rows
	for _, row := range board.selected {
		wins := true

		for _, valueSelected := range row {
			wins = wins && valueSelected
		}

		if wins {
			return true
		}
	}

	// check columns
	for j := 0; j < len(board.selected[0]); j++ {
		wins := true

		for i := 0; i < len(board.selected); i++ {
			wins = wins && board.selected[i][j]
		}

		if wins {
			return true
		}
	}

	return false
}

func (board *Board) sumUnmarked() int {
	sum := 0

	for i, row := range board.grid {
		for j, val := range row {
			if !board.selected[i][j] {
				sum += val
			}
		}
	}

	return sum
}

// Runs the game and returns the score of the winner. If no one wins, returns -1
func playGame(numbers []int, boards []Board) int {
	for _, number := range numbers {
		for _, board := range boards {
			board.drawNumber(number)

			if board.checkWin() {
				return number * board.sumUnmarked()
			}
		}
	}

	return -1
}

func playGameUntilLastWins(numbers []int, boards []Board) int {
	winnerScores := make([]int, len(boards))
	lastToWin := -1

	for i := range winnerScores {
		winnerScores[i] = -1
	}

	for _, number := range numbers {
		for boardIndex, board := range boards {
			board.drawNumber(number)

			if winnerScores[boardIndex] == -1 && board.checkWin() {
				winnerScores[boardIndex] = number * board.sumUnmarked()
				lastToWin = boardIndex
			}
		}
	}

	if lastToWin > -1 {
		return winnerScores[lastToWin]
	}

	return -1
}

func Run(input string) {
	numbers, boards := parseInput(input)
	winnerScore := playGame(numbers, boards)
	lastToWinScore := playGameUntilLastWins(numbers, boards)

	fmt.Println("Winner score:", winnerScore)
	fmt.Println("Last to win score:", lastToWinScore)
}
