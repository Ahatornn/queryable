package query

// Union объединяет текущую последовательность с другой, удаляя дубликаты.
// Возвращает новую Queryable[T], содержащую уникальные элементы из обеих последовательностей.
// Порядок: сначала уникальные элементы из q, затем из second, которых не было в q.
//
// Пустые (nil) последовательности обрабатываются как пустые.
// ⚠️ Внимание: вторая последовательность итерируется только если это необходимо
// (если внешний потребитель ещё ожидает данные).
func (q Queryable[T]) Union(second Queryable[T]) Queryable[T] {
	return func(yield func(T) bool) {
		seen := make(map[T]bool)
		stopped := false

		// Обёртка вокруг yield, чтобы отслеживать остановку
		yieldGuard := func(item T) bool {
			if stopped {
				return false
			}
			if !yield(item) {
				stopped = true
				return false
			}
			return true
		}

		// Первая последовательность
		if q != nil {
			q(func(item T) bool {
				if !seen[item] {
					seen[item] = true
					if !yieldGuard(item) {
						return false
					}
				}
				return true
			})
		}

		// Вторая последовательность — только если ещё не остановились
		if !stopped && second != nil {
			second(func(item T) bool {
				if !seen[item] {
					seen[item] = true
					return yieldGuard(item)
				}
				return true
			})
		}
	}
}
