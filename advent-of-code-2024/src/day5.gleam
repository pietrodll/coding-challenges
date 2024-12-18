import gleam/dict
import gleam/int
import gleam/list
import gleam/option
import gleam/order
import gleam/pair
import gleam/result
import gleam/set
import gleam/string
import utils

fn parse_ordering_rules(
  ordering_rules: String,
) -> Result(List(#(Int, Int)), String) {
  ordering_rules
  |> string.split("\n")
  |> list.try_map(fn(line) {
    case line |> string.split("|") {
      [first_str, second_str] -> {
        use first <- result.try(utils.parse_int(first_str))
        use second <- result.try(utils.parse_int(second_str))
        Ok(#(first, second))
      }
      _ -> Error("could not parse page numbers from line " <> line)
    }
  })
}

fn parse_update_pages(update_pages: String) -> Result(List(List(Int)), String) {
  update_pages
  |> string.split("\n")
  |> list.try_map(fn(line) {
    line |> string.split(",") |> list.try_map(utils.parse_int)
  })
}

pub fn parse_input(
  input: String,
) -> Result(#(List(#(Int, Int)), List(List(Int))), String) {
  case input |> string.trim |> string.split("\n\n") {
    [first_part, second_part] -> {
      use ordering_rules <- result.try(parse_ordering_rules(first_part))
      use update_pages <- result.try(parse_update_pages(second_part))

      Ok(#(ordering_rules, update_pages))
    }
    _ -> Error("could not parse input " <> input)
  }
}

pub fn generate_order(
  ordering_rules: List(#(Int, Int)),
) -> fn(Int, Int) -> order.Order {
  let lower =
    ordering_rules
    |> list.group(pair.first)
    |> dict.map_values(fn(_, values) {
      values |> list.map(pair.second) |> set.from_list
    })
  let greater =
    ordering_rules
    |> list.group(pair.second)
    |> dict.map_values(fn(_, values) {
      values |> list.map(pair.first) |> set.from_list
    })

  fn(a, b) {
    lower
    |> dict.get(a)
    |> result.then(fn(values) {
      case values |> set.contains(b) {
        True -> Ok(order.Lt)
        False -> Error(Nil)
      }
    })
    |> result.lazy_or(fn() {
      lower
      |> dict.get(b)
      |> result.then(fn(values) {
        case values |> set.contains(a) {
          True -> Ok(order.Gt)
          False -> Error(Nil)
        }
      })
    })
    |> result.lazy_or(fn() {
      greater
      |> dict.get(a)
      |> result.then(fn(values) {
        case values |> set.contains(b) {
          True -> Ok(order.Gt)
          False -> Error(Nil)
        }
      })
    })
    |> result.lazy_or(fn() {
      greater
      |> dict.get(b)
      |> result.then(fn(values) {
        case values |> set.contains(a) {
          True -> Ok(order.Lt)
          False -> Error(Nil)
        }
      })
    })
    |> result.unwrap(order.Eq)
  }
}

pub fn is_correctly_ordered(
  order_func: fn(Int, Int) -> order.Order,
  update: List(Int),
) {
  update == list.sort(update, order_func)
}

pub fn find_middle_rec(fast: List(a), slow: List(a)) -> option.Option(a) {
  case fast, slow {
    [], [] -> option.None
    [], [first, ..] -> option.Some(first)
    [_], [first, ..] -> option.Some(first)
    [_, _, ..rest1], [_, ..rest2] -> find_middle_rec(rest1, rest2)
    _, _ -> option.None
  }
}

pub fn find_middle(list: List(a)) -> option.Option(a) {
  find_middle_rec(list, list)
}

pub fn first_part(
  order_func: fn(Int, Int) -> order.Order,
  update_pages: List(List(Int)),
) -> Int {
  update_pages
  |> list.filter(fn(update) { is_correctly_ordered(order_func, update) })
  |> list.map(find_middle)
  |> option.values
  |> list.fold(0, int.add)
}

pub fn second_part(
  order_func: fn(Int, Int) -> order.Order,
  update_pages: List(List(Int)),
) -> Int {
  update_pages
  |> list.filter_map(fn(update) {
    let reordered = list.sort(update, order_func)
    case update == reordered {
      True -> Error(Nil)
      False -> find_middle(reordered) |> option.to_result(Nil)
    }
  })
  |> list.fold(0, int.add)
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use #(ordering_rules, update_pages) <- result.try(parse_input(input))
  let order_func = generate_order(ordering_rules)

  let first_part = first_part(order_func, update_pages)
  let second_part = second_part(order_func, update_pages)

  Ok(utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  ))
}
