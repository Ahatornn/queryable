package query

// ForEach Выполняет действие для каждого элемента
func (q Queryable[T]) ForEach(action func(T)) {
	if q == nil {
		return
	}
	q(func(item T) bool {
		action(item)
		return true
	})
}
