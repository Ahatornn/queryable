package query

// Union combines the current sequence with another, removing duplicates.
// Returns a new Queryable[T] containing unique elements from both sequences.
// Order: first unique elements from q, then from second that were not in q.
//
// Empty (nil) sequences are treated as empty.
// ⚠️ Note: the second sequence is iterated only if necessary
// (if the external consumer is still waiting for data).
func (q Queryable[T]) Union(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]bool)
		stopped := false

		yieldGuard := func(item T) bool {
			if stopped {
				return false
			}
			if !yield(item) {
				stopped = true
				return false
			}
			return true
		}

		if q != nil {
			q(func(item T) bool {
				if !seen[item] {
					seen[item] = true
					if !yieldGuard(item) {
						return false
					}
				}
				return true
			})
		}

		if !stopped && second != nil {
			second(func(item T) bool {
				if !seen[item] {
					seen[item] = true
					return yieldGuard(item)
				}
				return true
			})
		}
	}
}
