import days.Day

class TestDay extends Day:
  override def number: Int = 0
  override def firstPart(): String = ""
  override def secondPart(): String = ""

class MainTest extends munit.FunSuite {
  test("should run without errors") {
    runDay(TestDay())
  }
}
