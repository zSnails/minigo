package single

type LinkedList[T any] struct {
	Len  int
	Head *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
	node := NewNode(value)
	node.Next = l.Head
	l.Head = node
	l.Len++
}

func (l *LinkedList[T]) RemoveIf(predicate func(T) bool) {
	current := l.Head
	var prev *Node[T] = nil

	for current != nil {
		if predicate(current.Value) {
			if prev != nil {
				prev.Next = current.Next
			} else {
				l.Head = current.Next
			}
			current = current.Next
			l.Len--
		} else {
			prev = current
			current = current.Next
		}
	}
}

func (l *LinkedList[T]) FindFirst(predicate func(T) bool) (value T, found bool) {
	current := l.Head
	for current != nil {
		if predicate(current.Value) {
			return current.Value, true
		}
		current = current.Next
	}
	return *new(T), false
}

type Node[T any] struct {
	Next  *Node[T]
	Value T
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		Next:  nil,
	}
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}
