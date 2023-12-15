package utils

class Counter[T](private val count: Map[T, Int] = Map[T, Int]()):
  def incremented(key: T, value: Int): Counter[T] = Counter(
    count.updatedWith(key)(prev => Option(prev.getOrElse(0) + value))
  )

  def incremented(key: T): Counter[T] = incremented(key, 1)

  def apply(key: T): Int = count(key)

  def toMap: Map[T, Int] = count

  def sum: Int = count.valuesIterator.sum
