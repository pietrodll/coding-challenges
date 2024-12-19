import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/set
import gleam/string
import utils

pub type GuardDirection {
  Up
  Down
  Left
  Right
}

/// Parses a line into a triplet with:
///   - length
///   - list of obstacles positions
///   - optional position and direction of the guard
fn parse_line(
  line: String,
) -> #(Int, List(Int), option.Option(#(Int, GuardDirection))) {
  let length = string.length(line)

  let #(obstacles, guard) =
    line
    |> string.to_graphemes
    |> list.index_fold(#(list.new(), option.None), fn(acc, elem, idx) {
      let #(obs, grd) = acc

      case elem {
        "#" -> #([idx, ..obs], grd)
        "^" -> #(obs, option.Some(#(idx, Up)))
        "v" -> #(obs, option.Some(#(idx, Down)))
        ">" -> #(obs, option.Some(#(idx, Right)))
        "<" -> #(obs, option.Some(#(idx, Left)))
        _ -> acc
      }
    })

  #(length, obstacles, guard)
}

/// Parses the input into a triplet with
///   - size of the grid
///   - initial position
///   - list of obstacle coordinates
pub fn parse_input(
  input: String,
) -> Result(
  #(#(Int, Int), #(Int, Int, GuardDirection), set.Set(#(Int, Int))),
  String,
) {
  let #(height, length, obstacles, maybe_guard) =
    input
    |> string.trim
    |> string.split("\n")
    |> list.index_fold(#(0, -1, [], option.None), fn(acc, line, idx) {
      let #(line_length, line_obstacles, line_maybe_guard) = parse_line(line)
      let #(height, _, obstacles, maybe_guard) = acc

      #(
        height + 1,
        line_length,
        line_obstacles
          |> list.map(fn(j) { #(idx, j) })
          |> list.append(obstacles),
        option.or(
          maybe_guard,
          line_maybe_guard
            |> option.map(fn(line_guard) { #(idx, line_guard.0, line_guard.1) }),
        ),
      )
    })

  use guard <- result.try(
    maybe_guard |> option.to_result("guard position not found"),
  )

  Ok(#(#(height, length), guard, set.from_list(obstacles)))
}

fn turn_right(direction: GuardDirection) -> GuardDirection {
  case direction {
    Up -> Right
    Right -> Down
    Down -> Left
    Left -> Up
  }
}

fn next_position(
  grid_size: #(Int, Int),
  guard: #(Int, Int, GuardDirection),
) -> option.Option(#(Int, Int)) {
  let #(x, y, direction) = guard
  let #(delta_x, delta_y) = case direction {
    Up -> #(-1, 0)
    Right -> #(0, 1)
    Down -> #(1, 0)
    Left -> #(0, -1)
  }

  let next_x = x + delta_x
  let next_y = y + delta_y

  case
    next_x < grid_size.0 && next_x >= 0 && next_y < grid_size.1 && next_y >= 0
  {
    True -> option.Some(#(next_x, next_y))
    False -> option.None
  }
}

fn count_guard_positions_rec(
  seen_positions: set.Set(#(Int, Int)),
  grid_size: #(Int, Int),
  guard: #(Int, Int, GuardDirection),
  obstacles: set.Set(#(Int, Int)),
) -> set.Set(#(Int, Int)) {
  let #(x, y, direction) = guard

  case next_position(grid_size, guard) {
    option.Some(#(next_x, next_y)) -> {
      case set.contains(obstacles, #(next_x, next_y)) {
        False ->
          count_guard_positions_rec(
            seen_positions |> set.insert(#(next_x, next_y)),
            grid_size,
            #(next_x, next_y, direction),
            obstacles,
          )
        True ->
          count_guard_positions_rec(
            seen_positions,
            grid_size,
            #(x, y, turn_right(direction)),
            obstacles,
          )
      }
    }
    option.None -> seen_positions
  }
}

pub fn count_guard_positions(
  grid_size: #(Int, Int),
  guard: #(Int, Int, GuardDirection),
  obstacles: set.Set(#(Int, Int)),
) -> Int {
  let #(x, y, _) = guard

  count_guard_positions_rec(
    set.from_list([#(x, y)]),
    grid_size,
    guard,
    obstacles,
  )
  |> set.size
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use #(grid_size, guard, obstacles) <- result.try(parse_input(input))

  let seen_positions = count_guard_positions(grid_size, guard, obstacles)

  Ok(utils.AdventOfCodeResult(int.to_string(seen_positions), ""))
}
