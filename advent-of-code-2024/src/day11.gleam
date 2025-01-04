import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/order
import gleam/result
import gleam/string
import utils

pub fn parse_input(input: String) -> Result(List(Int), String) {
  input |> string.trim |> string.split(" ") |> list.try_map(utils.parse_int)
}

fn count_digits_rec(number: Int, count: Int) -> Int {
  case int.compare(number, 10) {
    order.Lt -> count + 1
    order.Gt | order.Eq -> count_digits_rec(number / 10, count + 1)
  }
}

fn count_digits(number: Int) -> Int {
  count_digits_rec(number, 0)
}

fn power_rec(number: Int, power: Int, acc: Int) -> Int {
  case power {
    0 -> acc
    _ -> power_rec(number, power - 1, acc * number)
  }
}

fn power(number: Int, power: Int) -> Int {
  power_rec(number, power, 1)
}

fn split_digits(number: Int) -> option.Option(#(Int, Int)) {
  let num_digits = count_digits(number)

  case int.is_even(num_digits) {
    True -> {
      let splitting_power = power(10, num_digits / 2)
      option.Some(#(number / splitting_power, number % splitting_power))
    }
    False -> option.None
  }
}

pub fn blink_stone(stone: Int) -> List(Int) {
  case stone {
    0 -> [1]
    _ ->
      case split_digits(stone) {
        option.Some(#(first, second)) -> [first, second]
        option.None -> [stone * 2024]
      }
  }
}

fn increment_stone(
  count_by_stone: dict.Dict(Int, Int),
  stone: Int,
) -> dict.Dict(Int, Int) {
  increment_stone_by(count_by_stone, stone, 1)
}

fn increment_stone_by(
  count_by_stone: dict.Dict(Int, Int),
  stone: Int,
  increment: Int,
) -> dict.Dict(Int, Int) {
  count_by_stone
  |> dict.upsert(stone, fn(previous) { option.unwrap(previous, 0) + increment })
}

pub fn blink_count(count_by_stone: dict.Dict(Int, Int)) -> dict.Dict(Int, Int) {
  count_by_stone
  |> dict.fold(dict.new(), fn(acc, stone, current_count) {
    blink_stone(stone)
    |> list.fold(acc, fn(acc1, new_stone) {
      increment_stone_by(acc1, new_stone, current_count)
    })
  })
}

pub fn count_stones_after_blink(stones: List(Int), blinks: Int) -> Int {
  stones
  |> list.fold(dict.new(), increment_stone)
  |> utils.call_multiple_times(blink_count, blinks)
  |> dict.values
  |> int.sum
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use stones <- result.map(parse_input(input))

  let first_part = count_stones_after_blink(stones, 25)
  let second_part = count_stones_after_blink(stones, 75)

  utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  )
}
