import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/string
import utils

pub type Letter {
  X
  M
  A
  S
}

fn parse_letter(char: String) -> Result(Letter, String) {
  case char {
    "X" -> Ok(X)
    "M" -> Ok(M)
    "A" -> Ok(A)
    "S" -> Ok(S)
    _ -> Error("cannot parse letter \"" <> char <> "\"")
  }
}

pub type LetterMatrix =
  dict.Dict(Int, dict.Dict(Int, Letter))

fn to_table(list: List(a)) -> dict.Dict(Int, a) {
  list |> list.index_map(fn(item, index) { #(index, item) }) |> dict.from_list
}

pub fn parse_input(input: String) -> Result(LetterMatrix, String) {
  input
  |> string.trim
  |> string.split("\n")
  |> list.try_map(fn(line) {
    line
    |> string.to_graphemes
    |> list.try_map(parse_letter)
    |> result.map(to_table)
  })
  |> result.map(to_table)
}

fn matrix_indices(size: Int) -> List(#(Int, Int)) {
  list.range(0, size - 1)
  |> list.flat_map(fn(i) {
    list.range(0, size - 1) |> list.map(fn(j) { #(i, j) })
  })
}

fn matrix_get(matrix: LetterMatrix, i: Int, j: Int) -> option.Option(Letter) {
  matrix
  |> dict.get(i)
  |> result.then(fn(line) { line |> dict.get(j) })
  |> option.from_result
}

fn check_xmas(
  matrix: LetterMatrix,
  i0: Int,
  j0: Int,
  direction: #(Int, Int),
) -> Bool {
  let #(delta_i, delta_j) = direction

  let #(is_valid, _, _) =
    [X, M, A, S]
    |> list.fold_until(#(True, i0, j0), fn(position, stage) {
      let #(_, i, j) = position
      matrix_get(matrix, i, j)
      |> option.map(fn(letter) {
        case letter == stage {
          True -> list.Continue(#(True, i + delta_i, j + delta_j))
          False -> list.Stop(#(False, -1, -1))
        }
      })
      |> option.unwrap(list.Stop(#(False, -1, -1)))
    })

  is_valid
}

fn count_xmas_from(matrix: LetterMatrix, i0: Int, j0: Int) {
  let directions = [
    #(1, 0),
    #(1, 1),
    #(0, 1),
    #(-1, 1),
    #(-1, 0),
    #(-1, -1),
    #(0, -1),
    #(1, -1),
  ]

  directions
  |> list.count(fn(direction) { check_xmas(matrix, i0, j0, direction) })
}

pub fn count_xmas(matrix: LetterMatrix) -> Int {
  let size = dict.size(matrix)

  matrix_indices(size)
  |> list.map(fn(indices) { count_xmas_from(matrix, indices.0, indices.1) })
  |> list.fold(0, int.add)
}

fn is_x_mas(matrix: LetterMatrix, i: Int, j: Int) -> Bool {
  [
    matrix_get(matrix, i, j),
    matrix_get(matrix, i - 1, j - 1),
    matrix_get(matrix, i - 1, j + 1),
    matrix_get(matrix, i + 1, j - 1),
    matrix_get(matrix, i + 1, j + 1),
  ]
  |> option.all
  |> option.map(fn(letters) {
    let assert [center, top_left, top_right, bottom_left, bottom_right] =
      letters
    center == A
    && {
      { top_left == M && bottom_right == S }
      || { top_left == S && bottom_right == M }
    }
    && {
      { top_right == M && bottom_left == S }
      || { top_right == S && bottom_left == M }
    }
  })
  |> option.unwrap(False)
}

pub fn count_x_mas(matrix: LetterMatrix) -> Int {
  matrix_indices(dict.size(matrix))
  |> list.count(fn(indices) { is_x_mas(matrix, indices.0, indices.1) })
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use letter_matrix <- result.map(parse_input(input))

  let first_part = count_xmas(letter_matrix)
  let second_part = count_x_mas(letter_matrix)

  utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  )
}
