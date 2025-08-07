package query

// Union объединяет текущую последовательность с другой, удаляя дубликаты.
// Возвращает новую Queryable[T], содержащую уникальные элементы из обеих последовательностей.
// Порядок: сначала уникальные элементы из q, затем из second, которых не было в q.
func (q Queryable[T]) Union(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]bool)

		// Обрабатываем первую последовательность
		q(func(item T) bool {
			if !seen[item] {
				seen[item] = true
				// Если yield вернул false — останавливаемся
				if !yield(item) {
					return false
				}
			}
			return true
		})

		// Если внешний yield ещё ожидает данные, продолжаем со второй
		second(func(item T) bool {
			if !seen[item] {
				seen[item] = true
				return yield(item)
			}
			return true
		})
	}
}
