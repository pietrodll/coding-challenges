package days

class Day6Test extends munit.FunSuite:
  test("should run first part on example") {
    val testDay = Day6("day6_test.txt")
    assertEquals(testDay.firstPart(), "288")
  }

  test("should run second part on example") {
    val testDay = Day6("day6_test.txt")
    assertEquals(testDay.getSingleRace, Race(71530, 940200))
    assertEquals(testDay.secondPart(), "71503")
  }
