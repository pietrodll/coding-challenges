import gleam/dict
import gleam/list
import gleam/option
import gleam/pair
import gleam/result
import gleam/string

pub type Matrix(a) {
  Matrix(items: dict.Dict(#(Int, Int), a), height: Int, width: Int)
}

pub fn get(grid: Matrix(a), pos: #(Int, Int)) -> option.Option(a) {
  grid.items |> dict.get(pos) |> option.from_result
}

pub fn set(grid: Matrix(a), pos: #(Int, Int), value: a) -> Matrix(a) {
  let #(i, j) = pos

  case 0 <= i && i < grid.height && 0 <= j && j < grid.width {
    True ->
      Matrix(grid.items |> dict.insert(pos, value), grid.height, grid.width)
    False -> grid
  }
}

pub fn from_text(
  text: String,
  parse_grapheme: fn(String) -> Result(a, b),
) -> Result(Matrix(a), b) {
  let lines = text |> string.trim |> string.split("\n")
  let height = list.length(lines)
  let width =
    lines |> list.first |> result.map(string.length) |> result.unwrap(0)

  use elements <- result.map(
    lines
    |> list.index_map(fn(line, i) {
      line
      |> string.to_graphemes
      |> list.index_map(fn(char, j) {
        use parsed_char <- result.map(parse_grapheme(char))
        #(#(i, j), parsed_char)
      })
    })
    |> list.flatten
    |> result.all,
  )

  let matrix = dict.from_list(elements)

  Matrix(matrix, height, width)
}

pub fn indices(height: Int, width: Int) -> List(#(Int, Int)) {
  list.range(0, height - 1)
  |> list.flat_map(fn(i) {
    list.range(0, width - 1) |> list.map(pair.new(i, _))
  })
}

pub type Direction {
  Up
  Down
  Left
  Right
  UpLeft
  UpRight
  DownLeft
  DownRight
}

fn delta(direction: Direction) -> #(Int, Int) {
  case direction {
    Down -> #(1, 0)
    DownLeft -> #(1, -1)
    DownRight -> #(1, 1)
    Left -> #(0, -1)
    Right -> #(0, 1)
    Up -> #(-1, 0)
    UpLeft -> #(-1, -1)
    UpRight -> #(-1, 1)
  }
}

pub fn next_position(position: #(Int, Int), direction: Direction) -> #(Int, Int) {
  let #(i, j) = position
  let #(delta_i, delta_j) = delta(direction)

  #(i + delta_i, j + delta_j)
}
