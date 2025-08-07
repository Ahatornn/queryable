package query

// Distinct returns a sequence without duplicates.
// Preserves the order of the first occurrence of each element.
func (q Queryable[T]) Distinct() Queryable[T] {
	if q == nil {
		return Empty[T]()
	}
	return func(yield func(T) bool) {
		seen := make(map[T]bool)

		q(func(item T) bool {
			if !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false
				}
			}
			return true
		})
	}
}
