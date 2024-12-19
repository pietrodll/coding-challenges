import day6
import gleam/set
import gleeunit/should

const input = "....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..."

pub fn parse_input_test() {
  let #(grid_size, guard, obstacles) = day6.parse_input(input) |> should.be_ok

  grid_size |> should.equal(#(10, 10))

  guard |> should.equal(#(6, 4, day6.Up))

  obstacles
  |> should.equal(
    set.from_list([
      #(0, 4),
      #(1, 9),
      #(3, 2),
      #(4, 7),
      #(6, 1),
      #(7, 8),
      #(8, 0),
      #(9, 6),
    ]),
  )
}

pub fn count_guard_positions_test() {
  let #(grid_size, guard, obstacles) = day6.parse_input(input) |> should.be_ok

  day6.count_guard_positions(grid_size, guard, obstacles) |> should.equal(41)
}
