package query

// First первый элемент или nil
func (q Queryable[T]) First() *T {
	if q == nil {
		return nil
	}
	var result *T
	q(func(item T) bool {
		temp := item // копируем значение
		result = &temp
		return false // останавливаем после первого
	})
	return result
}
