package double

import "testing"

func TestAdd(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)

	if l.Head != l.Tail {
		t.Fatalf("expected head and tail to be equal")
	}

	if l.Tail.Value != 1 && l.Head.Value != 1 {
		t.Fatalf("expected head and tail to be equal to 1, insead got head=%d, tail=%d", l.Head.Value, l.Tail.Value)
	}
}

func TestRemoveIf(t *testing.T) {
	l := NewLinkedList[int]()

	l.Add(1)
	l.Add(2)
	l.Add(3)
	l.Add(4)

	if l.Len != 4 {
		t.Fatalf("expected len to be 4, instead got %d", l.Len)
	}

	l.RemoveIf(func(i int) bool {
		return i == 3
	})

	if l.Len != 3 {
		t.Fatalf("expected len to be 3, instead got %d", l.Len)
	}

	if l.Head == l.Tail {
		t.Fatalf("expected head and tail to be different, instead got the same value")
	}

	l.RemoveIf(func(i int) bool {
		return i == 2 || i == 4
	})

	if l.Head != l.Tail {
		t.Fatalf("expected head and tail to be the same value, instead got different values")
	}
}
