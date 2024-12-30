import day1
import day2
import day3
import day4
import day5
import day6
import day7
import day8
import day9
import gleam/list
import utils

const days_functions: List(
  fn(String) -> Result(utils.AdventOfCodeResult, String),
) = [
  day1.run, day2.run, day3.run, day4.run, day5.run, day6.run, day7.run, day8.run,
  day9.run,
]

pub fn main() {
  days_functions
  |> list.index_map(fn(day_func, day_idx) {
    utils.run_day(day_idx + 1, day_func)
  })
}
