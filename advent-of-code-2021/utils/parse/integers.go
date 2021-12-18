package parse

import (
	"strconv"
	"strings"
)

func StringsToIntegers(strValues []string) []int {
	values := make([]int, len(strValues))

	for i, strVal := range strValues {
		if val, err := strconv.Atoi(strVal); err == nil {
			values[i] = val
		} else {
			panic(err)
		}
	}

	return values
}

func ParseIntegers(data string, sep string) []int {
	return StringsToIntegers(strings.Split(data, sep))
}
