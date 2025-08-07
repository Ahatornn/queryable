package query

// SumInt calculates the sum of the values extracted from the items by selector.
func (q Queryable[T]) SumInt(selector func(T) int) int {
	if q == nil {
		return 0
	}
	var sum int
	q(func(item T) bool {
		sum += selector(item)
		return true
	})
	return sum
}
