import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/set
import gleam/string
import utils

fn fold_grid(
  over grid: String,
  from initial: a,
  with fun: fn(a, Int, Int, String) -> a,
) -> a {
  grid
  |> string.split("\n")
  |> list.index_fold(initial, fn(acc, line, i) {
    line
    |> string.to_graphemes
    |> list.index_fold(acc, fn(acc1, char, j) { fun(acc1, i, j, char) })
  })
}

pub fn parse_input(
  input: String,
) -> #(#(Int, Int), dict.Dict(String, List(#(Int, Int)))) {
  input
  |> string.trim
  |> fold_grid(#(#(0, 0), dict.new()), fn(acc, i, j, char) {
    let #(size, positions_by_char) = acc

    let new_positions_by_char = case char {
      "." -> positions_by_char
      _ ->
        positions_by_char
        |> dict.upsert(char, fn(prev) {
          prev |> option.lazy_unwrap(list.new) |> list.prepend(#(i, j))
        })
    }

    let new_size = #(int.max(size.0, i + 1), int.max(size.0, j + 1))

    #(new_size, new_positions_by_char)
  })
}

fn is_in_grid(grid_size: #(Int, Int), position: #(Int, Int)) -> Bool {
  let #(h, w) = grid_size
  let #(i, j) = position

  0 <= i && i < h && 0 <= j && j < w
}

fn compute_antinodes_rec(
  acc: List(#(Int, Int)),
  grid_size: #(Int, Int),
  x: Int,
  y: Int,
  dx: Int,
  dy: Int,
) -> List(#(Int, Int)) {
  let next_x = x + dx
  let next_y = y + dy

  case is_in_grid(grid_size, #(next_x, next_y)) {
    False -> acc
    True ->
      compute_antinodes_rec(
        [#(next_x, next_y), ..acc],
        grid_size,
        next_x,
        next_y,
        dx,
        dy,
      )
  }
}

pub fn compute_antinodes_with_resonance(
  grid_size: #(Int, Int),
  p1: #(Int, Int),
  p2: #(Int, Int),
) -> List(#(Int, Int)) {
  let #(x1, y1) = p1
  let #(x2, y2) = p2

  let dx = x2 - x1
  let dy = y2 - y1

  list.new()
  |> compute_antinodes_rec(grid_size, x1, y1, dx, dy)
  |> compute_antinodes_rec(grid_size, x2, y2, -dx, -dy)
}

pub fn compute_antinodes_for_two(
  grid_size: #(Int, Int),
  p1: #(Int, Int),
  p2: #(Int, Int),
) -> List(#(Int, Int)) {
  let #(x1, y1) = p1
  let #(x2, y2) = p2

  let dx = x2 - x1
  let dy = y2 - y1

  [#(x1 - dx, y1 - dy), #(x2 + dx, y2 + dy)]
  |> list.filter(fn(position) { is_in_grid(grid_size, position) })
}

fn generate_antinodes(
  positions: List(#(Int, Int)),
  grid_size: #(Int, Int),
  compute_antinodes_fn: fn(#(Int, Int), #(Int, Int), #(Int, Int)) ->
    List(#(Int, Int)),
) -> List(#(Int, Int)) {
  positions
  |> list.combination_pairs
  |> list.flat_map(fn(pair) { compute_antinodes_fn(grid_size, pair.0, pair.1) })
}

pub fn count_antinodes(
  grid_size: #(Int, Int),
  positions_by_char: dict.Dict(String, List(#(Int, Int))),
  compute_antinodes_fn: fn(#(Int, Int), #(Int, Int), #(Int, Int)) ->
    List(#(Int, Int)),
) {
  positions_by_char
  |> dict.values
  |> list.flat_map(fn(positions) {
    generate_antinodes(positions, grid_size, compute_antinodes_fn)
  })
  |> set.from_list
  |> set.size
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  let #(grid_size, positions_by_char) = parse_input(input)

  let antinodes_count =
    count_antinodes(grid_size, positions_by_char, compute_antinodes_for_two)
  let antinodes_with_resonance_count =
    count_antinodes(
      grid_size,
      positions_by_char,
      compute_antinodes_with_resonance,
    )

  Ok(utils.AdventOfCodeResult(
    int.to_string(antinodes_count),
    int.to_string(antinodes_with_resonance_count),
  ))
}
