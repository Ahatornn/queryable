package query

// Empty пустая последовательность
func Empty[T any]() Queryable[T] {
	return func(yield func(T) bool) {}
}
