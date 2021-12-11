package day6

import (
	"fmt"

	"github.com/pietrodll/aoc2021/utils/parse"
)

func parseInput(input string) []int {
	return parse.ParseIntegers(input, ",")
}

type StateTracker struct {
	countByTimer []int
	maxTimer     int
	resetTo      int
}

func initStateTracker(initialState []int, maxTimer int, resetTo int) StateTracker {
	countByTimer := make([]int, maxTimer+1)

	for _, val := range initialState {
		if val > maxTimer {
			panic(fmt.Errorf("lanternfish timer cannot be %d", val))
		}

		countByTimer[val] += 1
	}

	return StateTracker{countByTimer, maxTimer, resetTo}
}

func (s *StateTracker) nextState() {
	goingToReset := s.countByTimer[0]

	for i := 0; i < s.maxTimer; i++ {
		s.countByTimer[i] = s.countByTimer[i+1]
	}

	s.countByTimer[s.resetTo] += goingToReset
	s.countByTimer[s.maxTimer] = goingToReset
}

func (s *StateTracker) simulateGrowth(days int) {
	for i := 0; i < days; i++ {
		s.nextState()
	}
}

func (s *StateTracker) countFishes() int {
	tot := 0

	for _, val := range s.countByTimer {
		tot += val
	}

	return tot
}

func Run(input string) {
	intialState := parseInput(input)

	tracker := initStateTracker(intialState, 8, 6)
	tracker.simulateGrowth(80)

	fmt.Println("Number of fishes after 80 days:", tracker.countFishes())

	tracker.simulateGrowth(256 - 80)

	fmt.Println("Number of fishes after 256 days:", tracker.countFishes())
}
