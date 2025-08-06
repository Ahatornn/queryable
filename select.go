package query

// Select Определяет проекцию
func (q Queryable[T]) Select(mapper func(T) any) Queryable[any] {
	if q == nil {
		return nil
	}
	return func(yield func(any) bool) {
		q(func(item T) bool {
			return yield(mapper(item))
		})
	}
}
