package query

// Count counting elements
func (q Queryable[T]) Count() int {
	if q == nil {
		return 0
	}
	var cnt int
	q(func(T) bool {
		cnt++
		return true
	})
	return cnt
}
