package days

import scala.io.Source

def computeAdjacentPositions(
    i: Int,
    j: Int,
    height: Int,
    width: Int
): List[(Int, Int)] = List(
  (i - 1, j - 1),
  (i - 1, j),
  (i - 1, j + 1),
  (i, j - 1),
  (i, j + 1),
  (i + 1, j - 1),
  (i + 1, j),
  (i + 1, j + 1)
).filter((i, j) => i >= 0 && j >= 0 && i < height && j < width)

class Day3(resourcePath: String = "day3.txt") extends Day:
  override def number: Int = 3

  private def isSymbol(s: Char): Boolean = !s.isDigit && s != '.'

  private def findPartNumbers(grid: List[String]): List[Int] =
    val height = grid.size
    val width = grid(0).size
    var partNumbers: List[Int] = List()

    for (line, i) <- grid.zipWithIndex do
      var current = ""
      var isPartNumber = false

      for (chr, j) <- line.zipWithIndex do
        if chr.isDigit then
          current += chr
          isPartNumber ||= computeAdjacentPositions(i, j, height, width).exists(
            (k, l) => isSymbol(grid(k).charAt(l))
          )
        else if current != "" then
          if isPartNumber then partNumbers = partNumbers :+ current.toInt
          current = ""
          isPartNumber = false

      if current != "" && isPartNumber then
        partNumbers = partNumbers :+ current.toInt

    partNumbers

  private def findGearRatios(grid: List[String]): List[Int] =
    val height = grid.size
    val width = grid(0).size
    var starNeighbors: Map[(Int, Int), List[Int]] = Map()

    for (line, i) <- grid.zipWithIndex do
      var current = ""
      var starPositions: Set[(Int, Int)] = Set()

      for (chr, j) <- line.zipWithIndex do
        if chr.isDigit then
          current += chr

          for (k, l) <- computeAdjacentPositions(i, j, height, width) do
            if grid(k).charAt(l) == '*' then
              starPositions += (k, l)
        else if current != "" then
          for pos <- starPositions do
            starNeighbors = starNeighbors.updatedWith(pos)(maybeList =>
              Option(maybeList.getOrElse(List()) :+ current.toInt)
            )
          current = ""
          starPositions = Set()

      if current != "" then
        for pos <- starPositions do
          starNeighbors = starNeighbors.updatedWith(pos)(maybeList =>
            Option(maybeList.getOrElse(List()) :+ current.toInt)
          )

    starNeighbors.valuesIterator
      .filter(_.length == 2)
      .map(nums => nums(0) * nums(1))
      .toList

  override def firstPart(): String =
    findPartNumbers(
      Source.fromResource(resourcePath).getLines.toList
    ).sum.toString

  override def secondPart(): String = findGearRatios(
    Source.fromResource(resourcePath).getLines.toList
  ).sum.toString
