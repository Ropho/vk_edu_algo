package sem3

type nodeQueue[T any] struct {
	val  T
	next *nodeQueue[T]
	prev *nodeQueue[T]
}

type Queue[T any] struct {
	start *nodeQueue[T]
	end   *nodeQueue[T]
	size  int
}

func (q *Queue[T]) Push(val T) {
	node := &nodeQueue[T]{
		val: val,
	}

	if q.start != nil {
		q.start.prev = node
		node.next = q.start
	} else {
		q.end = node
	}

	q.start = node

	q.size++
}

func (q *Queue[T]) Pop() {

	if q.end != nil {
		q.end = q.end.prev
		q.size--
	}
}

func (q Queue[T]) PeekEnd() T {

	var val T

	if q.end != nil {
		return q.end.val
	}

	return val
}
func (q Queue[T]) PeekStart() T {

	var val T

	if q.start != nil {
		return q.start.val
	}

	return val
}

func (q Queue[T]) Size() int {
	return q.size
}

type nodeStack[T any] struct {
	next *nodeStack[T]
	val  T
}
type Stack[T any] struct {
	end  *nodeStack[T]
	size int
}

func (q *Stack[T]) Push(val T) {
	node := &nodeStack[T]{
		val: val,
	}

	if q.end != nil {
		node.next = q.end
	}

	q.size++
	q.end = node
}

func (q *Stack[T]) Pop() {

	if q.end != nil {
		q.end = q.end.next
		q.size--
	}
}

func (q Stack[T]) Peek() T {

	var val T

	if q.end != nil {
		return q.end.val
	}

	return val
}

func (q Stack[T]) Size() int {
	return q.size
}
