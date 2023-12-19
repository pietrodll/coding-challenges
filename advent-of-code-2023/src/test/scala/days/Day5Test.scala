package days

class Day5Test extends munit.FunSuite:
  def checkMapping(categoryMaps: List[CategoryMap], numbers: List[Long]): Unit =
    categoryMaps.zipWithIndex.foldLeft(numbers(0))((num, tup) => {
      val (map, index) = tup
      val result = map.map(num)
      assertEquals(result, numbers(index + 1), clue((num, map)))
      result
    })

  test("should map numbers correctly") {
    val testDay = Day5("day5_test.txt")
    val initialValues = testDay.almanach.initialValues
    val categoryMaps = testDay.almanach.categoryMaps

    assertEquals(categoryMaps.size, 7)

    assertEquals(initialValues(0), 79L)
    checkMapping(categoryMaps, List(79, 81, 81, 81, 74, 78, 78, 82))

    assertEquals(initialValues(1), 14L)
    checkMapping(categoryMaps, List(14, 14, 53, 49, 42, 42, 43, 43))

    assertEquals(initialValues(2), 55L)
    checkMapping(categoryMaps, List(55, 57, 57, 53, 46, 82, 82, 86))

    assertEquals(initialValues(3), 13L)
    checkMapping(categoryMaps, List(13, 13, 52, 41, 34, 34, 35, 35))
  }

  test("should run first part on example") {
    val testDay = Day5("day5_test.txt")
    assertEquals(testDay.firstPart(), "35")
  }

  test("should map ranges correctly") {
    val catMap = CategoryMap(
      "seed",
      "soil",
      List(RangeDescriptor(50, 98, 2), RangeDescriptor(52, 50, 48))
    )

    assertEquals(
      RangeDescriptor(52, 50, 48).mapRange(79, 79 + 14),
      (List((81L, 95L)), List())
    )
    assertEquals(catMap.mapRange(79, 79 + 14), List((81L, 95L)))
  }

  test("should run second part on example") {
    val testDay = Day5("day5_test.txt")
    assertEquals(testDay.secondPart(), "46")
  }
