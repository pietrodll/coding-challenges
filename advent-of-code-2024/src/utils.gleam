import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/option
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

pub fn index_fold_until(
  over list: List(a),
  from initial: b,
  with fun: fn(b, Int, a) -> list.ContinueOrStop(b),
) -> b {
  list
  |> list.fold_until(#(initial, 0), fn(acc_and_index, element) {
    let #(acc, index) = acc_and_index
    case fun(acc, index, element) {
      list.Stop(x) -> list.Stop(#(x, index + 1))
      list.Continue(x) -> list.Continue(#(x, index + 1))
    }
  })
  |> fn(tup) { tup.0 }
}

fn append_reverse(target: List(a), source: List(a)) -> List(a) {
  case source {
    [] -> target
    [first, ..rest] -> append_reverse([first, ..target], rest)
  }
}

fn remove_index_rec(acc: List(a), list: List(a), index: Int) -> List(a) {
  case list {
    [] -> acc
    [first, ..rest] -> {
      case index {
        0 -> append_reverse(acc, rest)
        _ -> remove_index_rec([first, ..acc], rest, index - 1)
      }
    }
  }
}

pub fn remove_index(list: List(a), index: Int) -> List(a) {
  remove_index_rec([], list, index) |> list.reverse
}

pub fn parse_int(string: String) -> Result(Int, String) {
  int.parse(string)
  |> result.replace_error("error parsing int \"" <> string <> "\"")
}

pub fn call_multiple_times(arg: a, func: fn(a) -> a, times: Int) -> a {
  case times <= 0 {
    True -> arg
    False -> call_multiple_times(func(arg), func, times - 1)
  }
}

pub fn counter(list: List(a)) -> dict.Dict(a, Int) {
  list.fold(list, dict.new(), fn(count, item) {
    count
    |> dict.upsert(item, fn(previous_count) {
      option.unwrap(previous_count, 0) + 1
    })
  })
}
