package query

// FirstOrDefault первый или значение по умолчанию
func (q Queryable[T]) FirstOrDefault(defaultValue T) T {
	if first := q.First(); first != nil {
		return *first
	}
	return defaultValue
}
