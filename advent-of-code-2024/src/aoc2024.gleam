import day1
import gleam/list
import utils

const days_functions: List(
  fn(String) -> Result(utils.AdventOfCodeResult, String),
) = [day1.run]

pub fn main() {
  days_functions
  |> list.index_map(fn(day_func, day_idx) {
    utils.run_day(day_idx + 1, day_func)
  })
}
