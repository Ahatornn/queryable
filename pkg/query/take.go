package query

// Take the first n elements
func (q Queryable[T]) Take(n int) Queryable[T] {
	if q == nil || n <= 0 {
		return Empty[T]()
	}
	return func(yield func(T) bool) {
		var taken int
		q(func(item T) bool {
			if taken >= n {
				return false // останавливаем внешний итератор
			}
			if !yield(item) {
				return false
			}
			taken++
			return true
		})
	}
}
