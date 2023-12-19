package days

import scala.io.Source

case class Race(time: Long, distance: Long)

def parseRaces(timeLine: String, distanceLine: String): List[Race] =
  val timeValues =
    timeLine.substring(6).trim().split("\\s+").map(_.toLong)
  val distanceValues =
    distanceLine.substring(9).trim().split("\\s+").map(_.toLong)

  timeValues
    .lazyZip(distanceValues)
    .map((time, distance) => Race(time, distance))
    .toList

def parseSingleRace(timeLine: String, distanceLine: String): Race =
  Race(
    timeLine.substring(6).filterNot(_.isSpaceChar).toLong,
    distanceLine.substring(9).trim().filterNot(_.isSpaceChar).toLong
  )

// D = T * (Tmax - T)
// f = T * (Tmax - T) - Dmin, 0 < T < Tmax
//   = - T^2 + Tmax * T - Dmin
// Delta = Tmax^2 - 4Dmin
// R1 = (Tmax - sqrt(Tmax^2 - 4Dmin)) / 2
// R2 = (Tmax + sqrt(Tmax^2 - 4Dmin)) / 2
// N = (ceil(R2) - 1) - (floor(R1) + 1) + 1 = ceil(R2) - floor(R1) - 1

class Day6(val resourcePath: String = "day6.txt") extends Day:
  override def number: Int = 6

  def countWaysToWin(tMax: Long, dMin: Long): Long =
    val delta = tMax * tMax - 4 * dMin
    val deltaRoot = math.sqrt(delta)
    val r1 = (tMax - deltaRoot) / 2
    val r2 = (tMax + deltaRoot) / 2

    math.ceil(r2).toLong - math.floor(r1).toLong - 1

  private def getRaces: List[Race] =
    val lines = Source.fromResource(resourcePath).getLines.toList
    parseRaces(lines(0), lines(1))

  override def firstPart(): String =
    getRaces
      .map(race => countWaysToWin(race.time, race.distance))
      .reduce(_ * _)
      .toString

  def getSingleRace: Race =
    val lines = Source.fromResource(resourcePath).getLines.toList
    parseSingleRace(lines(0), lines(1))

  override def secondPart(): String =
    val race = getSingleRace
    countWaysToWin(race.time, race.distance).toString
