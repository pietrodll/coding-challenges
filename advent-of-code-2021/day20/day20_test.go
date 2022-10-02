package day20

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `..#.#..#####.#.#.#.###.##.....###.##.#..###.####..#####..#....#..#..##..###..######.###...####..#..#####..##..#.#####...##.#.#..#.##..#.#......#.###.######.###.####...#.##.##..#..#..#####.....#.#....###..#.##......#.....#..#..#..##..#...##.######.####.####.#.#...#.......#..#.#.#...####.##.#......#..#...##.#.##..#...##.#.##..###.#......#.#.......#.#.#.####.###.##...#.....####.#..#..#.##.#....##..#.####....##...##..#...#......#.#.......#.......##..####..#...#.#.#...##..#.#..###..#####........#..####......#..#

#..#.
#....
##..#
..#..
..###`

func TestParseInput(t *testing.T) {
	_, image := parseInput(input)

	expected := [][]pixel{
		{pixel(1), pixel(0), pixel(0), pixel(1), pixel(0)},
		{pixel(1), pixel(0), pixel(0), pixel(0), pixel(0)},
		{pixel(1), pixel(1), pixel(0), pixel(0), pixel(1)},
		{pixel(0), pixel(0), pixel(1), pixel(0), pixel(0)},
		{pixel(0), pixel(0), pixel(1), pixel(1), pixel(1)},
	}

	assert.Equal(t, expected, image)
}

func TestIncreaseImage(t *testing.T) {
	_, image := parseInput(input)

	expected := [][]pixel{
		{pixel(0), pixel(0), pixel(0), pixel(0), pixel(0), pixel(0), pixel(0)},
		{pixel(0), pixel(1), pixel(0), pixel(0), pixel(1), pixel(0), pixel(0)},
		{pixel(0), pixel(1), pixel(0), pixel(0), pixel(0), pixel(0), pixel(0)},
		{pixel(0), pixel(1), pixel(1), pixel(0), pixel(0), pixel(1), pixel(0)},
		{pixel(0), pixel(0), pixel(0), pixel(1), pixel(0), pixel(0), pixel(0)},
		{pixel(0), pixel(0), pixel(0), pixel(1), pixel(1), pixel(1), pixel(0)},
		{pixel(0), pixel(0), pixel(0), pixel(0), pixel(0), pixel(0), pixel(0)},
	}

	assert.Equal(t, expected, increaseImage(image, pixel(0)))
}

func TestFindAlgoIndex(t *testing.T) {
	_, image := parseInput(input)
	assert.Equal(t, 34, findAlgoIndex(image, 2, 2))
}

func TestNextStep(t *testing.T) {
	algo, image := parseInput(input)
	image = increaseImage(image, pixel(0))

	image = nextStep(algo, image, pixel(0), pixel(0))

	_, expected := parseInput(`...

.........
..##.##..
.#..#.#..
.##.#..#.
.####..#.
..#..##..
...##..#.
....#.#..
.........`)

	assert.Equal(t, expected, image)
}

func TestApplyAlgo(t *testing.T) {
	algo, image := parseInput(input)
	image = applyAlgo(algo, image, 2)

	_, expected := parseInput(`...

...........
........#..
..#..#.#...
.#.#...###.
.#...##.#..
.#.....#.#.
..#.#####..
...#.#####.
....##.##..
.....###...
...........`)

	assert.Equal(t, expected, image)
}

func TestCountPixels(t *testing.T) {
	algo, image := parseInput(input)
	image = applyAlgo(algo, image, 2)

	assert.Equal(t, 35, countPixels(image))
}
