package days

import scala.io.Source
import utils.Counter

case class Card(id: Int, winningNumbers: Set[Int], numbers: List[Int])

class Day4(resourcePath: String = "day4.txt") extends Day {
  override def number: Int = 4

  def parseCard(str: String): Card =
    val colonIndex = str.indexOf(':')
    val pipeIndex = str.indexOf('|')

    Card(
      str.substring(5, colonIndex).trim().toInt,
      str
        .substring(colonIndex + 1, pipeIndex)
        .trim()
        .split("\\s+")
        .map(_.toInt)
        .toSet,
      str.substring(pipeIndex + 1).trim().split("\\s+").map(_.toInt).toList
    )

  def countMatchingNumbers(card: Card): Int =
    card.numbers.count(card.winningNumbers.contains)

  def computePoints(card: Card): Int =
    math.pow(2, countMatchingNumbers(card) - 1).toInt

  def countCopies(cards: IterableOnce[Card]): Counter[Int] =
    cards
      .foldLeft(Counter[Int]())((counter, card) =>
        (1 to countMatchingNumbers(card)).foldLeft(
          counter.incremented(card.id)
        )((co, shift) => co.incremented(card.id + shift, co(card.id)))
      )

  override def firstPart(): String =
    Source
      .fromResource(resourcePath)
      .getLines
      .map(parseCard)
      .map(computePoints)
      .sum
      .toString

  override def secondPart(): String =
    countCopies(
      Source.fromResource(resourcePath).getLines.map(parseCard)
    ).sum.toString
}
