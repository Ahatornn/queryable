package query

// Any Проверяет, есть ли хотя бы один элемент?
func (q Queryable[T]) Any() bool {
	if q == nil {
		return false
	}
	var has bool
	q(func(T) bool {
		has = true
		return false // останавливаемся на первом
	})
	return has
}
