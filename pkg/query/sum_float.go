package query

// SumFloat calculates the sum of the values extracted from the items by selector.
func (q Queryable[T]) SumFloat(selector func(T) float32) float32 {
	if q == nil {
		return 0
	}
	var sum float32
	q(func(item T) bool {
		sum += selector(item)
		return true
	})
	return sum
}
