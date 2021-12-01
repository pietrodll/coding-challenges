package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pietrodll/aoc2021/utils"
)

func parseValues(input string) []int {
	values := strings.Split(input, "\n")

	result := make([]int, len(values))

	for i, val := range values {
		intVal, err := strconv.Atoi(val)

		if err != nil {
			panic(err)
		}

		result[i] = intVal
	}

	return result
}

func sumBy(values []int, groupSize int) []int {
	if len(values) < groupSize {
		return make([]int, 0)
	}

	result := make([]int, len(values)-groupSize+1)

	for i := 0; i <= len(values)-groupSize; i++ {
		groupVal := 0

		for _, val := range values[i : i+groupSize] {
			groupVal += val
		}

		result[i] = groupVal
	}

	return result
}

func countIncrease(values []int) int {
	count := 0

	if len(values) == 0 {
		return 0
	}

	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			count++
		}
	}

	return count
}

func main() {
	values := parseValues(utils.LoadInput(1))
	increases := countIncrease(values)
	groupIncreases := countIncrease(sumBy(values, 3))

	fmt.Println("Count:", increases)
	fmt.Println("Count by 3:", groupIncreases)
}
