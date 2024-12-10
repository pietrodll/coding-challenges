import gleam/int
import gleam/list
import gleam/option
import gleam/regexp
import gleam/result
import gleam/string
import utils

pub type Operation {
  Mul(left: Int, right: Int)
  Do
  DoNot
}

pub fn extract_operations(input: String) -> Result(List(Operation), String) {
  use re <- result.try(
    regexp.from_string("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")
    |> result.map_error(fn(compile_error) { compile_error.error }),
  )

  regexp.scan(re, input)
  |> list.try_map(fn(match) {
    case match.content {
      "mul(" <> _ -> {
        case match.submatches {
          [option.Some(left_str), option.Some(right_str)] -> {
            use left <- result.try(
              int.parse(left_str)
              |> result.replace_error("cannot parse integer " <> left_str),
            )
            use right <- result.try(
              int.parse(right_str)
              |> result.replace_error("cannot parse integer " <> left_str),
            )
            Ok(Mul(left, right))
          }
          _ -> Error("invalid multiplication " <> string.inspect(match))
        }
      }
      "do()" -> Ok(Do)
      "don't()" -> Ok(DoNot)
      _ -> Error("Unrecognized operation " <> match.content)
    }
  })
}

pub fn compute_result(operations: List(Operation)) -> Int {
  operations
  |> list.map(fn(op) {
    case op {
      Mul(left, right) -> left * right
      _ -> 0
    }
  })
  |> list.fold(0, int.add)
}

pub fn compute_result_including_dos(operations: List(Operation)) -> Int {
  operations
  |> list.fold(#(True, 0), fn(acc, op) {
    let #(is_enabled, result) = acc

    case op {
      Mul(left, right) ->
        case is_enabled {
          True -> #(is_enabled, result + { left * right })
          False -> #(is_enabled, result)
        }
      Do -> #(True, result)
      DoNot -> #(False, result)
    }
  })
  |> fn(tup) { tup.1 }
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use operations <- result.try(extract_operations(input))

  let first_part = compute_result(operations)
  let second_part = compute_result_including_dos(operations)

  Ok(utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  ))
}
