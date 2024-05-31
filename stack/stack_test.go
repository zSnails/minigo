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
	_ = s.Push(1)
	_ = s.Push(12)
	_ = s.Push(29)
	_ = s.Push(90)

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
	_ = s.Push(1)
	_ = s.Push(2)
	_ = s.Push(3)
	_ = s.Push(4)
	_ = s.Push(5)
	err := s.Push(5)
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}
}

func TestPop2(t *testing.T) {
	s := NewStack[int](5)
	_ = s.Push(1)
	_ = s.Push(2)
	_ = s.Push(3)
	_ = s.Push(4)
	_ = s.Push(5)
	err := s.Push(6)
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}

	_, _ = s.Pop()
	_, _ = s.Pop()
	_, _ = s.Pop()
	_, _ = s.Pop()
	_, _ = s.Pop()
	_, err = s.Pop()
	if err == nil {
		t.Fatalf("expected err to not be nil")
	}
}
