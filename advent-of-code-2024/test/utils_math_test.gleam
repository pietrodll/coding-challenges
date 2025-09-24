import gleeunit/should
import utils_math

pub fn gcd_test() {
  utils_math.gcd(10, 5) |> should.equal(5)

  utils_math.gcd(10, 3) |> should.equal(1)

  utils_math.gcd(15, 12) |> should.equal(3)
}

pub fn bezout_identity_test() {
  utils_math.bezout_identity(10, 3) |> should.equal(#(1, -3))

  utils_math.bezout_identity(21, 12) |> should.equal(#(-1, 2))
}
