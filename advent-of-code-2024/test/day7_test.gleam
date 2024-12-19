import day7
import gleam/list
import gleeunit/should

const input = "190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20"

pub fn can_be_resolved_test() {
  input
  |> day7.parse_input
  |> should.be_ok
  |> list.map(fn(equation) { day7.can_be_resolved(equation.0, equation.1) })
  |> should.equal([True, True, False, False, False, False, False, False, True])
}

pub fn can_be_resolved_with_concatenation_test() {
  input
  |> day7.parse_input
  |> should.be_ok
  |> list.map(fn(equation) {
    day7.can_be_resolved_with_concatenation(equation.0, equation.1)
  })
  |> should.equal([True, True, False, True, True, False, True, False, True])
}
