package days

import scala.compiletime.ops.boolean
import scala.io.Source

case class CubeSubset(red: Int, green: Int, blue: Int)

def parseCubeSubset(str: String): CubeSubset =
  str
    .trim()
    .split(", ")
    .map(colorStr => colorStr.split(" "))
    .map(colorSplit => (colorSplit(0).toInt, colorSplit(1)))
    .foldLeft(CubeSubset(0, 0, 0))((cubeSet, valueAndColor) =>
      valueAndColor._2 match
        case "red"   => cubeSet.copy(red = valueAndColor._1)
        case "green" => cubeSet.copy(green = valueAndColor._1)
        case "blue"  => cubeSet.copy(blue = valueAndColor._1)
    )

case class Game(id: Int, records: List[CubeSubset])

def parseGame(line: String): Game =
  val gameAndRecords = line.trim().split(": ")
  assert(gameAndRecords.size == 2)

  val gameId = gameAndRecords(0).substring(5).toInt

  val cubeSubsets = gameAndRecords(1).split("; ").map(parseCubeSubset).toList

  Game(gameId, cubeSubsets)

def computeGameMinimumSet(game: Game): CubeSubset =
  game.records.reduce((a, b) =>
    CubeSubset(a.red.max(b.red), a.green.max(b.green), a.blue.max(b.blue))
  )

class Day2(
    resourcePath: String = "day2.txt",
    redCount: Int = 12,
    greenCount: Int = 13,
    blueCount: Int = 14
) extends Day:
  override def number: Int = 2

  def isPossible(game: Game): Boolean = game.records.forall(cubeSubset =>
    cubeSubset.red <= redCount && cubeSubset.green <= greenCount && cubeSubset.blue <= blueCount
  )

  override def firstPart(): String = Source
    .fromResource(resourcePath)
    .getLines()
    .map(parseGame)
    .filter(isPossible)
    .map(game => game.id)
    .sum
    .toString

  override def secondPart(): String = Source
    .fromResource(resourcePath)
    .getLines()
    .map(parseGame)
    .map(computeGameMinimumSet)
    .map(minimumSet => minimumSet.red * minimumSet.green * minimumSet.blue)
    .sum
    .toString
