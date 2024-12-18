import day5
import gleam/list
import gleeunit/should

const input = "47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47"

pub fn parse_input_test() {
  input |> day5.parse_input |> should.be_ok
}

pub fn is_correctly_ordered_test() {
  let #(ordering_rules, _) = parse_input_test()
  let order_fun = day5.generate_order(ordering_rules)

  day5.is_correctly_ordered(order_fun, [75, 47, 61, 53, 29])
  |> should.be_true

  day5.is_correctly_ordered(order_fun, [97, 61, 53, 29, 13])
  |> should.be_true

  day5.is_correctly_ordered(order_fun, [75, 29, 13])
  |> should.be_true

  day5.is_correctly_ordered(order_fun, [75, 97, 47, 61, 53])
  |> should.be_false

  day5.is_correctly_ordered(order_fun, [61, 13, 29])
  |> should.be_false

  day5.is_correctly_ordered(order_fun, [97, 13, 75, 29, 47])
  |> should.be_false
}

pub fn find_middle_test() {
  [
    #([75, 47, 61, 53, 29], 61),
    #([97, 61, 53, 29, 13], 53),
    #([75, 29, 13], 29),
  ]
  |> list.each(fn(list_and_expected) {
    day5.find_middle(list_and_expected.0)
    |> should.be_some
    |> should.equal(list_and_expected.1)
  })
}
