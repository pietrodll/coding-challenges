package days

class Day4Test extends munit.FunSuite:
  test("should run first part on example") {
    val testDay = Day4("day4_test.txt")
    assertEquals(testDay.firstPart(), "13")
  }

  test("should run second part on example") {
    val testDay = Day4("day4_test.txt")
    assertEquals(testDay.secondPart(), "30")
  }
