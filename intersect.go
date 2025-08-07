package query

// Intersect возвращает элементы, которые присутствуют и в текущей последовательности,
// и в `second`. Дубликаты удаляются. Порядок — как в первой последовательности.
func (q Queryable[T]) Intersect(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		// Сначала собираем все элементы из second
		secondSet := make(map[T]bool)
		second(func(item T) bool {
			secondSet[item] = true
			return true
		})

		// Теперь проходим по первой последовательности
		seen := make(map[T]bool) // чтобы не дублировать результат
		q(func(item T) bool {
			// Элемент должен быть в second и ещё не выдан
			if secondSet[item] && !seen[item] {
				seen[item] = true
				if !yield(item) {
					return false // остановка от потребителя
				}
			}
			return true
		})
	}
}
