import gleam/int
import gleam/list
import gleam/result
import gleam/string
import utils

fn parse_line(line: String) -> Result(#(Int, List(Int)), String) {
  case line |> string.split(": ") {
    [res_str, values_str] -> {
      use res <- result.try(utils.parse_int(res_str))
      use values <- result.map(
        values_str |> string.split(" ") |> list.try_map(utils.parse_int),
      )

      #(res, values)
    }
    _ -> Error("could not parse line \"" <> line <> "\"")
  }
}

pub fn parse_input(input: String) -> Result(List(#(Int, List(Int))), String) {
  input |> string.trim |> string.split("\n") |> list.try_map(parse_line)
}

fn can_be_resolved_rec(test_value: Int, acc: Int, numbers: List(Int)) -> Bool {
  case numbers {
    [] -> test_value == acc
    [first, ..rest] ->
      can_be_resolved_rec(test_value, acc + first, rest)
      || can_be_resolved_rec(test_value, acc * first, rest)
  }
}

pub fn can_be_resolved(test_value: Int, numbers: List(Int)) -> Bool {
  case numbers {
    [] -> False
    [first, ..rest] -> can_be_resolved_rec(test_value, first, rest)
  }
}

fn concat_numbers(first: Int, second: Int) -> Int {
  let assert Ok(first_digits) = int.digits(first, 10)
  let assert Ok(second_digits) = int.digits(second, 10)
  let assert Ok(result) =
    int.undigits(list.append(first_digits, second_digits), 10)
  result
}

fn can_be_resolved_with_concatenation_rec(
  test_value: Int,
  acc: Int,
  numbers: List(Int),
) -> Bool {
  case numbers {
    [] -> test_value == acc
    [first, ..rest] ->
      can_be_resolved_with_concatenation_rec(test_value, acc + first, rest)
      || can_be_resolved_with_concatenation_rec(test_value, acc * first, rest)
      || can_be_resolved_with_concatenation_rec(
        test_value,
        concat_numbers(acc, first),
        rest,
      )
  }
}

pub fn can_be_resolved_with_concatenation(
  test_value: Int,
  numbers: List(Int),
) -> Bool {
  case numbers {
    [] -> False
    [first, ..rest] ->
      can_be_resolved_with_concatenation_rec(test_value, first, rest)
  }
}

pub fn sum_up_resolvable(
  equations: List(#(Int, List(Int))),
  resolvable_func: fn(Int, List(Int)) -> Bool,
) {
  equations
  |> list.filter_map(fn(equation) {
    let #(test_value, numbers) = equation

    case resolvable_func(test_value, numbers) {
      True -> Ok(test_value)
      False -> Error(Nil)
    }
  })
  |> list.fold(0, int.add)
}

pub fn first_part(equations: List(#(Int, List(Int)))) -> Int {
  sum_up_resolvable(equations, can_be_resolved)
}

pub fn second_part(equations: List(#(Int, List(Int)))) -> Int {
  sum_up_resolvable(equations, can_be_resolved_with_concatenation)
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use equations <- result.map(parse_input(input))

  let first_part = first_part(equations)
  let second_part = second_part(equations)

  utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  )
}
