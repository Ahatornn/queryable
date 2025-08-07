package query

// Except возвращает элементы из текущей последовательности, которых нет в `second`.
// Дубликаты удаляются. Порядок — как в первой последовательности.
func (q Queryable[T]) Except(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		// Собираем все элементы из second
		excludeSet := make(map[T]bool)
		second(func(item T) bool {
			excludeSet[item] = true
			return true
		})

		// Проходим по первой последовательности и пропускаем те, что есть в excludeSet
		seen := make(map[T]bool) // чтобы не выдавать дубликаты
		q(func(item T) bool {
			if !excludeSet[item] && !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false // остановка от потребителя
				}
			}
			return true
		})
	}
}
