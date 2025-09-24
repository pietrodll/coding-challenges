import day13
import gleam/list
import gleam/option
import gleeunit/should

const input = "Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279"

fn must_parse_input(input: String) -> List(day13.Game) {
  input |> day13.parse_input |> should.be_ok
}

pub fn parse_input_test() {
  input
  |> must_parse_input
  |> should.equal([
    day13.Game(
      day13.Coords(94, 34),
      day13.Coords(22, 67),
      day13.Coords(8400, 5400),
    ),
    day13.Game(
      day13.Coords(26, 66),
      day13.Coords(67, 21),
      day13.Coords(12_748, 12_176),
    ),
    day13.Game(
      day13.Coords(17, 86),
      day13.Coords(84, 37),
      day13.Coords(7870, 6450),
    ),
    day13.Game(
      day13.Coords(69, 23),
      day13.Coords(27, 71),
      day13.Coords(18_641, 10_279),
    ),
  ])
}

pub fn game_price_test() {
  input
  |> must_parse_input
  |> list.map(day13.game_cost)
  |> should.equal([option.Some(280), option.None, option.Some(200), option.None])
}
