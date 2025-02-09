package roundrobin

import (
	"github.com/spyzhov/roundrobin/internal"
)

type LinkedListNoSplit[T any] struct {
	next *internal.LinkedList[T]
}

func NewLinkedListNoSplit[T any](array []T) *LinkedListNoSplit[T] {
	return &LinkedListNoSplit[T]{
		next: internal.NewCircledLinkedList(array),
	}
}

//go:nosplit
func (c *LinkedListNoSplit[T]) Next() (value T) {
	value = c.next.Value
	c.next = c.next.Next
	return value
}
