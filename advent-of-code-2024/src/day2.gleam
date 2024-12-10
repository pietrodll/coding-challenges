import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/string
import utils

fn parse_line(line: String) -> Result(List(Int), String) {
  line
  |> string.trim
  |> string.split(on: " ")
  |> list.try_map(utils.parse_int)
}

pub fn parse_input(input: String) -> Result(List(List(Int)), String) {
  input
  |> string.trim
  |> string.split(on: "\n")
  |> list.try_map(parse_line)
}

fn compute_diffs(report: List(Int)) -> List(Int) {
  case report {
    [] -> []
    [first, ..rest] ->
      rest
      |> list.map_fold(from: first, with: fn(previous, item) {
        #(item, item - previous)
      })
      |> fn(tup) { tup.1 }
  }
}

pub type DiffStatus {
  Unsafe(index: Int)
  Safe(increasing: Bool)
}

fn diff_status(index: Int, diff: Int) -> DiffStatus {
  let abs_diff = int.absolute_value(diff)
  let is_safe = abs_diff >= 1 && abs_diff <= 3

  case is_safe {
    True -> Safe(increasing: diff > 0)
    False -> Unsafe(index)
  }
}

fn is_current_diff_safe(
  prev_status: DiffStatus,
  curr_status: DiffStatus,
) -> Bool {
  case prev_status, curr_status {
    Unsafe(_), _ | _, Unsafe(_) -> False
    Safe(prev_increasing), Safe(curr_increasing) ->
      curr_increasing == prev_increasing
  }
}

pub fn compute_status(report: List(Int)) -> option.Option(DiffStatus) {
  report
  |> compute_diffs
  |> utils.index_fold_until(
    option.None,
    fn(maybe_prev_status, index, curr_diff) {
      case maybe_prev_status {
        option.None ->
          case diff_status(index, curr_diff) {
            Unsafe(idx) -> list.Stop(option.Some(Unsafe(idx)))
            other -> list.Continue(option.Some(other))
          }
        option.Some(prev_status) ->
          case
            is_current_diff_safe(prev_status, diff_status(index, curr_diff))
          {
            True -> list.Continue(option.Some(prev_status))
            False -> list.Stop(option.Some(Unsafe(index)))
          }
      }
    },
  )
}

pub fn is_safe(report: List(Int)) -> Bool {
  report
  |> compute_status
  |> option.map(fn(status) {
    case status {
      Unsafe(_) -> False
      Safe(_) -> True
    }
  })
  |> option.unwrap(False)
}

pub fn is_safe_with_problem_handler(report: List(Int)) -> Bool {
  report
  |> compute_status
  |> option.map(fn(status) {
    case status {
      Unsafe(problem_index) ->
        is_safe(utils.remove_index(report, problem_index))
        || is_safe(utils.remove_index(report, problem_index + 1))
      Safe(_) -> True
    }
  })
  |> option.unwrap(False)
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use parsed_input <- result.try(parse_input(input))

  let safe_reports = parsed_input |> list.count(is_safe)
  let safe_reports_with_problem_handler =
    parsed_input |> list.count(is_safe_with_problem_handler)

  Ok(utils.AdventOfCodeResult(
    int.to_string(safe_reports),
    int.to_string(safe_reports_with_problem_handler),
  ))
}
