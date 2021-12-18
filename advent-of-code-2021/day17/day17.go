package day17

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/pietrodll/aoc2021/utils/parse"
)

type ProbeState struct {
	X, Y, Vx, Vy int
}

func newProbeState(vx, vy int) ProbeState {
	return ProbeState{0, 0, vx, vy}
}

func nextStep(p ProbeState) ProbeState {
	var nextVx int

	if p.Vx == 0 {
		nextVx = 0
	} else if p.Vx > 0 {
		nextVx = p.Vx - 1
	} else {
		nextVx = p.Vx + 1
	}

	return ProbeState{p.X + p.Vx, p.Y + p.Vy, nextVx, p.Vy - 1}
}

type TargetArea struct {
	MinX, MaxX, MinY, MaxY int
}

func parseInput(input string) TargetArea {
	pattern := regexp.MustCompile(`-?\d+`)

	match := pattern.FindAllString(input, 4)
	values := parse.StringsToIntegers(match)

	return TargetArea{values[0], values[1], values[2], values[3]}
}

func (t TargetArea) Contains(p ProbeState) bool {
	return p.X <= t.MaxX && p.X >= t.MinX && p.Y <= t.MaxY && p.Y >= t.MinY
}

func (t TargetArea) WillNeverContain(p ProbeState) bool {
	// we consider a positive initial Vx

	if p.X > t.MaxX {
		// if the probe is past the target in the X direction, it will never come back to enter it
		return true
	}

	if p.Y < t.MinY && p.Vy <= 0 {
		// if the probe is under the target, it will never come up to enter it
		return true
	}

	return false
}

func isValidInitialVelocity(t TargetArea, vx, vy int) bool {
	p := newProbeState(vx, vy)

	for !t.WillNeverContain(p) {
		if t.Contains(p) {
			return true
		}

		p = nextStep(p)
	}

	return false
}

func findVxFromFinalPosition(finalPos int) (int, bool) {
	// solve the equation finalPos = vx * (vx + 1) / 2

	if finalPos == 0 {
		return 0, true
	}

	max, min := finalPos, 1

	for min < max {
		mid := (min + max) / 2

		if candidate := mid * (mid + 1) / 2; candidate == finalPos {
			return mid, true
		} else if candidate > finalPos {
			max = mid
		} else {
			min = mid + 1
		}
	}

	return 0, false
}

func findFirstValidVx(t TargetArea) int {
	for xFinalPos := t.MinX; xFinalPos <= t.MaxX; xFinalPos++ {
		if vx, ok := findVxFromFinalPosition(xFinalPos); ok {
			return vx
		}
	}

	panic(errors.New("cannot find suitable vx"))
}

func findHighestPossibleAltitute(t TargetArea) int {
	maxAltitude := 0
	vy := 0
	vx := findFirstValidVx(t)

	invalidCount := 0
	for invalidCount < 1000 {
		// we stop when we find 1000 invalid velocities in a row
		if isValidInitialVelocity(t, vx, vy) {
			invalidCount = 0

			if alt := vy * (vy + 1) / 2; alt > maxAltitude {
				maxAltitude = alt
			}
		} else {
			invalidCount++
		}

		vy++
	}

	return maxAltitude
}

func countValidInitialVelocities(t TargetArea) int {
	count := 0

	vx := findFirstValidVx(t)
	minVy := t.MinY - 1

	invalidCountX := 0

	for invalidCountX < 500 {
		invalidCountY := 0
		isVxValid := false
		vy := minVy

		for invalidCountY < 500 {
			if isValidInitialVelocity(t, vx, vy) {
				count++
				invalidCountY = 0
				isVxValid = true
			} else {
				invalidCountY++
			}

			vy++
		}

		if isVxValid {
			invalidCountX = 0
		} else {
			invalidCountX++
		}

		vx++
	}

	return count
}

func Run(input string) {
	target := parseInput(input)

	fmt.Println("Highest possible altitude:", findHighestPossibleAltitute(target))
	fmt.Println("Total valid initial velocities:", countValidInitialVelocities(target))
}
