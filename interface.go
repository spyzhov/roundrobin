package roundrobin

type RoundRobin[T any] interface {
	Next() T
}
