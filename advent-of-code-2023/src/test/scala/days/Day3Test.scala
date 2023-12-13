package days

class Day3Test extends munit.FunSuite:
  test("should run first part on example") {
    val testDay = Day3("day3_test.txt")
    assertEquals(testDay.firstPart(), "4361")
  }

  test("should run second part on example") {
    val testDay = Day3("day3_test.txt")
    assertEquals(testDay.secondPart(), "467835")
  }
