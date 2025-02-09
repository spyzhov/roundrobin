package internal

type LinkedList[T any] struct {
	Value T
	Next  *LinkedList[T]
}

func NewCircledLinkedList[T any](array []T) *LinkedList[T] {
	if len(array) == 0 {
		return nil
	}
	first := &LinkedList[T]{
		Value: array[0],
		Next:  nil,
	}
	list := first
	for i := 1; i < len(array); i++ {
		list.Next = &LinkedList[T]{
			Value: array[i],
			Next:  nil,
		}
		list = list.Next
	}
	list.Next = first
	return first
}
