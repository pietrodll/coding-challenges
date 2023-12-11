package days

class Day2Test extends munit.FunSuite:
  test("should parse game correctly") {
    assertEquals(
      parseGame("Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"),
      Game(
        1,
        List(CubeSubset(4, 0, 3), CubeSubset(1, 2, 6), CubeSubset(0, 2, 0))
      )
    )
  }

  test("should run first part on example") {
    val testDay = Day2(resourcePath = "day2_test.txt")
    assertEquals(testDay.firstPart(), "8")
  }

  test("should run second part on example") {
    val testDay = Day2(resourcePath = "day2_test.txt")
    assertEquals(testDay.secondPart(), "2286")
  }
