import gleam/int
import gleam/list
import gleam/option
import gleam/pair
import gleam/result
import utils
import utils_matrix

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
  utils_matrix.Matrix(Letter)

pub fn parse_input(input: String) -> Result(LetterMatrix, String) {
  utils_matrix.from_text(input, parse_letter)
}

fn check_xmas(
  matrix: LetterMatrix,
  starting_position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> Bool {
  [X, M, A, S]
  |> list.fold_until(#(True, starting_position), fn(acc, stage) {
    let #(_, position) = acc

    case utils_matrix.get(matrix, position) {
      option.Some(letter) if letter == stage ->
        list.Continue(#(True, utils_matrix.next_position(position, direction)))
      _ -> list.Stop(#(False, #(-1, -1)))
    }
  })
  |> pair.first
}

const directions = [
  utils_matrix.Up,
  utils_matrix.Down,
  utils_matrix.Left,
  utils_matrix.Right,
  utils_matrix.UpLeft,
  utils_matrix.UpRight,
  utils_matrix.DownLeft,
  utils_matrix.DownRight,
]

fn count_xmas_from(matrix: LetterMatrix, position: #(Int, Int)) {
  directions
  |> list.count(check_xmas(matrix, position, _))
}

pub fn count_xmas(matrix: LetterMatrix) -> Int {
  utils_matrix.indices(matrix.height, matrix.width)
  |> list.map(count_xmas_from(matrix, _))
  |> int.sum
}

fn is_x_mas(matrix: LetterMatrix, position: #(Int, Int)) -> Bool {
  [
    position,
    utils_matrix.next_position(position, utils_matrix.UpLeft),
    utils_matrix.next_position(position, utils_matrix.UpRight),
    utils_matrix.next_position(position, utils_matrix.DownLeft),
    utils_matrix.next_position(position, utils_matrix.DownRight),
  ]
  |> list.map(utils_matrix.get(matrix, _))
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
  utils_matrix.indices(matrix.height, matrix.width)
  |> list.count(is_x_mas(matrix, _))
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
