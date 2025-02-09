package roundrobin

import (
	"sync"

	"github.com/spyzhov/roundrobin/internal"
)

type LinkedListMutex[T any] struct {
	next *internal.LinkedList[T]
	mu   sync.Mutex
}

func NewLinkedListMutex[T any](array []T) *LinkedListMutex[T] {
	return &LinkedListMutex[T]{
		next: internal.NewCircledLinkedList(array),
	}
}

func (c *LinkedListMutex[T]) Next() (value T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value = c.next.Value
	c.next = c.next.Next
	return value
}
