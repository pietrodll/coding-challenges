import day2
import gleam/list
import gleam/option
import gleeunit/should

const input = "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"

const expected_input: List(List(Int)) = [
  [7, 6, 4, 2, 1], [1, 2, 7, 8, 9], [9, 7, 6, 2, 1], [1, 3, 2, 4, 5],
  [8, 6, 4, 4, 1], [1, 3, 6, 7, 9],
]

pub fn parse_input_test() {
  day2.parse_input(input)
  |> should.be_ok
  |> should.equal(expected_input)
}

pub fn is_safe_test() {
  expected_input
  |> list.map(day2.is_safe)
  |> should.equal([True, False, False, False, False, True])
}

pub fn compute_status_test() {
  expected_input
  |> list.map(day2.compute_status)
  |> option.all
  |> should.be_some
  |> should.equal([
    day2.Safe(False),
    day2.Unsafe(1),
    day2.Unsafe(2),
    day2.Unsafe(1),
    day2.Unsafe(2),
    day2.Safe(True),
  ])
}

pub fn is_safe_with_problem_handler_test() {
  expected_input
  |> list.map(day2.is_safe_with_problem_handler)
  |> should.equal([True, False, False, True, True, True])
}
