package query

// ForEach Performs an action on each element
func (q Queryable[T]) ForEach(action func(T)) {
	if q == nil {
		return
	}
	q(func(item T) bool {
		action(item)
		return true
	})
}
