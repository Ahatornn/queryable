package query

// Except returns elements from the current sequence that are not in `second`.
// Duplicates are removed. The order is the same as in the first sequence.
func (q Queryable[T]) Except(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		excludeSet := make(map[T]bool)
		if second != nil {
			second(func(item T) bool {
				excludeSet[item] = true
				return true
			})
		}

		if q == nil {
			return
		}

		seen := make(map[T]bool)
		q(func(item T) bool {
			if !excludeSet[item] && !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
