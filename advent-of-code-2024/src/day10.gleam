import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/set
import gleam/string
import utils

pub type Grid(a) {
  Grid(items: dict.Dict(#(Int, Int), a), Int, Int)
}

pub fn get(grid: Grid(a), pos: #(Int, Int)) -> option.Option(a) {
  grid.items |> dict.get(pos) |> option.from_result
}

pub type Move {
  Up
  Down
  Left
  Right
}

pub fn next(move: Move, pos: #(Int, Int)) -> #(Int, Int) {
  let #(i, j) = pos

  case move {
    Up -> #(i - 1, j)
    Down -> #(i + 1, j)
    Left -> #(i, j - 1)
    Right -> #(i, j + 1)
  }
}

pub fn parse_input(input: String) -> Result(Grid(Int), String) {
  let lines = input |> string.trim |> string.split("\n")
  let height = list.length(lines)
  let width = list.first(lines) |> result.map(string.length) |> result.unwrap(0)

  use elements <- result.map(
    lines
    |> list.index_map(fn(line, i) {
      line
      |> string.to_graphemes
      |> list.index_map(fn(char, j) {
        use parsed_char <- result.map(utils.parse_int(char))
        #(#(i, j), parsed_char)
      })
    })
    |> list.flatten
    |> result.all,
  )

  let matrix = dict.from_list(elements)

  Grid(matrix, height, width)
}

const allowed_moves = [Up, Down, Left, Right]

fn next_positions(
  grid: Grid(Int),
  pos: #(Int, Int),
  value: Int,
) -> List(#(Int, Int)) {
  allowed_moves
  |> list.filter_map(fn(move) {
    let next = next(move, pos)
    use next_value <- result.try(grid |> get(next) |> option.to_result(Nil))

    case next_value == value + 1 {
      True -> Ok(next)
      False -> Error(Nil)
    }
  })
}

fn summits_reachable_from(
  grid: Grid(Int),
  pos: #(Int, Int),
) -> set.Set(#(Int, Int)) {
  case grid |> get(pos) {
    option.None -> set.new()
    option.Some(value) if value == 9 -> set.new() |> set.insert(pos)
    option.Some(value) -> {
      next_positions(grid, pos, value)
      |> list.map(summits_reachable_from(grid, _))
      |> list.fold(set.new(), set.union)
    }
  }
}

pub fn score_starting_from(grid: Grid(Int), pos: #(Int, Int)) -> Int {
  grid |> summits_reachable_from(pos) |> set.size
}

pub fn find_trailheads(grid: Grid(Int)) -> List(#(Int, Int)) {
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

pub fn total_score(grid: Grid(Int)) -> Int {
  grid
  |> find_trailheads
  |> list.map(score_starting_from(grid, _))
  |> int.sum
}

pub fn rating_starting_from(grid: Grid(Int), pos: #(Int, Int)) -> Int {
  case grid |> get(pos) {
    option.None -> 0
    option.Some(value) if value == 9 -> 1
    option.Some(value) -> {
      next_positions(grid, pos, value)
      |> list.map(rating_starting_from(grid, _))
      |> int.sum
    }
  }
}

pub fn total_rating(grid: Grid(Int)) -> Int {
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
