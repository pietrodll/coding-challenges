package day3

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Parses the list of binary numbers, converting it to a list of decimal integers. Also returns the number of bits.
func parseInput(input string) ([]uint64, int) {
	binaries := strings.Split(input, "\n")

	if len(binaries) == 0 {
		panic(errors.New("empty input"))
	}

	bitNum := len(binaries[0])

	values := make([]uint64, len(binaries))

	for index, binary := range binaries {
		val, err := strconv.ParseUint(binary, 2, 64)

		if err != nil {
			panic(err)
		}

		values[index] = val
	}

	return values, bitNum
}

// Gets a bit of a number based on its index. The rightmost (least significant) bit has index 0
func getBitByIndex(value uint64, index int) uint64 {
	return (value & (1 << index)) >> index
}

// Finds the most common bit in an array of integer, depending on its index. In case of equality, considers
// that 1 is most common
func findMostCommonBitByIndex(values []uint64, index int) uint64 {
	counter := map[uint64]int{0: 0, 1: 0}

	for _, val := range values {
		counter[getBitByIndex(val, index)] += 1
	}

	if counter[0] > counter[1] {
		return 0
	} else {
		return 1
	}
}

func findEpsilonAndGamma(values []uint64, bitNum int) (int, int) {
	epsilon, gamma := uint64(0), uint64(0)

	for index := 0; index < bitNum; index++ {
		bit := findMostCommonBitByIndex(values, index)

		gamma += bit << index
		epsilon += (1 - bit) << index
	}

	return int(epsilon), int(gamma)
}

func getConsumption(values []uint64, bitNum int) int {
	epsilon, gamma := findEpsilonAndGamma(values, bitNum)
	return epsilon * gamma
}

func findRating(values []uint64, bitNum int, bitSelector func(values []uint64, index int) uint64) int {
	ratingValues := values

	for index := bitNum - 1; index >= 0; index-- {
		bit := bitSelector(ratingValues, index)

		newRatingValues := []uint64{}

		for _, val := range ratingValues {
			if getBitByIndex(val, index) == bit {
				newRatingValues = append(newRatingValues, val)
			}
		}

		ratingValues = newRatingValues

		if len(ratingValues) == 1 {
			return int(ratingValues[0])
		}
	}

	if len(ratingValues) == 1 {
		return int(ratingValues[0])
	}

	panic(errors.New("cannot find rating value"))
}

func findOxygen(values []uint64, bitNum int) int {
	return findRating(values, bitNum, findMostCommonBitByIndex)
}

func findLeastCommonBitByIndex(values []uint64, index int) uint64 {
	return 1 - findMostCommonBitByIndex(values, index)
}

func findCo2(values []uint64, bitNum int) int {
	return findRating(values, bitNum, findLeastCommonBitByIndex)
}

func getLifeSupportRating(values []uint64, bitNum int) int {
	return findCo2(values, bitNum) * findOxygen(values, bitNum)
}

func Run(input string) {
	values, bitNum := parseInput(input)

	fmt.Println("Consumption:", getConsumption(values, bitNum))
	fmt.Println("Life support rating:", getLifeSupportRating(values, bitNum))
}
