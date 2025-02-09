package roundrobin

type Chan[T any] struct {
	next chan T
}

func NewChan[T any](array []T) *Chan[T] {
	c := &Chan[T]{
		next: make(chan T, len(array)),
	}
	c.init(array)
	return c
}

func (c *Chan[T]) Next() (value T) {
	value = c.pop()
	defer c.put(value)
	return value
}

func (c *Chan[T]) Close() error {
	close(c.next)
	return nil
}

func (c *Chan[T]) pop() T {
	return <-c.next
}

func (c *Chan[T]) put(value T) {
	c.next <- value
}

func (c *Chan[T]) init(array []T) {
	c.next = make(chan T, len(array))
	for i := 0; i < len(array); i++ {
		c.put(array[i])
	}
}
