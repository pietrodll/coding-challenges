import day1
import gleeunit/should

const input = "3   4
4   3
2   5
1   3
3   9
3   3"

pub fn parse_input_test() {
  day1.parse_input(input)
  |> should.be_ok
  |> should.equal(#([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3]))
}

pub fn compute_distance_test() {
  day1.compute_distance([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3])
  |> should.equal(11)
}

pub fn compute_similarity_test() {
  day1.compute_similarity([3, 4, 2, 1, 3, 3], [4, 3, 5, 3, 9, 3])
  |> should.equal(31)
}
