package query

// Queryable[T] — ленивая последовательность элементов типа T.
// Поддерживает цепочки: .Where().Select().Take().First() и т.д.
type Queryable[T any] func(yield func(T) bool)
