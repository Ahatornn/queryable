package query

// Concat concatenates the current sequence with another.
// Returns a new Queryable[T] with elements from q first, then from second.
func (q Queryable[T]) Concat(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		if q != nil {
			q(func(item T) bool {
				return yield(item)
			})
		}

		if second != nil {
			second(func(item T) bool {
				return yield(item)
			})
		}
	}
}
