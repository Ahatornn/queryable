package query

// All Проверяет все элементы на удовлетворение условию
func (q Queryable[T]) All(predicate func(T) bool) bool {
	if q == nil {
		return true // по соглашению: все элементы пустого множества удовлетворяют любому условию
	}
	var any bool
	var all = true
	q(func(item T) bool {
		any = true
		if !predicate(item) {
			all = false
			return false // останавливаем
		}
		return true
	})
	return !any || all // если нет элементов — true, иначе проверяем all
}
