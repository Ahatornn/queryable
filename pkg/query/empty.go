package query

// Empty is an empty sequence
func Empty[T comparable]() Queryable[T] {
	return func(yield func(T) bool) {}
}
