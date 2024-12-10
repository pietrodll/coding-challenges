import day3
import gleeunit/should

pub fn extract_operations_test() {
  "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
  |> day3.extract_operations
  |> should.be_ok
  |> should.equal([
    day3.Mul(2, 4),
    day3.Mul(5, 5),
    day3.Mul(11, 8),
    day3.Mul(8, 5),
  ])

  "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
  |> day3.extract_operations
  |> should.be_ok
  |> should.equal([
    day3.Mul(2, 4),
    day3.DoNot,
    day3.Mul(5, 5),
    day3.Mul(11, 8),
    day3.Do,
    day3.Mul(8, 5),
  ])
}

pub fn compute_result_test() {
  [day3.Mul(2, 4), day3.Mul(5, 5), day3.Mul(11, 8), day3.Mul(8, 5)]
  |> day3.compute_result
  |> should.equal(161)
}

pub fn compute_result_including_dos_test() {
  [
    day3.Mul(2, 4),
    day3.DoNot,
    day3.Mul(5, 5),
    day3.Mul(11, 8),
    day3.Do,
    day3.Mul(8, 5),
  ]
  |> day3.compute_result_including_dos
  |> should.equal(48)
}
