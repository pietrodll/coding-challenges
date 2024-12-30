import day9
import gleeunit/should

const input = "2333133121414131402"

pub fn compact_files_test() {
  input
  |> day9.parse_input
  |> should.be_ok
  |> day9.compact_files
  |> should.equal([
    day9.File(0, 2),
    day9.File(9, 2),
    day9.File(8, 1),
    day9.File(1, 3),
    day9.File(8, 3),
    day9.File(2, 1),
    day9.File(7, 3),
    day9.File(3, 3),
    day9.File(6, 1),
    day9.File(4, 2),
    day9.File(6, 1),
    day9.File(5, 4),
    day9.File(6, 2),
  ])
}

pub fn compact_files_without_splitting_test() {
  input
  |> day9.parse_input
  |> should.be_ok
  |> day9.compact_files_without_splitting
  |> should.equal([
    day9.File(0, 2),
    day9.File(9, 2),
    day9.File(2, 1),
    day9.File(1, 3),
    day9.File(7, 3),
    day9.Empty(1),
    day9.File(4, 2),
    day9.Empty(1),
    day9.File(3, 3),
    day9.Empty(4),
    day9.File(5, 4),
    day9.Empty(1),
    day9.File(6, 4),
    day9.Empty(5),
    day9.File(8, 4),
    day9.Empty(2),
  ])
}

pub fn compute_checksum_test() {
  let parsed_input =
    input
    |> day9.parse_input
    |> should.be_ok

  parsed_input
  |> day9.compact_files
  |> day9.compute_checksum
  |> should.equal(1928)

  parsed_input
  |> day9.compact_files_without_splitting
  |> day9.compute_checksum
  |> should.equal(2858)
}
