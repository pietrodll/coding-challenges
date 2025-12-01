import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/result
import gleam/string
import utils
import utils_matrix

fn parse_direction(char: String) -> Result(utils_matrix.Direction, String) {
  case char {
    ">" -> Ok(utils_matrix.Right)
    "^" -> Ok(utils_matrix.Up)
    "<" -> Ok(utils_matrix.Left)
    "v" -> Ok(utils_matrix.Down)
    _ -> Error("Cannot parse direction '" <> char <> "'")
  }
}

pub type WarehouseElement {
  None
  Robot
  Wall
  Box
  // The last two values are only used for the 2nd part
  BoxLeft
  BoxRight
}

fn parse_element(char: String) -> Result(WarehouseElement, String) {
  case char {
    "." -> Ok(None)
    "O" -> Ok(Box)
    "@" -> Ok(Robot)
    "#" -> Ok(Wall)
    _ -> Error("Cannot parse '" <> "'")
  }
}

pub fn parse_map(
  map: String,
) -> Result(utils_matrix.Matrix(WarehouseElement), String) {
  utils_matrix.from_text(map, parse_element)
}

pub fn parse_input(input: String) {
  case input |> string.split("\n\n") {
    [first, second] -> {
      use map <- result.try(parse_map(first))
      use moves <- result.try(
        second
        |> string.replace("\n", "")
        |> string.to_graphemes
        |> list.try_map(parse_direction),
      )

      Ok(#(map, moves))
    }
    _ -> Error("Cannot parse input")
  }
}

fn can_move_box(
  map: utils_matrix.Matrix(WarehouseElement),
  position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> Bool {
  let next_position = utils_matrix.next_position(position, direction)

  case utils_matrix.get(map, next_position) {
    option.Some(element) ->
      case element {
        Box -> can_move_box(map, next_position, direction)
        None -> True
        _ -> False
      }
    option.None -> False
  }
}

fn move_boxes(
  map: utils_matrix.Matrix(WarehouseElement),
  replacement: WarehouseElement,
  position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> utils_matrix.Matrix(WarehouseElement) {
  let next_position = utils_matrix.next_position(position, direction)

  case utils_matrix.get(map, next_position) {
    option.Some(element) ->
      case element {
        None ->
          map
          |> utils_matrix.set(position, replacement)
          |> utils_matrix.set(next_position, Box)
        Box ->
          move_boxes(
            map |> utils_matrix.set(position, replacement),
            Box,
            next_position,
            direction,
          )
        _ -> map
      }
    option.None -> map
  }
}

pub fn move_robot_once(
  map: utils_matrix.Matrix(WarehouseElement),
  robot_position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> #(utils_matrix.Matrix(WarehouseElement), #(Int, Int)) {
  let next_position = utils_matrix.next_position(robot_position, direction)

  case utils_matrix.get(map, next_position) {
    option.Some(element) ->
      case element {
        Box ->
          case can_move_box(map, next_position, direction) {
            True -> #(
              move_boxes(
                map |> utils_matrix.set(robot_position, None),
                Robot,
                next_position,
                direction,
              ),
              next_position,
            )
            False -> #(map, robot_position)
          }
        None -> #(
          map
            |> utils_matrix.set(robot_position, None)
            |> utils_matrix.set(next_position, Robot),
          next_position,
        )
        _ -> #(map, robot_position)
      }
    option.None -> #(map, robot_position)
  }
}

fn move_robot_rec(
  map: utils_matrix.Matrix(WarehouseElement),
  robot_position: #(Int, Int),
  directions: List(utils_matrix.Direction),
) -> utils_matrix.Matrix(WarehouseElement) {
  case directions {
    [] -> map
    [first, ..rest] -> {
      let #(next_map, next_robot_position) =
        move_robot_once(map, robot_position, first)
      move_robot_rec(next_map, next_robot_position, rest)
    }
  }
}

fn find_robot(map: utils_matrix.Matrix(WarehouseElement)) {
  utils_matrix.indices(map.height, map.width)
  |> list.find_map(fn(position) {
    utils_matrix.get(map, position)
    |> option.to_result(Nil)
    |> result.try(fn(element) {
      case element {
        Robot -> Ok(position)
        _ -> Error(Nil)
      }
    })
  })
}

pub fn move_robot(
  map: utils_matrix.Matrix(WarehouseElement),
  directions: List(utils_matrix.Direction),
) {
  case find_robot(map) {
    Ok(robot_position) -> move_robot_rec(map, robot_position, directions)
    _ -> map
  }
}

pub fn sum_gps_coordinates(
  map: utils_matrix.Matrix(WarehouseElement),
  directions: List(utils_matrix.Direction),
) {
  let updated_matrix = move_robot(map, directions)

  utils_matrix.indices(updated_matrix.height, updated_matrix.width)
  |> list.filter_map(fn(position) {
    utils_matrix.get(updated_matrix, position)
    |> option.to_result(Nil)
    |> result.try(fn(element) {
      case element {
        Box | BoxLeft -> Ok(position)
        _ -> Error(Nil)
      }
    })
  })
  |> list.map(fn(box_position) { { 100 * box_position.0 } + box_position.1 })
  |> int.sum
}

// ----- Part 2 -----

fn to_wide_positions(pos: #(Int, Int)) -> #(#(Int, Int), #(Int, Int)) {
  let #(i, j) = pos

  #(#(i, 2 * j), #(i, { 2 * j } + 1))
}

fn widen_map(map: utils_matrix.Matrix(WarehouseElement)) {
  let wide_items =
    utils_matrix.indices(map.height, map.width)
    |> list.fold(dict.new(), fn(acc, pos) {
      let #(left, right) = to_wide_positions(pos)

      case utils_matrix.get(map, pos) {
        option.Some(element) ->
          case element {
            Box ->
              acc |> dict.insert(left, BoxLeft) |> dict.insert(right, BoxRight)
            BoxLeft | BoxRight -> acc
            None -> acc |> dict.insert(left, None) |> dict.insert(right, None)
            Robot -> acc |> dict.insert(left, Robot) |> dict.insert(right, None)
            Wall -> acc |> dict.insert(left, Wall) |> dict.insert(right, Wall)
          }
        option.None -> acc
      }
    })

  utils_matrix.Matrix(wide_items, map.height, 2 * map.width)
}

fn can_move_box_wide(
  map: utils_matrix.Matrix(WarehouseElement),
  position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> Bool {
  case direction, utils_matrix.get(map, position) {
    utils_matrix.Up, option.Some(BoxLeft)
    | utils_matrix.Down, option.Some(BoxLeft)
    -> {
      let next_position_left =
        position
        |> utils_matrix.next_position(direction)
      let next_position_right =
        position
        |> utils_matrix.next_position(utils_matrix.Right)
        |> utils_matrix.next_position(direction)

      case
        utils_matrix.get(map, next_position_left),
        utils_matrix.get(map, next_position_right)
      {
        option.Some(next_left), option.Some(next_right) -> {
          case next_left, next_right {
            None, None -> True
            None, BoxLeft ->
              can_move_box_wide(map, next_position_right, direction)
            BoxRight, None ->
              can_move_box_wide(map, next_position_left, direction)
            BoxLeft, BoxRight ->
              can_move_box_wide(map, next_position_left, direction)
            _, _ -> False
          }
        }
        _, _ -> False
      }
    }
    utils_matrix.Up, option.Some(BoxRight)
    | utils_matrix.Down, option.Some(BoxRight)
    -> {
      let next_position_left =
        position
        |> utils_matrix.next_position(utils_matrix.Left)
        |> utils_matrix.next_position(direction)
      let next_position_right =
        position
        |> utils_matrix.next_position(direction)

      case
        utils_matrix.get(map, next_position_left),
        utils_matrix.get(map, next_position_right)
      {
        option.Some(next_left), option.Some(next_right) -> {
          case next_left, next_right {
            None, None -> True
            None, BoxLeft ->
              can_move_box_wide(map, next_position_right, direction)
            BoxRight, None ->
              can_move_box_wide(map, next_position_left, direction)
            BoxLeft, BoxRight ->
              can_move_box_wide(map, next_position_left, direction)
            _, _ -> False
          }
        }
        _, _ -> False
      }
    }
    utils_matrix.Right, option.Some(BoxLeft) -> {
      let next_next_position =
        position
        |> utils_matrix.next_position(direction)
        |> utils_matrix.next_position(direction)

      case utils_matrix.get(map, next_next_position) {
        option.Some(None) -> True
        option.Some(BoxLeft) ->
          can_move_box_wide(map, next_next_position, direction)
        _ -> False
      }
    }
    utils_matrix.Left, option.Some(BoxRight) -> {
      let next_next_position =
        position
        |> utils_matrix.next_position(direction)
        |> utils_matrix.next_position(direction)

      case utils_matrix.get(map, next_next_position) {
        option.Some(None) -> True
        option.Some(BoxRight) ->
          can_move_box_wide(map, next_next_position, direction)
        _ -> False
      }
    }
    // Should not happen
    _, _ -> False
  }
}

fn move_boxes_wide(
  map: utils_matrix.Matrix(WarehouseElement),
  replacement: WarehouseElement,
  position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> utils_matrix.Matrix(WarehouseElement) {
  let next_position = position |> utils_matrix.next_position(direction)

  case map |> utils_matrix.get(next_position) {
    option.None -> map
    option.Some(next_element) -> {
      todo
    }
  }
}

pub fn run(input: String) {
  use #(map, directions) <- result.try(parse_input(input))
  let first_part = sum_gps_coordinates(map, directions)

  Ok(utils.AdventOfCodeResult(int.to_string(first_part), ""))
}
