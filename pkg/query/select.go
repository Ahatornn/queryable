package query

// Select defines the projection. It is difficult to use now because of any
func (q Queryable[T]) Select(mapper func(T) any) Queryable[any] {
	if q == nil {
		return Empty[any]()
	}
	return func(yield func(any) bool) {
		q(func(item T) bool {
			return yield(mapper(item))
		})
	}
}
