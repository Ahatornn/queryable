package query

// SumInt вычисляет сумму значений, извлечённых из элементов с помощью selector.
// Пример: q.Sum(func(item T) int { return item.Age })
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
