import gleam/int
import gleam/io
import gleam/result
import gleam/string
import simplifile

pub type AdventOfCodeResult {
  AdventOfCodeResult(first: String, second: String)
}

pub fn read_day_file(day: Int) -> Result(String, String) {
  simplifile.read("data/day" <> int.to_string(day) <> ".txt")
  |> result.map_error(fn(err) {
    "error while opening file for day "
    <> int.to_string(day)
    <> ": "
    <> string.inspect(err)
  })
}

pub fn run_day(day: Int, run: fn(String) -> Result(AdventOfCodeResult, String)) {
  io.println("Running day " <> int.to_string(day))
  case read_day_file(day) |> result.try(run) {
    Ok(AdventOfCodeResult(first, second)) -> {
      io.println("First part: " <> first)
      io.println("Second part: " <> second)
    }
    Error(err) ->
      io.println_error(
        "Error while running day " <> int.to_string(day) <> ": " <> err,
      )
  }
}
