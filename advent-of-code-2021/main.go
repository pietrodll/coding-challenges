package main

import (
	"github.com/pietrodll/aoc2021/day1"
	"github.com/pietrodll/aoc2021/day2"
	"github.com/pietrodll/aoc2021/day3"
	"github.com/pietrodll/aoc2021/day4"
	"github.com/pietrodll/aoc2021/day5"
	"github.com/pietrodll/aoc2021/day6"
	"github.com/pietrodll/aoc2021/utils"
)

func main() {
	utils.RunDay(day1.Run, 1)
	utils.RunDay(day2.Run, 2)
	utils.RunDay(day3.Run, 3)
	utils.RunDay(day4.Run, 4)
	utils.RunDay(day5.Run, 5)
	utils.RunDay(day6.Run, 6)
}
