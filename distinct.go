package query

// Distinct возвращает последовательность без дубликатов.
// Сохраняет порядок первого вхождения каждого элемента.
func (q Queryable[T]) Distinct() Queryable[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]bool)

		q(func(item T) bool {
			if !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false // остановка от потребителя
				}
			}
			return true
		})
	}
}
