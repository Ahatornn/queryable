package query

// Any Checks if there is at least one element?
func (q Queryable[T]) Any() bool {
	if q == nil {
		return false
	}
	var has bool
	q(func(T) bool {
		has = true
		return false
	})
	return has
}
