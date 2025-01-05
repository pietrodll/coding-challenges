import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/set
import utils
import utils_matrix

pub fn parse_input(input: String) -> Result(utils_matrix.Matrix(Int), String) {
  utils_matrix.from_text(input, utils.parse_int)
}

const allowed_moves = [
  utils_matrix.Up,
  utils_matrix.Down,
  utils_matrix.Left,
  utils_matrix.Right,
]

fn next_positions(
  grid: utils_matrix.Matrix(Int),
  pos: #(Int, Int),
  value: Int,
) -> List(#(Int, Int)) {
  allowed_moves
  |> list.filter_map(fn(move) {
    let next = utils_matrix.next_position(pos, move)
    use next_value <- result.try(
      grid |> utils_matrix.get(next) |> option.to_result(Nil),
    )

    case next_value == value + 1 {
      True -> Ok(next)
      False -> Error(Nil)
    }
  })
}

fn summits_reachable_from(
  grid: utils_matrix.Matrix(Int),
  pos: #(Int, Int),
) -> set.Set(#(Int, Int)) {
  case grid |> utils_matrix.get(pos) {
    option.None -> set.new()
    option.Some(value) if value == 9 -> set.new() |> set.insert(pos)
    option.Some(value) -> {
      next_positions(grid, pos, value)
      |> list.map(summits_reachable_from(grid, _))
      |> list.fold(set.new(), set.union)
    }
  }
}

pub fn score_starting_from(
  grid: utils_matrix.Matrix(Int),
  pos: #(Int, Int),
) -> Int {
  grid |> summits_reachable_from(pos) |> set.size
}

pub fn find_trailheads(grid: utils_matrix.Matrix(Int)) -> List(#(Int, Int)) {
  grid.items
  |> dict.to_list
  |> list.filter_map(fn(pos_and_value) {
    let #(pos, value) = pos_and_value
    case value == 0 {
      True -> Ok(pos)
      False -> Error(Nil)
    }
  })
}

pub fn total_score(grid: utils_matrix.Matrix(Int)) -> Int {
  grid
  |> find_trailheads
  |> list.map(score_starting_from(grid, _))
  |> int.sum
}

pub fn rating_starting_from(
  grid: utils_matrix.Matrix(Int),
  pos: #(Int, Int),
) -> Int {
  case grid |> utils_matrix.get(pos) {
    option.None -> 0
    option.Some(value) if value == 9 -> 1
    option.Some(value) -> {
      next_positions(grid, pos, value)
      |> list.map(rating_starting_from(grid, _))
      |> int.sum
    }
  }
}

pub fn total_rating(grid: utils_matrix.Matrix(Int)) -> Int {
  grid
  |> find_trailheads
  |> list.map(rating_starting_from(grid, _))
  |> int.sum
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use grid <- result.map(parse_input(input))

  let total_score = total_score(grid)
  let total_rating = total_rating(grid)

  utils.AdventOfCodeResult(
    int.to_string(total_score),
    int.to_string(total_rating),
  )
}
