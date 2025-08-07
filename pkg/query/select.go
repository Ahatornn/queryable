package query

// Select Определяет проекцию. Пока сложно использовать из за any
func (q Queryable[T]) Select(mapper func(T) any) Queryable[any] {
	if q == nil {
		return Empty[any]()
	}
	return func(yield func(any) bool) {
		q(func(item T) bool {
			return yield(mapper(item))
		})
	}
}
