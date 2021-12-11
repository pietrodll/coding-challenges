package parse

import (
	"strconv"
	"strings"
)

func ParseIntegers(data string, sep string) []int {
	strValues := strings.Split(data, sep)
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
