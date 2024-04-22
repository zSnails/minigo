package stack

import (
	"testing"
)

func TestPushSizeZero(t *testing.T) {
	s := NewStack[int](0)
	err := s.Push(23)
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}
}

func TestPopSizeZero(t *testing.T) {
	s := NewStack[int](0)
	_, err := s.Pop()
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}
}

func TestPop(t *testing.T) {
	s := NewStack[int](10)
	s.Push(1)
	s.Push(12)
	s.Push(29)
	s.Push(90)

	val, err := s.Pop()
	if err != nil {
		t.Fatalf("popping here should not produce an error")
	}
	if val != 90 {
		t.Fatalf("expected val to equal 90 instead got %d", val)
	}
}

func TestPush(t *testing.T) {
	s := NewStack[int](5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	err := s.Push(5)
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}
}

func TestPop2(t *testing.T) {
	s := NewStack[int](5)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	err := s.Push(6)
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}

    s.Pop()
    s.Pop()
    s.Pop()
    s.Pop()
    s.Pop()
    _, err = s.Pop()
    if err == nil {
		t.Fatalf("expected err to not be nil")
    }
}
