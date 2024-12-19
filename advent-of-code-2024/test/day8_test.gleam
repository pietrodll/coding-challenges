import day8
import gleam/dict
import gleeunit/should

const input = "............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............"

pub fn parse_input_test() {
  let #(size, positions_by_char) = day8.parse_input(input)

  size |> should.equal(#(12, 12))

  dict.size(positions_by_char) |> should.equal(2)

  dict.get(positions_by_char, "0")
  |> should.be_ok
  |> should.equal([#(4, 4), #(3, 7), #(2, 5), #(1, 8)])

  dict.get(positions_by_char, "A")
  |> should.be_ok
  |> should.equal([#(9, 9), #(8, 8), #(5, 6)])
}

pub fn count_antinodes_test() {
  let #(size, positions_by_char) = day8.parse_input(input)

  day8.count_antinodes(size, positions_by_char, day8.compute_antinodes_for_two)
  |> should.equal(14)
}

pub fn count_antinodes_with_resonance_test() {
  let #(size, positions_by_char) = day8.parse_input(input)

  day8.count_antinodes(
    size,
    positions_by_char,
    day8.compute_antinodes_with_resonance,
  )
  |> should.equal(34)
}
