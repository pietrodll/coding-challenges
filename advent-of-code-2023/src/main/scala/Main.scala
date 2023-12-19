import days.{Day, Day1, Day2, Day3, Day4, Day5, Day6}

def allDays = List(Day1(), Day2(), Day3(), Day4(), Day5(), Day6())

def runDay(day: Day): Unit =
  println(s"Running day ${day.number}")
  println(s"First part: ${day.firstPart()}")
  println(s"Second part: ${day.secondPart()}")
  println("")

@main def run: Unit = allDays.foreach(runDay)
