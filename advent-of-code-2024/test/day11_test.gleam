import day11
import gleam/dict
import gleam/list
import gleeunit/should

const input = "125 17"

pub fn parse_input_test() {
  input |> day11.parse_input |> should.be_ok |> should.equal([125, 17])
}

pub fn blink_stone_test() {
  day11.blink_stone(125) |> should.equal([253_000])
  day11.blink_stone(17) |> should.equal([1, 7])
  day11.blink_stone(253_000) |> should.equal([253, 0])
  day11.blink_stone(0) |> should.equal([1])
}

pub fn blink_count_test() {
  [
    [#(125, 1), #(17, 1)],
    [#(253_000, 1), #(1, 1), #(7, 1)],
    [#(253, 1), #(0, 1), #(2024, 1), #(14_168, 1)],
    [#(512_072, 1), #(1, 1), #(20, 1), #(24, 1), #(28_676_032, 1)],
    [
      #(512, 1),
      #(72, 1),
      #(2024, 1),
      #(2, 2),
      #(0, 1),
      #(4, 1),
      #(2867, 1),
      #(6032, 1),
    ],
    [
      #(1_036_288, 1),
      #(7, 1),
      #(2, 1),
      #(20, 1),
      #(24, 1),
      #(4048, 2),
      #(1, 1),
      #(8096, 1),
      #(28, 1),
      #(67, 1),
      #(60, 1),
      #(32, 1),
    ],
  ]
  |> list.map(dict.from_list)
  |> list.reduce(fn(previous, next) {
    day11.blink_count(previous) |> should.equal(next)
    next
  })
  |> should.be_ok
}

pub fn count_stones_after_blink_test() {
  [125, 17] |> day11.count_stones_after_blink(6) |> should.equal(22)
  [125, 17] |> day11.count_stones_after_blink(25) |> should.equal(55_312)
}
