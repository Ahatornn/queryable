package query

// ToQueryable creates a wrapper for lazy operations from a slice
func ToQueryable[T comparable](items []T) Queryable[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}
