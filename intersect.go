package query

// Intersect возвращает элементы, которые присутствуют и в текущей последовательности,
// и в `second`. Дубликаты удаляются. Порядок — как в первой последовательности.
//
// Пустые (nil) последовательности обрабатываются как пустые наборы.
// ⚠️ Внимание: эта операция полностью буферизует `second` в память (создаёт map),
// поэтому она не ленивая по отношению к `second`.
func (q Queryable[T]) Intersect(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		// Собираем множество из second
		secondSet := make(map[T]bool)
		if second != nil {
			second(func(item T) bool {
				secondSet[item] = true
				return true
			})
		}

		// Если q nil — ничего не делаем
		if q == nil {
			return
		}

		seen := make(map[T]bool) // для дедупликации
		q(func(item T) bool {
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
