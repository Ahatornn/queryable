package query

// Intersect returns elements that are both in the current sequence,
// and `second`. Duplicates are removed. The order is the same as in the first sequence.
//
// Empty (nil) sequences are treated as empty sets.
// ⚠️ Warning: this operation completely buffers `second` into memory (creates a map),
// so it is not lazy with respect to `second`.
func (q Queryable[T]) Intersect(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		secondSet := make(map[T]bool)
		if second != nil {
			second(func(item T) bool {
				secondSet[item] = true
				return true
			})
		}

		if q == nil {
			return
		}

		seen := make(map[T]bool)
		q(func(item T) bool {
			if secondSet[item] && !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
