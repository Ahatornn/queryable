package query

// Repeat создаёт Queryable, возвращающий элемент `item` ровно `count` раз.
func Repeat[T any](item T, count int) Queryable[T] {
	return func(yield func(T) bool) {
		for i := 0; i < count; i++ {
			if !yield(item) {
				return
			}
		}
	}
}
