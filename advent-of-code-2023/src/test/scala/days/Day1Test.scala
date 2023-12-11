package days

class Day1Test extends munit.FunSuite:
  test("should run first part on example") {
    val testDay = Day1("day1_test_1.txt")

    assertEquals(testDay.firstPart(), "142")
  }

  test("should find spelled digits correctly") {
    assertEquals(findFirstAndLastDigitsIncludingSpelled("two1nine"), (2, 9))
    assertEquals(findFirstAndLastDigitsIncludingSpelled("eightwothree"), (8, 3))
    assertEquals(
      findFirstAndLastDigitsIncludingSpelled("abcone2threexyz"),
      (1, 3)
    )
    assertEquals(findFirstAndLastDigitsIncludingSpelled("xtwone3four"), (2, 4))
    assertEquals(
      findFirstAndLastDigitsIncludingSpelled("4nineeightseven2"),
      (4, 2)
    )
    assertEquals(findFirstAndLastDigitsIncludingSpelled("zoneight234"), (1, 4))
    assertEquals(
      findFirstAndLastDigitsIncludingSpelled("7pqrstsixteen"),
      (7, 6)
    )
  }

  test("should run second part on example") {
    val testDay = Day1("day1_test_2.txt")

    assertEquals(testDay.secondPart(), "281")
  }
