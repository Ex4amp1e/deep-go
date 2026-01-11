package homework2

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type CircularQueue[T Number] struct {
	values []T
	size   int
	count  int
	front  int
	rear   int
}

func NewCircularQueue[T Number](size int) *CircularQueue[T] {
	if size <= 0 {
		panic("size must be > 0")
	}

	return &(CircularQueue[T]{
		values: make([]T, size),
		size:   size,
		rear:   -1,
	})
}

func (q *CircularQueue[T]) Push(value T) bool {
	if q.Full() {
		return false
	}

	// Increase rear
	q.rear = (q.rear + 1) % q.size

	// Set value
	q.values[q.rear] = value

	q.count++
	return true
}

func (q *CircularQueue[T]) Pop() bool {
	if q.Empty() {
		return false
	}

	// Take value - not implemented

	// Increase front
	q.front = (q.front + 1) % q.size

	q.count--
	return true
}

func (q *CircularQueue[T]) Front() T {
	if !q.Empty() {
		return q.values[q.front]
	}
	return T(-1)
}

func (q *CircularQueue[T]) Back() T {
	if !q.Empty() {
		return q.values[q.rear]
	}
	return T(-1)
}

func (q *CircularQueue[T]) Empty() bool {
	return q.count == 0
}

func (q *CircularQueue[T]) Full() bool {
	return q.count == q.size
}
