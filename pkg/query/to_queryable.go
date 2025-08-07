package query

// ToQueryable Создаёт обёртку для ленивых операций из слайса
func ToQueryable[T comparable](items []T) Queryable[T] {
	return func(yield func(T) bool) {
		for _, item := range items {
			if !yield(item) {
				return
			}
		}
	}
}
