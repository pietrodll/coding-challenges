pub fn gcd(a: Int, b: Int) -> Int {
  case b {
    0 -> a
    _ -> gcd(b, a % b)
  }
}

fn bezout_identity_rec(
  a: Int,
  b: Int,
  acc: #(Int, Int, Int, Int),
) -> #(Int, Int, Int, Int) {
  case b {
    0 -> acc
    _ -> {
      let q = a / b
      let r = a % b

      bezout_identity_rec(b, r, #(
        acc.2,
        acc.3,
        acc.0 - { q * acc.2 },
        acc.1 - { q * acc.3 },
      ))
    }
  }
}

pub fn bezout_identity(a: Int, b: Int) -> #(Int, Int) {
  let #(u, v, _, _) = bezout_identity_rec(a, b, #(1, 0, 0, 1))

  #(u, v)
}
