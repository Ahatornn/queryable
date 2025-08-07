package query

// Empty пустая последовательность
func Empty[T comparable]() Queryable[T] {
	return func(yield func(T) bool) {}
}
