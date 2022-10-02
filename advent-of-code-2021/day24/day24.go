package day24

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func getDigits(number int) []int {
	n := number
	digits := make([]int, 0)

	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	reverse(digits)
	return digits
}

func assembleDigits(digits []int) int {
	result := 0
	powOf10 := 1

	for i := len(digits) - 1; i >= 0; i-- {
		result += digits[i] * powOf10
		powOf10 *= 10
	}

	return result
}

type ModelNumberInputInterface struct {
	modelNumberDigits []int
	index             int
}

func (i *ModelNumberInputInterface) GetInput() int {
	result := i.modelNumberDigits[i.index]
	i.index++
	return result
}

func InitAluStateFromDigits(modelNumberDigits []int) AluState {
	variables := make([]int, 4)
	return AluState{&ModelNumberInputInterface{modelNumberDigits, 0}, variables}
}

func InitAluState(modelNumber int) AluState {
	return InitAluStateFromDigits(getDigits(modelNumber))
}

var instructionPattern = regexp.MustCompile(`^(inp|add|mul|div|mod|eql) (w|x|y|z)( (w|x|y|z|-?\d+))?$`)

func parseInstruction(instructionStr string) Instruction {
	match := instructionPattern.FindAllStringSubmatch(instructionStr, 1)

	if len(match) != 1 {
		panic(fmt.Errorf("cannot parse %s", instructionStr))
	}

	submatches := match[0]
	base := BaseInstruction{[]rune(submatches[2])[0]}

	if submatches[1] == "inp" {
		return InputInstruction{base}
	}

	if len(submatches) != 5 {
		panic(fmt.Errorf("cannot parse %s", instructionStr))
	}

	var (
		rightVar rune
		rightNum int
	)

	if num, err := strconv.Atoi(submatches[4]); err != nil {
		rightVar = []rune(submatches[4])[0]
	} else {
		rightNum = num
	}

	operation := OperatorInstruction{base, rightVar, rightNum}

	switch submatches[1] {
	case "add":
		return AddInstruction{operation}
	case "mul":
		return MultiplyInstruction{operation}
	case "div":
		return DivideInstruction{operation}
	case "mod":
		return ModuloInstruction{operation}
	default:
		return EqualInstruction{operation}
	}
}

func parseInput(input string) []Instruction {
	lines := strings.Split(input, "\n")
	instructions := make([]Instruction, len(lines))

	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}

	return instructions
}

func isValidModelNumber(digits []int, instructions []Instruction) bool {
	state := InitAluStateFromDigits(digits)
	state.Execute(instructions)

	return state.GetVar('z') == 0
}

func decrementDigits(digits []int) {
	decrementIndex := len(digits) - 1
	digits[decrementIndex]--

	for decrementIndex > 0 && digits[decrementIndex] == 0 {
		digits[decrementIndex] = 9
		decrementIndex--
		digits[decrementIndex]--
	}
}

func findHighestValidModelNumber(instructions []Instruction) int {
	digits := make([]int, 14)

	for i := range digits {
		digits[i] = 9
	}

	compiled := compileInstructions(instructions)
	fmt.Println("Before compiling:", len(instructions), "After:", len(compiled))

	for digits[0] >= 1 {
		if isValidModelNumber(digits, compiled) {
			return assembleDigits(digits)
		}

		decrementDigits(digits)
	}

	panic(errors.New("could not find a valid model number"))
}

func Run(input string) {
	instructions := parseInput(input)
	fmt.Println("Highest valid model number:", findHighestValidModelNumber(instructions))
}
