package day13

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/pietrodll/aoc2021/utils/parse"
)

type Point struct {
	X int
	Y int
}

type Fold interface {
	transform(Point) Point
}

type XFold struct {
	X int
}

func (fold XFold) transform(p Point) Point {
	if p.X > fold.X {
		return Point{2*fold.X - p.X, p.Y}
	}

	return p
}

type YFold struct {
	Y int
}

func (fold YFold) transform(p Point) Point {
	if p.Y > fold.Y {
		return Point{p.X, 2*fold.Y - p.Y}
	}

	return p
}

func parseInput(input string) ([]Point, []Fold) {
	blocks := strings.Split(input, "\n\n")

	strPoints := strings.Split(blocks[0], "\n")
	points := make([]Point, len(strPoints))

	for i, strPoint := range strPoints {
		vals := parse.ParseIntegers(strPoint, ",")
		points[i] = Point{vals[0], vals[1]}
	}

	strFolds := strings.Split(blocks[1], "\n")
	folds := make([]Fold, len(strFolds))
	foldPattern := regexp.MustCompile(`^fold along (x|y)=(\d+)$`)

	for i, strFold := range strFolds {
		groups := foldPattern.FindAllStringSubmatch(strFold, -1)[0]

		if val, err := strconv.Atoi(groups[2]); err != nil {
			panic(err)
		} else if groups[1] == "x" {
			folds[i] = XFold{val}
		} else {
			folds[i] = YFold{val}
		}

	}

	return points, folds
}

func foldAndReduce(points []Point, fold Fold) []Point {
	folded := make(map[Point]bool)

	for _, point := range points {
		folded[fold.transform(point)] = true
	}

	result := make([]Point, len(folded))
	i := 0

	for point := range folded {
		result[i] = point
		i++
	}

	return result
}

func executeFolds(points []Point, folds []Fold) []Point {
	result := points

	for _, fold := range folds {
		result = foldAndReduce(result, fold)
	}

	return result
}

func countPointsAfterFirstFold(points []Point, folds []Fold) int {
	return len(foldAndReduce(points, folds[0]))
}

func displayPoints(points []Point) string {
	maxX, maxY := 0, 0

	for _, point := range points {
		if point.X > maxX {
			maxX = point.X
		}

		if point.Y > maxY {
			maxY = point.Y
		}
	}

	grid := make([][]rune, maxY+1)

	for y := range grid {
		line := make([]rune, maxX+1)
		for x := range line {
			line[x] = '.'
		}
		grid[y] = line
	}

	for _, point := range points {
		grid[point.Y][point.X] = '#'
	}

	lines := make([]string, len(grid))

	for i := range grid {
		lines[i] = string(grid[i])
	}

	return strings.Join(lines, "\n")
}

func Run(input string) {
	points, folds := parseInput(input)

	fmt.Println("Visible points after first fold:", countPointsAfterFirstFold(points, folds))
	fmt.Println("After all folds:")
	fmt.Println(displayPoints(executeFolds(points, folds)))
}
