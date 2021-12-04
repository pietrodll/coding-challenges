package day2

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MoveInstruction struct {
	Direction string
	Value     int
}

func parseInstructions(input string) []MoveInstruction {
	lines := strings.Split(input, "\n")

	values := make([]MoveInstruction, len(lines))

	for index, line := range lines {
		split := strings.Split(line, " ")
		val, err := strconv.Atoi(split[1])

		if err != nil {
			panic(err)
		}

		values[index] = MoveInstruction{split[0], val}
	}

	return values
}

type Position struct {
	Horiz int
	Depth int
}

func (pos *Position) update(instuction MoveInstruction) {
	switch instuction.Direction {
	case "forward":
		pos.Horiz += instuction.Value
	case "up":
		pos.Depth -= instuction.Value
	case "down":
		pos.Depth += instuction.Value
	default:
		panic(errors.New("invalid direction"))
	}
}

type PositionWithAim struct {
	Position
	Aim int
}

func (pos *PositionWithAim) update(instruction MoveInstruction) {
	switch instruction.Direction {
	case "forward":
		pos.Horiz += instruction.Value
		pos.Depth += pos.Aim * instruction.Value
	case "up":
		pos.Aim -= instruction.Value
	case "down":
		pos.Aim += instruction.Value
	default:
		panic(errors.New("invalid direction"))
	}
}

type UpdatableByInstruction interface {
	update(MoveInstruction)
}

func executeInstructions(instructions []MoveInstruction, pos UpdatableByInstruction) {
	for _, instruction := range instructions {
		pos.update(instruction)
	}
}

func Run(input string) {
	instructions := parseInstructions(input)
	position := Position{0, 0}
	executeInstructions(instructions, &position)

	positionWithAim := PositionWithAim{Position{0, 0}, 0}
	executeInstructions(instructions, &positionWithAim)

	fmt.Println("Final position:", position.Horiz*position.Depth)
	fmt.Println("Final position with aim:", positionWithAim.Horiz*positionWithAim.Depth)
}
