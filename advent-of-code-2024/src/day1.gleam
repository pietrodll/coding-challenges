import gleam/dict
import gleam/int
import gleam/list
import gleam/result
import gleam/string
import utils

fn split_line(line: String) -> Result(#(String, String), String) {
  case line |> string.trim |> string.split(on: "   ") {
    [left, right] -> Ok(#(left, right))
    _ -> Error("could not split line " <> line)
  }
}

fn parse_line(line: String) -> Result(#(Int, Int), String) {
  use #(left, right) <- result.try(split_line(line))
  use left_int <- result.try(utils.parse_int(left))
  use right_int <- result.map(utils.parse_int(right))
  #(left_int, right_int)
}

pub fn parse_input(content: String) -> Result(#(List(Int), List(Int)), String) {
  content
  |> string.trim
  |> string.split(on: "\n")
  |> list.try_map(parse_line)
  |> result.map(list.unzip)
}

pub fn compute_distance(left: List(Int), right: List(Int)) -> Int {
  let left_sorted = list.sort(left, by: int.compare)
  let right_sorted = list.sort(right, by: int.compare)

  list.map2(left_sorted, right_sorted, fn(left, right) {
    int.absolute_value(left - right)
  })
  |> int.sum
}

pub fn compute_similarity(left: List(Int), right: List(Int)) -> Int {
  let counts = utils.counter(right)

  list.fold(left, 0, fn(total, number) {
    total + { number * { dict.get(counts, number) |> result.unwrap(0) } }
  })
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use #(left, right) <- result.try(parse_input(input))
  let distance = compute_distance(left, right)
  let similarity = compute_similarity(left, right)

  Ok(utils.AdventOfCodeResult(
    int.to_string(distance),
    int.to_string(similarity),
  ))
}
