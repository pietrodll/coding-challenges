package main

import (
	"flag"
	"fmt"

	"github.com/pietrodll/aoc2021/day1"
	"github.com/pietrodll/aoc2021/day10"
	"github.com/pietrodll/aoc2021/day11"
	"github.com/pietrodll/aoc2021/day12"
	"github.com/pietrodll/aoc2021/day13"
	"github.com/pietrodll/aoc2021/day14"
	"github.com/pietrodll/aoc2021/day15"
	"github.com/pietrodll/aoc2021/day16"
	"github.com/pietrodll/aoc2021/day17"
	"github.com/pietrodll/aoc2021/day18"
	"github.com/pietrodll/aoc2021/day19"
	"github.com/pietrodll/aoc2021/day2"
	"github.com/pietrodll/aoc2021/day20"
	"github.com/pietrodll/aoc2021/day21"
	"github.com/pietrodll/aoc2021/day22"
	"github.com/pietrodll/aoc2021/day23"
	"github.com/pietrodll/aoc2021/day24"
	"github.com/pietrodll/aoc2021/day25"
	"github.com/pietrodll/aoc2021/day3"
	"github.com/pietrodll/aoc2021/day4"
	"github.com/pietrodll/aoc2021/day5"
	"github.com/pietrodll/aoc2021/day6"
	"github.com/pietrodll/aoc2021/day7"
	"github.com/pietrodll/aoc2021/day8"
	"github.com/pietrodll/aoc2021/day9"
	"github.com/pietrodll/aoc2021/utils"
)

var runnables = []func(string){
	day1.Run,
	day2.Run,
	day3.Run,
	day4.Run,
	day5.Run,
	day6.Run,
	day7.Run,
	day8.Run,
	day9.Run,
	day10.Run,
	day11.Run,
	day12.Run,
	day13.Run,
	day14.Run,
	day15.Run,
	day16.Run,
	day17.Run,
	day18.Run,
	day19.Run,
	day20.Run,
	day21.Run,
	day22.Run,
	day23.Run,
	day24.Run,
	day25.Run,
}

func main() {
	dayPtr := flag.Int("day", 0, "the day to run. If not provided or equal to zero, all days are run")
	flag.Parse()

	if *dayPtr == 0 {
		for dayIndex, runFunc := range runnables {
			utils.RunDay(runFunc, dayIndex+1)
		}
	} else if *dayPtr <= len(runnables) {
		utils.RunDay(runnables[*dayPtr-1], *dayPtr)
	} else {
		panic(fmt.Errorf("invalid day: %d", *dayPtr))
	}
}
