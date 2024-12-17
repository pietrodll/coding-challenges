import day4
import gleeunit/should

const input = "MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX"

pub fn parse_input_test() {
  input |> day4.parse_input |> should.be_ok
}

pub fn count_xmas_test() {
  input
  |> day4.parse_input
  |> should.be_ok
  |> day4.count_xmas
  |> should.equal(18)
}

pub fn count_x_mas_test() {
  input
  |> day4.parse_input
  |> should.be_ok
  |> day4.count_x_mas
  |> should.equal(9)
}
