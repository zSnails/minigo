package single

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1)
	list.Add(2)
	assert.Equal(t, 2, list.Len, "Expected size not met")
}

func TestRemoveIf(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(2)
	list.Add(2)
	list.Add(1)
	assert.Equal(t, 3, list.Len, "Expected size not met")
	list.RemoveIf(func(i int) bool {
		return i == 2
	})
	assert.Equal(t, 1, list.Len, "Expected size not met")
}

func TestFindFirst(t *testing.T) {
	list := NewLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	assert.Equal(t, 4, list.Len, "Expected size not met")
	value, found := list.FindFirst(func(i int) bool {
		return i == 2
	})
	if !found {
		assert.FailNow(t, "Value 2 not found, unreachable case")
		return
	}
	assert.Equal(t, value, 2, "Mismatched values")
}
