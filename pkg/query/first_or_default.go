package query

// FirstOrDefault first or default value
func (q Queryable[T]) FirstOrDefault(defaultValue T) T {
	if first := q.First(); first != nil {
		return *first
	}
	return defaultValue
}
