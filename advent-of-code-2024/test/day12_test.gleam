import day12
import gleam/int
import gleam/list
import gleam/set
import gleeunit/should
import utils_matrix

const input1 = "AAAA
BBCD
BBCC
EEEC"

const input2 = "OOOOO
OXOXO
OOOOO
OXOXO
OOOOO"

const input3 = "RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE"

fn parse(input: String) {
  input |> day12.parse_input |> should.be_ok
}

pub fn find_regions_test() {
  input1
  |> parse
  |> day12.find_regions
  |> list.length
  |> should.equal(5)

  input2
  |> parse
  |> day12.find_regions
  |> list.length
  |> should.equal(5)

  input3
  |> parse
  |> day12.find_regions
  |> list.length
  |> should.equal(11)
}

pub fn compute_region_price_test() {
  let matrix1 = parse(input1)

  [
    #(#(0, 0), 40),
    #(#(1, 0), 32),
    #(#(1, 2), 40),
    #(#(1, 3), 4),
    #(#(3, 0), 24),
  ]
  |> list.each(fn(test_case) {
    let #(region_position, expected_price) = test_case

    day12.compute_region_price(matrix1, region_position)
    |> should.equal(expected_price)
  })
}

fn compute_total_price(m: utils_matrix.Matrix(String)) -> Int {
  m
  |> day12.find_regions
  |> list.map(day12.compute_region_price(m, _))
  |> int.sum
}

pub fn compute_total_price_test() {
  input1
  |> parse
  |> compute_total_price
  |> should.equal(140)

  input2
  |> parse
  |> compute_total_price
  |> should.equal(772)

  input3
  |> parse
  |> compute_total_price
  |> should.equal(1930)
}

pub fn get_region_fences_test() {
  let parsed_input1 = parse(input1)

  parsed_input1
  |> day12.get_region_fences(#(0, 0))
  |> should.equal(
    set.from_list([
      day12.Fence(#(0, 0), utils_matrix.Up),
      day12.Fence(#(0, 0), utils_matrix.Left),
      day12.Fence(#(0, 0), utils_matrix.Down),
      day12.Fence(#(0, 1), utils_matrix.Up),
      day12.Fence(#(0, 2), utils_matrix.Up),
      day12.Fence(#(0, 3), utils_matrix.Up),
      day12.Fence(#(0, 1), utils_matrix.Down),
      day12.Fence(#(0, 2), utils_matrix.Down),
      day12.Fence(#(0, 3), utils_matrix.Down),
      day12.Fence(#(0, 3), utils_matrix.Right),
    ]),
  )
}

pub fn compute_number_of_sides_test() {
  let parsed_input1 = input1 |> parse

  parsed_input1
  |> day12.compute_number_of_sides(#(0, 0))
  |> should.equal(4)
}

pub fn compute_region_price_with_discount_test() {
  let parsed_input = input3 |> parse

  parsed_input
  |> day12.compute_region_price_with_discount(#(0, 0))
  |> should.equal(120)
}

fn compute_total_price_with_discount(m: utils_matrix.Matrix(String)) -> Int {
  m
  |> day12.find_regions
  |> list.map(day12.compute_region_price_with_discount(m, _))
  |> int.sum
}

pub fn compute_total_price_with_discount_test() {
  input1
  |> parse
  |> compute_total_price_with_discount
  |> should.equal(80)

  input2
  |> parse
  |> compute_total_price_with_discount
  |> should.equal(436)

  input3
  |> parse
  |> compute_total_price_with_discount
  |> should.equal(1206)
}
