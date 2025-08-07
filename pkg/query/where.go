package query

// Where filtering
func (q Queryable[T]) Where(predicate func(T) bool) Queryable[T] {
	if q == nil {
		return nil
	}
	return func(yield func(T) bool) {
		q(func(item T) bool {
			if predicate(item) {
				return yield(item)
			}
			return true
		})
	}
}
