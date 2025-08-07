package query

// Except возвращает элементы из текущей последовательности, которых нет в `second`.
// Дубликаты удаляются. Порядок — как в первой последовательности.
func (q Queryable[T]) Except(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		// Обработка second: если nil — считаем, что он пустой
		excludeSet := make(map[T]bool)
		if second != nil {
			second(func(item T) bool {
				excludeSet[item] = true
				return true
			})
		}

		// Обработка q: если nil — ничего не итерируем
		if q == nil {
			return
		}

		seen := make(map[T]bool) // для дедупликации
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
