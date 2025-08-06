package query

// Concat объединяет текущую последовательность с другой.
// Возвращает новую Queryable[T], в которой сначала идут элементы из q, потом из second.
func (q Queryable[T]) Concat(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		// Сначала итерируем текущую последовательность
		if q != nil {
			q(func(item T) bool {
				return yield(item)
			})
		}

		// Затем — вторую
		if second != nil {
			second(func(item T) bool {
				return yield(item)
			})
		}
	}
}
