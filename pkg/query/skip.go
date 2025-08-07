package query

// Skip skips the first n elements
func (q Queryable[T]) Skip(n int) Queryable[T] {
	if q == nil || n <= 0 {
		return q
	}
	return func(yield func(T) bool) {
		var skipped int
		q(func(item T) bool {
			if skipped < n {
				skipped++
				return true
			}
			return yield(item)
		})
	}
}
