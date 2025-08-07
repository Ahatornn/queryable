package query

func (q Queryable[T]) ToSlice() []T {
	if q == nil {
		return nil
	}
	result := make([]T, 0, 16)
	q(func(item T) bool {
		result = append(result, item)
		return true
	})
	return result
}
