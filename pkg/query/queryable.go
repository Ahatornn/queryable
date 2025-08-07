package query

// Queryable[T] is a lazy sequence of elements of type T.
// Supports chaining: .Where().Skip().First() etc.
type Queryable[T comparable] func(yield func(T) bool)
