package days

import scala.io.Source

case class RangeDescriptor(destination: Long, source: Long, length: Long) {
  def map(num: Long): Option[Long] =
    if num >= source && num < source + length then
      Option(num - source + destination)
    else Option.empty

    /** Takes a range as parameter (left bound included, right bound excluded)
      * and returns a pair of mapped and unmapped ranges. Splitting between
      * mapped and unmapped is necessary to know which ranges need to pass
      * through the following mappers and which must not.
      */
  def mapRange(from: Long, to: Long): (List[(Long, Long)], List[(Long, Long)]) =
    if from >= source + length || to <= source then (List(), List((from, to)))
    else if to <= source + length && from >= source then
      (List((from - source + destination, to - source + destination)), List())
    else if to <= source + length then
      (
        List((destination, to - source + destination)),
        List(
          (from, source)
        )
      )
    else if from >= source then
      (
        List(
          (from - source + destination, destination + length)
        ),
        List(
          (source + length, to)
        )
      )
    else
      (
        List(
          (destination, destination + length)
        ),
        List(
          (from, source),
          (source + length, to)
        )
      )
}

def parseRangeDescriptor(line: String): RangeDescriptor =
  val values = line.trim().split(' ').map(_.toLong)
  RangeDescriptor(values(0), values(1), values(2))

case class CategoryMap(
    sourceCategory: String,
    destinationCategory: String,
    ranges: List[RangeDescriptor]
) {
  def map(num: Long): Long =
    ranges.iterator
      .map(range => range.map(num))
      .collectFirst { case Some(value) => value }
      .getOrElse(num)

  def mapRange(from: Long, to: Long): List[(Long, Long)] =
    val (mapped, unmapped) = ranges
      .foldLeft((List[(Long, Long)](), List[(Long, Long)]((from, to))))(
        (prevMappedAndUnmapped, range) => {
          val (prevMapped, prevUnmapped) = prevMappedAndUnmapped
          var nextMapped = prevMapped
          var nextUnmapped = List[(Long, Long)]()

          for prevUnmappedRange <- prevUnmapped do
            val (mapped, unmapped) = range.mapRange(
              prevUnmappedRange._1,
              prevUnmappedRange._2
            )
            nextMapped :++= mapped
            nextUnmapped :++= unmapped

          (nextMapped, nextUnmapped)
        }
      )

    mapped :++ unmapped
}

val categoryMapRegex = """^(\w+)-to-(\w+)\smap:$""".r

def parseCategoryMap(lines: Iterator[String]): CategoryMap =
  val firstLine = lines.next()
  categoryMapRegex
    .findFirstMatchIn(firstLine)
    .map(m =>
      CategoryMap(
        m.group(1),
        m.group(2),
        lines.map(parseRangeDescriptor).toList
      )
    )
    .get

case class Almanach(initialValues: List[Long], categoryMaps: List[CategoryMap])

def parseAlmanach(lines: Iterator[String]): Almanach =
  val initialValues =
    lines.next().substring(7).split(' ').map(_.toLong).toList
  var categoryMaps: List[CategoryMap] = List()
  lines.next()

  while lines.hasNext do
    var categoryMapSection: List[String] = List()
    var nextLine = lines.next()

    while nextLine != "" do
      categoryMapSection :+= nextLine
      nextLine = lines.nextOption().getOrElse("")

    categoryMaps :+= parseCategoryMap(categoryMapSection.iterator)

  Almanach(initialValues, categoryMaps)

class Day5(val resourcePath: String = "day5.txt") extends Day:
  override def number: Int = 5

  val almanach: Almanach = parseAlmanach(
    Source.fromResource(resourcePath).getLines
  )

  override def firstPart(): String =
    almanach.initialValues
      .map(value =>
        almanach.categoryMaps.foldLeft(value)((v, cat) => cat.map(v))
      )
      .min
      .toString

  override def secondPart(): String =
    almanach.initialValues
      .grouped(2)
      .flatMap(fromAndLength =>
        almanach.categoryMaps.foldLeft(
          List((fromAndLength(0), fromAndLength(0) + fromAndLength(1)))
        )((ranges, cat) =>
          ranges.flatMap(range => cat.mapRange(range._1, range._2))
        )
      )
      .map(_._1)
      .min
      .toString
