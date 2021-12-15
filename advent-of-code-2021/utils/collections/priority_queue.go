package collections

type PriorityElement interface {
	GetPriority() int
	Key() interface{}
}

// Implementation of a minimal priority queue
type PriorityQueue struct {
	items     []PriorityElement
	positions map[interface{}]int
}

func parent(i int) int {
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func (q *PriorityQueue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *PriorityQueue) Len() int {
	return len(q.items)
}

func (q *PriorityQueue) swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
	q.positions[q.items[i].Key()] = i
	q.positions[q.items[j].Key()] = j
}

func (q *PriorityQueue) percolateUp(start int) {
	i, par := start, parent(start)

	for par >= 0 && q.items[i].GetPriority() < q.items[par].GetPriority() {
		q.swap(i, par)
		i, par = par, parent(par)
	}
}

func (q *PriorityQueue) minChild(i int) int {
	if rightChild(i) >= len(q.items) {
		return leftChild(i)
	}

	if q.items[rightChild(i)].GetPriority() < q.items[leftChild(i)].GetPriority() {
		return rightChild(i)
	}

	return leftChild(i)
}

func (q *PriorityQueue) percolateDown(start int) {
	i, mc := start, q.minChild(start)

	for mc < len(q.items) && q.items[mc].GetPriority() < q.items[i].GetPriority() {
		q.swap(i, mc)
		i, mc = mc, q.minChild(mc)
	}
}

func (q *PriorityQueue) Add(x PriorityElement) {
	q.items = append(q.items, x)
	q.positions[x.Key()] = len(q.items) - 1
	q.percolateUp(len(q.items) - 1)
}

func (q *PriorityQueue) Pop() PriorityElement {
	lastIndex := len(q.items) - 1
	q.swap(0, lastIndex)
	result := q.items[lastIndex]
	q.items = q.items[:lastIndex]
	delete(q.positions, result.Key())
	q.percolateDown(0)
	return result
}

func (q *PriorityQueue) GetMin() PriorityElement {
	return q.items[0]
}

func (q *PriorityQueue) IncreasedPriority(x PriorityElement) {
	q.percolateDown(q.positions[x.Key()])
}

func (q *PriorityQueue) DecreasedPriority(x PriorityElement) {
	q.percolateUp(q.positions[x.Key()])
}

func NewPriorityQueue(elements ...PriorityElement) PriorityQueue {
	q := PriorityQueue{make([]PriorityElement, 0), make(map[interface{}]int)}

	for _, element := range elements {
		q.Add(element)
	}

	return q
}
