package double

import (
	"fmt"
	"strings"
)

type LinkedList[T any] struct {
	Len  int
	Head *Node[T]
	Tail *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
	node := NewNode(value)
	node.Next = l.Head
	if l.Head != nil {
		l.Head.Previous = node
	}
	if l.Tail == nil {
		l.Tail = node
	}
	l.Head = node
	l.Len++
}
func (t *LinkedList[T]) String() string {
	sb := strings.Builder{}
	fmt.Fprint(&sb, "[")
	current := t.Head
	for current != nil {
		fmt.Fprintf(&sb, "%v", current.Value)
		if current.Next != nil {
			fmt.Fprint(&sb, ", ")
		}
		current = current.Next
	}
	fmt.Fprint(&sb, "]")
	return sb.String()
}

func (l *LinkedList[T]) ForEach(a func(T)) {
	for curr := l.Head; curr != nil; curr = curr.Next {
		a(curr.Value)
	}
}

func (l *LinkedList[T]) ForEachReverse(a func(T)) {
	for curr := l.Tail; curr != nil; curr = curr.Previous {
		a(curr.Value)
	}
}

func (l *LinkedList[T]) RemoveIf(predicate func(T) bool) {
	current := l.Head

	for current != nil {
		nextNode := current.Next
		if predicate(current.Value) {
			if current.Previous != nil {
				current.Previous.Next = current.Next
			}

			if current.Next != nil {
				current.Next.Previous = current.Previous
			}

			if current == l.Head {
				l.Head = current.Next
			}

			if current == l.Tail {
				l.Tail = current.Previous
			}

			l.Len--
		}

		current = nextNode
	}
}

func (l *LinkedList[T]) FindFirst(predicate func(T) bool) (T, bool) {
	for curr := l.Head; curr != nil; curr = curr.Next {
		if predicate(curr.Value) {
			return curr.Value, true
		}
	}

	return *new(T), false
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value:    value,
		Next:     nil,
		Previous: nil,
	}
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		Len:  0,
		Head: nil,
		Tail: nil,
	}
}

type Node[T any] struct {
	Value    T
	Next     *Node[T]
	Previous *Node[T]
}
