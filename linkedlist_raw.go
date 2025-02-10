package roundrobin

import (
	"github.com/spyzhov/roundrobin/internal"
)

type LinkedListRaw[T any] struct {
	next *internal.LinkedList[T]
}

func NewLinkedListRaw[T any](array []T) *LinkedListRaw[T] {
	return &LinkedListRaw[T]{
		next: internal.NewCircledLinkedList(array),
	}
}

func (c *LinkedListRaw[T]) Next() (value T) {
	value = c.next.Value
	c.next = c.next.Next
	return value
}
