import gleam/int
import gleam/list
import gleam/option
import gleam/pair
import gleam/result
import gleam/set
import utils
import utils_matrix

pub fn parse_input(input: String) -> Result(utils_matrix.Matrix(String), String) {
  utils_matrix.from_text(input, Ok)
}

const directions = [
  utils_matrix.Up,
  utils_matrix.Down,
  utils_matrix.Left,
  utils_matrix.Right,
]

fn get_neighbor(
  matrix: utils_matrix.Matrix(String),
  position: #(Int, Int),
  direction: utils_matrix.Direction,
) -> option.Option(String) {
  utils_matrix.get(matrix, utils_matrix.next_position(position, direction))
}

fn explore_region_rec(
  matrix: utils_matrix.Matrix(String),
  position: #(Int, Int),
  value: String,
  seen_indices: set.Set(#(Int, Int)),
  from: a,
  with: fn(a, #(Int, Int), String) -> a,
) -> #(set.Set(#(Int, Int)), a) {
  directions
  |> list.fold(
    #(seen_indices |> set.insert(position), with(from, position, value)),
    fn(acc, dir) {
      let #(seen_indices_iter, from_iter) = acc

      let next_position = utils_matrix.next_position(position, dir)

      case set.contains(seen_indices_iter, next_position) {
        True -> acc
        False ->
          case utils_matrix.get(matrix, next_position) {
            option.Some(next_value) if next_value == value -> {
              explore_region_rec(
                matrix,
                next_position,
                next_value,
                seen_indices_iter,
                from_iter,
                with,
              )
            }
            _ -> acc
          }
      }
    },
  )
}

fn explore_region(
  matrix: utils_matrix.Matrix(String),
  starting_position: #(Int, Int),
  from: a,
  with: fn(a, #(Int, Int), String) -> a,
) -> a {
  case utils_matrix.get(matrix, starting_position) {
    option.Some(value) ->
      explore_region_rec(
        matrix,
        starting_position,
        value,
        set.new(),
        from,
        with,
      )
      |> pair.second
    option.None -> from
  }
}

fn find_regions_rec(
  matrix: utils_matrix.Matrix(String),
  remaining_indices: set.Set(#(Int, Int)),
  found_regions: List(#(Int, Int)),
) -> List(#(Int, Int)) {
  case set.to_list(remaining_indices) {
    [] -> found_regions
    [position, ..rest] -> {
      let updated_found_regions = [position, ..found_regions]

      // remove all the indices of the region from remaining_indices
      let updated_remaining_indices =
        explore_region(
          matrix,
          position,
          set.from_list(rest),
          fn(indices, pos, _value) { set.delete(indices, pos) },
        )

      find_regions_rec(matrix, updated_remaining_indices, updated_found_regions)
    }
  }
}

/// Returns the list of the matrix's regions, given by a list of points: one for each region.
pub fn find_regions(matrix: utils_matrix.Matrix(String)) -> List(#(Int, Int)) {
  let all_indices =
    set.from_list(utils_matrix.indices(matrix.height, matrix.width))

  find_regions_rec(matrix, all_indices, [])
}

pub fn perimeter_contribution(
  matrix: utils_matrix.Matrix(String),
  position: #(Int, Int),
) -> Int {
  case utils_matrix.get(matrix, position) {
    option.Some(letter) -> {
      directions
      |> list.map(get_neighbor(matrix, position, _))
      |> list.count(fn(maybe_neighbor_letter) {
        case maybe_neighbor_letter {
          option.Some(neighbor_letter) -> neighbor_letter != letter
          option.None -> True
        }
      })
    }
    option.None -> 0
  }
}

fn compute_perimeter(
  matrix: utils_matrix.Matrix(String),
  region_position: #(Int, Int),
) -> Int {
  explore_region(matrix, region_position, 0, fn(acc, position, _value) {
    acc + perimeter_contribution(matrix, position)
  })
}

fn compute_area(
  matrix: utils_matrix.Matrix(String),
  region_position: #(Int, Int),
) -> Int {
  explore_region(matrix, region_position, 0, fn(acc, _position, _value) {
    acc + 1
  })
}

pub fn compute_region_price(
  matrix: utils_matrix.Matrix(String),
  region_position: #(Int, Int),
) -> Int {
  compute_perimeter(matrix, region_position)
  * compute_area(matrix, region_position)
}

/// Represents a fence of the region. Given by the point position and the direction
/// where the fence is compared to the point.
pub type Fence {
  Fence(position: #(Int, Int), direction: utils_matrix.Direction)
}

pub fn get_region_fences(
  matrix: utils_matrix.Matrix(String),
  region_position: #(Int, Int),
) -> set.Set(Fence) {
  matrix
  |> explore_region(region_position, set.new(), fn(acc, pos, letter) {
    directions
    |> list.filter(fn(dir) {
      case get_neighbor(matrix, pos, dir) {
        // If the neighbor is empty, it means we're on the edge of the map
        option.None -> True
        // If the letter is different, we're on the edge of the region
        option.Some(neighbor_letter) -> neighbor_letter != letter
      }
    })
    |> list.map(Fence(pos, _))
    |> list.fold(acc, set.insert)
  })
}

/// Returns a Fence in the same direction of the one passed as parameter, but
/// shifted in the direction passed as parameter.
fn next_fence(fence: Fence, direction: utils_matrix.Direction) -> Fence {
  let next_position = utils_matrix.next_position(fence.position, direction)
  Fence(next_position, fence.direction)
}

/// Moves along the direction passed as parameter as long as they are in the same direction of the starting fence passed as parameter.
fn explore_fence(
  fences: set.Set(Fence),
  fence: Fence,
  direction: utils_matrix.Direction,
  from: a,
  with: fn(a, Fence) -> a,
) -> a {
  let next_fence = next_fence(fence, direction)

  case set.contains(fences, next_fence) {
    True ->
      explore_fence(fences, next_fence, direction, with(from, fence), with)
    False -> with(from, fence)
  }
}

fn explore_side(
  region_fences: set.Set(Fence),
  fence: Fence,
  seen_fences: set.Set(Fence),
) -> set.Set(Fence) {
  case fence.direction {
    utils_matrix.Left | utils_matrix.Right ->
      seen_fences
      |> set.insert(fence)
      |> explore_fence(region_fences, fence, utils_matrix.Up, _, set.insert)
      |> explore_fence(region_fences, fence, utils_matrix.Down, _, set.insert)
    _ ->
      seen_fences
      |> set.insert(fence)
      |> explore_fence(region_fences, fence, utils_matrix.Left, _, set.insert)
      |> explore_fence(region_fences, fence, utils_matrix.Right, _, set.insert)
  }
}

fn compute_number_of_sides_from_fences(region_fences: set.Set(Fence)) -> Int {
  region_fences
  |> set.fold(#(0, set.new()), fn(acc, fence) {
    let #(count, seen_fences) = acc

    case set.contains(seen_fences, fence) {
      True -> acc
      False -> #(count + 1, explore_side(region_fences, fence, seen_fences))
    }
  })
  |> pair.first
}

pub fn compute_number_of_sides(
  matrix: utils_matrix.Matrix(String),
  region_position: #(Int, Int),
) -> Int {
  get_region_fences(matrix, region_position)
  |> compute_number_of_sides_from_fences
}

pub fn compute_region_price_with_discount(
  matrix: utils_matrix.Matrix(String),
  region_position: #(Int, Int),
) -> Int {
  compute_number_of_sides(matrix, region_position)
  * compute_area(matrix, region_position)
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use matrix <- result.map(parse_input(input))

  let regions = find_regions(matrix)

  let first = regions |> list.map(compute_region_price(matrix, _)) |> int.sum
  let second =
    regions
    |> list.map(compute_region_price_with_discount(matrix, _))
    |> int.sum

  utils.AdventOfCodeResult(int.to_string(first), int.to_string(second))
}
