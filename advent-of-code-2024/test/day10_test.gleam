import day10
import gleam/set
import gleeunit/should

const input = "89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732"

pub fn find_trailheads_test() {
  input
  |> day10.parse_input
  |> should.be_ok
  |> day10.find_trailheads
  |> set.from_list
  |> should.equal(
    set.from_list([
      #(0, 2),
      #(0, 4),
      #(2, 4),
      #(4, 6),
      #(5, 2),
      #(5, 5),
      #(6, 0),
      #(6, 6),
      #(7, 1),
    ]),
  )
}

pub fn score_starting_from_test() {
  let grid = input |> day10.parse_input |> should.be_ok

  day10.score_starting_from(grid, #(0, 2)) |> should.equal(5)
  day10.score_starting_from(grid, #(0, 4)) |> should.equal(6)
}

pub fn total_score_test() {
  input
  |> day10.parse_input
  |> should.be_ok
  |> day10.total_score
  |> should.equal(36)
}

pub fn rating_starting_from_test() {
  let grid = input |> day10.parse_input |> should.be_ok

  day10.rating_starting_from(grid, #(0, 2)) |> should.equal(20)
  day10.rating_starting_from(grid, #(0, 4)) |> should.equal(24)
}

pub fn total_rating_test() {
  input
  |> day10.parse_input
  |> should.be_ok
  |> day10.total_rating
  |> should.equal(81)
}
