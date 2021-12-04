package utils

import "fmt"

func RunDay(runFunc func(input string), day int) {
	fmt.Printf("------------ DAY %d ------------\n", day)
	input := LoadInput(day)
	runFunc(input)
	fmt.Println()
}
