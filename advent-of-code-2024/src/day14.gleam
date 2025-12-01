import gleam/dict
import gleam/int
import gleam/list
import gleam/result
import gleam/string
import utils

pub type Robot {
  Robot(position: #(Int, Int), speed: #(Int, Int))
}

fn parse_coords(str: String) -> Result(#(Int, Int), String) {
  case str {
    "p=" <> coords | "v=" <> coords -> {
      case string.split(coords, ",") {
        [x_str, y_str] -> {
          use x <- result.try(utils.parse_int(x_str))
          use y <- result.try(utils.parse_int(y_str))
          Ok(#(x, y))
        }
        _ -> Error("Cannot parse coordinates '" <> coords <> "'")
      }
    }
    _ -> Error("Cannot parse '" <> str <> "'")
  }
}

fn parse_line(line: String) -> Result(Robot, String) {
  case line |> string.split(" ") {
    [position_str, speed_str] -> {
      use position <- result.try(parse_coords(position_str))
      use speed <- result.try(parse_coords(speed_str))
      Ok(Robot(position, speed))
    }
    _ -> Error("Cannot parse line '" <> line <> "'")
  }
}

pub fn parse_input(input: String) -> Result(List(Robot), String) {
  input |> string.trim |> string.split("\n") |> list.try_map(parse_line)
}

/// Computes the position of the robot after the number of steps passed as parameter. Assumes
/// that the position is valid (i.e. within the bounds of the grid)
pub fn move(
  robot: Robot,
  steps: Int,
  grid_height: Int,
  grid_width: Int,
) -> #(Int, Int) {
  let #(x, y) = robot.position
  let #(vx, vy) = robot.speed

  let x_factor = steps % grid_width
  let y_factor = steps % grid_height

  // Modulo can return negative values, make sure everything is positive
  let vx_mod = { vx % grid_width } + grid_width
  let vy_mod = { vy % grid_height } + grid_height

  let new_x = { x + { vx_mod * x_factor } } % grid_width
  let new_y = { y + { vy_mod * y_factor } } % grid_height

  #(new_x, new_y)
}

pub type Quadrant {
  TopLeft
  TopRight
  BottomLeft
  BottomRight
  Middle
}

pub fn quadrant(x: Int, y: Int, grid_height: Int, grid_width: Int) -> Quadrant {
  let middle_width = grid_width / 2
  let middle_height = grid_height / 2

  case x == middle_width || y == middle_height {
    True -> Middle
    False ->
      case x < middle_width {
        True ->
          case y < middle_height {
            True -> TopLeft
            False -> BottomLeft
          }
        False ->
          case y < middle_height {
            True -> TopRight
            False -> BottomRight
          }
      }
  }
}

fn count_by_quadrant(
  positions: List(#(Int, Int)),
  grid_height: Int,
  grid_width: Int,
) -> dict.Dict(Quadrant, Int) {
  positions
  |> list.map(fn(pos) { quadrant(pos.0, pos.1, grid_height, grid_width) })
  |> utils.counter
}

pub fn safety_factor(
  robots: List(Robot),
  steps: Int,
  grid_height: Int,
  grid_width: Int,
) -> Int {
  robots
  |> list.map(move(_, steps, grid_height, grid_width))
  |> count_by_quadrant(grid_height, grid_width)
  |> dict.delete(Middle)
  |> dict.values
  |> int.product
}

pub fn run(input: String) {
  use robots <- result.map(parse_input(input))

  let grid_height = 103
  let grid_width = 101

  let first_part = safety_factor(robots, 100, grid_height, grid_width)

  utils.AdventOfCodeResult(int.to_string(first_part), "")
}
