import gleam/int
import gleam/list
import gleam/order
import gleam/pair
import gleam/result
import gleam/string
import utils

pub type Block {
  File(id: Int, size: Int)
  Empty(size: Int)
}

pub fn parse_input(input: String) -> Result(List(Block), String) {
  let digits = input |> string.trim |> string.to_graphemes
  use disk_map <- result.map(list.try_map(digits, utils.parse_int))

  disk_map
  |> list.index_map(fn(size, index) {
    case int.is_even(index) {
      True -> File(index / 2, size)
      False -> Empty(size)
    }
  })
}

fn add_file(compressed: List(Block), id: Int, size: Int) -> List(Block) {
  case compressed {
    [File(first_id, first_size), ..rest] if first_id == id -> [
      File(first_id, first_size + size),
      ..rest
    ]
    _ -> [File(id, size), ..compressed]
  }
}

fn add_empty(compressed: List(Block), size: Int) -> List(Block) {
  case compressed {
    [Empty(s), ..rest] -> [Empty(s + size), ..rest]
    _ -> [Empty(size), ..compressed]
  }
}

fn compact_files_rec(
  compressed: List(Block),
  to_compress: List(Block),
  to_take: List(#(Int, Int)),
  uncompressed_size: Int,
) -> List(Block) {
  case to_compress, to_take, uncompressed_size {
    [], _, _ | _, [], _ | _, _, 0 -> compressed
    [first, ..rest], [#(id_to_take, size_to_take), ..rest_to_take], _ -> {
      case first {
        File(id, size) -> {
          let available_size = int.min(size, uncompressed_size)
          compact_files_rec(
            add_file(compressed, id, available_size),
            rest,
            to_take,
            uncompressed_size - available_size,
          )
        }
        Empty(size_to_fill) -> {
          let available_size_to_take = int.min(size_to_take, uncompressed_size)

          case int.compare(size_to_fill, available_size_to_take) {
            order.Gt ->
              compact_files_rec(
                add_file(compressed, id_to_take, available_size_to_take),
                add_empty(rest, size_to_fill - available_size_to_take),
                rest_to_take,
                uncompressed_size - available_size_to_take,
              )
            order.Lt ->
              compact_files_rec(
                add_file(compressed, id_to_take, size_to_fill),
                rest,
                [
                  #(id_to_take, available_size_to_take - size_to_fill),
                  ..rest_to_take
                ],
                uncompressed_size - size_to_fill,
              )
            order.Eq ->
              compact_files_rec(
                add_file(compressed, id_to_take, available_size_to_take),
                rest,
                rest_to_take,
                uncompressed_size - available_size_to_take,
              )
          }
        }
      }
    }
  }
}

pub fn compact_files(disk_map: List(Block)) -> List(Block) {
  let uncompressed_size =
    disk_map
    |> list.fold(0, fn(acc, block) {
      case block {
        File(_, size) -> acc + size
        _ -> acc
      }
    })

  let to_take =
    disk_map
    |> list.filter_map(fn(block) {
      case block {
        File(id, size) -> Ok(#(id, size))
        Empty(_) -> Error(Nil)
      }
    })
    |> list.reverse

  compact_files_rec([], disk_map, to_take, uncompressed_size)
  |> list.reverse
}

fn reverse_append_removing(
  target: List(Block),
  source: List(Block),
  id_to_remove: Int,
) -> List(Block) {
  case source {
    [] -> target
    [first, ..rest] ->
      case first {
        File(id, size) if id == id_to_remove ->
          reverse_append_removing(add_empty(target, size), rest, id_to_remove)
        File(id, size) ->
          reverse_append_removing(
            add_file(target, id, size),
            rest,
            id_to_remove,
          )
        Empty(size) ->
          reverse_append_removing(add_empty(target, size), rest, id_to_remove)
      }
  }
}

fn move_one_file_rec(
  target: List(Block),
  source: List(Block),
  id: Int,
  size: Int,
) -> List(Block) {
  case source {
    [] -> target
    [first, ..rest] ->
      case first {
        Empty(empty_size) if empty_size > size ->
          target
          |> add_file(id, size)
          |> add_empty(empty_size - size)
          |> reverse_append_removing(rest, id)
        Empty(empty_size) if empty_size == size ->
          target
          |> add_file(id, size)
          |> reverse_append_removing(rest, id)
        File(file_id, file_size) if id == file_id ->
          target
          |> add_file(file_id, file_size)
          |> reverse_append_removing(rest, id)
        _ -> move_one_file_rec([first, ..target], rest, id, size)
      }
  }
}

fn move_one_file(disk_map: List(Block), id: Int, size: Int) -> List(Block) {
  move_one_file_rec([], disk_map, id, size) |> list.reverse
}

pub fn pretty_print(disk_map: List(Block)) -> String {
  disk_map
  |> list.map(fn(block) {
    case block {
      File(id, size) -> int.to_string(id) |> string.repeat(size)
      Empty(size) -> string.repeat(".", size)
    }
  })
  |> string.concat
}

pub fn compact_files_without_splitting(disk_map: List(Block)) -> List(Block) {
  disk_map
  |> list.filter_map(fn(block) {
    case block {
      File(id, size) -> Ok(#(id, size))
      Empty(_) -> Error(Nil)
    }
  })
  |> list.reverse
  |> list.fold(disk_map, fn(acc, to_take) {
    let #(id, size) = to_take
    move_one_file(acc, id, size)
  })
}

/// Need to compute:
/// index * id + (index + 1) * id + ... + (index + size - 1) * id
fn chunk_checksum(index: Int, id: Int, size: Int) {
  id * { { index * size } + { { size * { size - 1 } } / 2 } }
}

pub fn compute_checksum(blocks: List(Block)) -> Int {
  blocks
  |> list.fold(#(0, 0), fn(acc, block) {
    let #(checksum, index) = acc
    case block {
      File(id, size) -> {
        #(checksum + chunk_checksum(index, id, size), index + size)
      }
      Empty(size) -> #(checksum, index + size)
    }
  })
  |> pair.first
}

pub fn run(input: String) -> Result(utils.AdventOfCodeResult, String) {
  use disk_map <- result.map(parse_input(input))

  let first_part = disk_map |> compact_files |> compute_checksum
  let second_part =
    disk_map |> compact_files_without_splitting |> compute_checksum

  utils.AdventOfCodeResult(
    int.to_string(first_part),
    int.to_string(second_part),
  )
}
