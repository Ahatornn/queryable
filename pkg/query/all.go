package query

// All Checks all elements to see if the condition is met
func (q Queryable[T]) All(predicate func(T) bool) bool {
	if q == nil {
		return true
	}
	var any bool
	var all = true
	q(func(item T) bool {
		any = true
		if !predicate(item) {
			all = false
			return false
		}
		return true
	})
	return !any || all
}
