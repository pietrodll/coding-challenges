package days

import scala.io.Source

val spelledDigits = Map(
  "one" -> 1,
  "two" -> 2,
  "three" -> 3,
  "four" -> 4,
  "five" -> 5,
  "six" -> 6,
  "seven" -> 7,
  "eight" -> 8,
  "nine" -> 9
)

def findSpelledDigitAtIndex(str: String, idx: Int): Option[Int] =
  spelledDigits.toSeq
    .find((spelled, value) =>
      spelled.size <= str.length - idx && str
        .slice(idx, idx + spelled.size)
        .equals(spelled)
    )
    .map((spelled, value) => value)

def findFirstAndLastDigit(line: String): (Int, Int) =
  var first = -1
  var last = -1

  for chr <- line do
    if chr.isDigit then
      if first == -1 then first = chr.asDigit
      last = chr.asDigit

  (first, last)

def findFirstAndLastDigitsIncludingSpelled(line: String): (Int, Int) =
  var first = -1
  var last = -1

  for (chr, index) <- line.zipWithIndex do
    var v = -1
    if chr.isDigit then v = chr.asDigit
    else
      var spelledDigit = findSpelledDigitAtIndex(line, index)
      if spelledDigit.isDefined then v = spelledDigit.get

    if v != -1 then
      if first == -1 then first = v
      last = v

  (first, last)

class Day1(resourcePath: String = "day1.txt") extends Day:
  override def number: Int = 1

  private def computeTotal(lineParser: (String) => (Int, Int)): Int =
    Source
      .fromResource(resourcePath)
      .getLines()
      .map(lineParser)
      .map((first, last) => 10 * first + last)
      .sum

  override def firstPart(): String = computeTotal(
    findFirstAndLastDigit
  ).toString

  override def secondPart(): String = computeTotal(
    findFirstAndLastDigitsIncludingSpelled
  ).toString
