package day24

type InputInterface interface {
	GetInput() int
}

type AluState struct {
	input     InputInterface
	variables []int
}

var variableIndexes = map[rune]int{
	'w': 0,
	'x': 1,
	'y': 2,
	'z': 3,
}

func (s *AluState) GetPtr(variable rune) *int {
	return &s.variables[variableIndexes[variable]]
}

func (s *AluState) GetVar(variable rune) int {
	return s.variables[variableIndexes[variable]]
}

func (s *AluState) GetVarSafe(variable rune) (int, bool) {
	index, present := variableIndexes[variable]

	if present {
		return s.variables[index], true
	}

	return 0, false
}

type Instruction interface {
	ApplyTo(state *AluState)
}

func (s *AluState) Execute(instructions []Instruction) {
	for _, instruction := range instructions {
		instruction.ApplyTo(s)
	}
}

type BaseInstruction struct {
	left rune
}

type InputInstruction struct {
	BaseInstruction
}

func (i InputInstruction) ApplyTo(state *AluState) {
	*state.GetPtr(i.left) = state.input.GetInput()
}

type OperatorInstruction struct {
	BaseInstruction
	rightVar rune
	rightNum int
}

func (i OperatorInstruction) GetRight(s *AluState) int {
	if value, present := s.GetVarSafe(i.rightVar); present {
		return value
	}

	return i.rightNum
}

type AddInstruction struct {
	OperatorInstruction
}

func (i AddInstruction) ApplyTo(state *AluState) {
	if _, present := variableIndexes[i.rightVar]; present {
		*state.GetPtr(i.left) += *state.GetPtr(i.rightVar)
	} else {
		*state.GetPtr(i.left) += i.rightNum
	}
}

type MultiplyInstruction struct {
	OperatorInstruction
}

func (i MultiplyInstruction) ApplyTo(state *AluState) {
	if _, present := variableIndexes[i.rightVar]; present {
		*state.GetPtr(i.left) *= *state.GetPtr(i.rightVar)
	} else {
		*state.GetPtr(i.left) *= i.rightNum
	}
}

type DivideInstruction struct {
	OperatorInstruction
}

func (i DivideInstruction) ApplyTo(state *AluState) {
	if _, present := variableIndexes[i.rightVar]; present {
		*state.GetPtr(i.left) /= *state.GetPtr(i.rightVar)
	} else {
		*state.GetPtr(i.left) /= i.rightNum
	}
}

type ModuloInstruction struct {
	OperatorInstruction
}

func (i ModuloInstruction) ApplyTo(state *AluState) {
	if _, present := variableIndexes[i.rightVar]; present {
		*state.GetPtr(i.left) %= state.GetVar(i.rightVar)
	} else {
		*state.GetPtr(i.left) %= i.rightNum
	}
}

type EqualInstruction struct {
	OperatorInstruction
}

func (i EqualInstruction) ApplyTo(state *AluState) {
	value := i.rightNum

	if _, present := variableIndexes[i.rightVar]; present {
		value = state.GetVar(i.rightVar)
	}

	if ptr := state.GetPtr(i.left); *ptr == value {
		*ptr = 1
	} else {
		*ptr = 0
	}

}

type SetInstruction struct {
	OperatorInstruction
}

func (i SetInstruction) ApplyTo(state *AluState) {
	if _, present := variableIndexes[i.rightVar]; present {
		*state.GetPtr(i.left) = state.GetVar(i.rightVar)
	} else {
		*state.GetPtr(i.left) = i.rightNum
	}
}

func compileInstructions(instructions []Instruction) []Instruction {
	compiled := make([]Instruction, 0)

	i := 0
	for i < len(instructions)-1 {
		curr, okCurr := instructions[i].(MultiplyInstruction)
		next, okNext := instructions[i+1].(AddInstruction)

		if okCurr && okNext && curr.rightVar == rune(0) && curr.rightNum == 0 && next.left == curr.left {
			// group multiplication by 0 and addition -> set
			setInstr := SetInstruction{OperatorInstruction{BaseInstruction{curr.left}, next.rightVar, next.rightNum}}
			compiled = append(compiled, setInstr)
			i += 2
		} else if okCurr && curr.rightNum == 1 {
			// remove multiplication by 1
			i++
		} else if divInstr, ok := instructions[i].(DivideInstruction); ok && divInstr.rightNum == 1 {
			// remove division by 1
			i++
		} else if addInstr, ok := instructions[i].(AddInstruction); ok && addInstr.rightVar == rune(0) && addInstr.rightNum == 0 {
			// remove addition with 0
			i++
		} else {
			compiled = append(compiled, instructions[i])
			i++
		}
	}

	return compiled
}
