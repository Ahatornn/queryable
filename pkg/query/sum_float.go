package query

// SumInt вычисляет сумму значений, извлечённых из элементов с помощью selector.
// Пример: q.Sum(func(item T) int { return item.Amount })
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
