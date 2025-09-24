import gleam/float
import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/string
import utils
import utils_math

pub type Coords {
  Coords(x: Int, y: Int)
}

pub type Game {
  Game(a: Coords, b: Coords, prize: Coords)
}

fn parse_line(line: String) -> Result(Coords, String) {
  use #(prefix, coords) <- result.try(
    line
    |> string.trim
    |> string.split_once(": ")
    |> result.replace_error("Cannot parse '" <> line <> "'"),
  )

  case prefix {
    "Button A" | "Button B" | "Prize" -> {
      case coords |> string.split(", ") {
        [x_str, y_str] -> {
          use x <- result.try(x_str |> string.drop_start(2) |> utils.parse_int)
          use y <- result.try(y_str |> string.drop_start(2) |> utils.parse_int)
          Ok(Coords(x, y))
        }
        _ -> Error("Cannot parse coordinates '" <> coords <> "'")
      }
    }
    _ -> Error("Cannot parse '" <> line <> "'")
  }
}

fn parse_game(game_str: String) -> Result(Game, String) {
  case game_str |> string.split("\n") {
    [a_str, b_str, prize_str] -> {
      use a <- result.try(parse_line(a_str))
      use b <- result.try(parse_line(b_str))
      use prize <- result.try(parse_line(prize_str))
      Ok(Game(a, b, prize))
    }
    _ -> Error("Cannot parse '" <> game_str <> "'")
  }
}

pub fn parse_input(content: String) -> Result(List(Game), String) {
  content |> string.trim |> string.split("\n\n") |> list.try_map(parse_game)
}

/// Folds a range of integers using the provided reducer function. The upper bound is exclusive
fn fold_range(from: Int, to: Int, init: a, with: fn(a, Int) -> a) -> a {
  case from >= to {
    True -> init
    False -> fold_range(from + 1, to, with(init, from), with)
  }
}

/// Returns the number of time the A and B buttons need to be pressed to win the
/// game. Returns an empty option if the game cannot be won.
pub fn game_cost(game: Game) -> option.Option(Int) {
  let Game(a, b, prize) = game

  // We're looking for (a, b) such that:
  // a.x * a + b.x * b = prize.x
  // a.y * a + b.y * b = prize.y

  // Considering that these are line equations, we're looking for the intersections
  // of those lines with integer coordinates
  // Three cases:
  // 1. The lines are parallel and non equal -> no solution
  // 2. The lines intersect: solution if the intersection has integer coordinates
  // 3. The lines are equal: it boils down to solving a Diophantine equation

  case { a.x * b.y } - { a.y - b.x } {
    // Lines are parallel
    0 ->
      case { a.x * prize.y } - { a.y - prize.x } {
        // Last member is proportional with same ratio: the lines are equal
        0 -> {
          // Solve diophantine equation, only the X equation is enough
          let gcd = utils_math.gcd(a.x, b.x)
          // The right member must be divisible by the GCD the coefficients
          case prize.x % gcd == 0 {
            True -> {
              let factor = prize.x / gcd
              let solution = utils_math.bezout_identity(a.x, b.x)
              let a0 = solution.0 * factor
              let b0 = solution.1 * factor

              // Solutions for the X direction are in the form (for any k)
              // a = a0 + k * b.x
              // b = b0 - k * a.x

              // a and b must be positive
              // a >= 0 -> k * b.x >= -a0 -> k >= -a0 / b.x -> k >= ceil(-a0 / b.x)
              // b >= 0 -> b0 >= k * a.x -> k <= b0 / a.x -> k <= floor(b0 / a.x)
              let from =
                float.truncate(float.ceiling(
                  { -1.0 *. int.to_float(a0) } /. int.to_float(b.x),
                ))
              use to <- option.then(
                int.floor_divide(b0, a.x) |> option.from_result,
              )

              fold_range(from, to + 1, option.None, fn(maybe_current_min, k) {
                let ak = a0 + { k * b.x }
                let bk = b0 - { k * a.x }
                case { 3 * ak } + bk {
                  cost if cost >= 0 ->
                    case maybe_current_min {
                      option.Some(current_min) ->
                        option.Some(int.min(current_min, cost))
                      option.None -> option.Some(cost)
                    }
                  _ -> maybe_current_min
                }
              })
            }
            False -> option.None
          }
        }
        // The two equations are not equivalent: lines are parallel and different, no solution
        _ -> option.None
      }
    // Lines intersect
    _ -> {
      let denominator = { a.x * b.y } - { a.y * b.x }
      let a_numerator = { b.y * prize.x } - { b.x * prize.y }
      let b_numerator = { a.x * prize.y } - { a.y * prize.x }

      // Check if coordinates are integers
      case a_numerator % denominator == 0 && b_numerator % denominator == 0 {
        True -> {
          // Single solution, the intersection point
          let a_count = a_numerator / denominator
          let b_count = b_numerator / denominator

          case a_count >= 0 && b_count >= 0 {
            True -> option.Some({ 3 * a_count } + b_count)
            False -> option.None
          }
        }
        False -> option.None
      }
    }
  }
}

fn shift(game: Game) -> Game {
  Game(
    game.a,
    game.b,
    Coords(game.prize.x + 10_000_000_000_000, game.prize.y + 10_000_000_000_000),
  )
}

fn total_cost(games: List(Game)) -> Int {
  games |> list.map(game_cost) |> option.values |> int.sum
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use parsed_input <- result.map(parse_input(input))

  let first_part = parsed_input |> total_cost

  let second_part = parsed_input |> list.map(shift) |> total_cost

  utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  )
}
