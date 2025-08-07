package query

// First is the first element or nil
func (q Queryable[T]) First() *T {
	if q == nil {
		return nil
	}
	var result *T
	q(func(item T) bool {
		temp := item
		result = &temp
		return false
	})
	return result
}
