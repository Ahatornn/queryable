package query

import "testing"

func TestEmpty(t *testing.T) {
	t.Parallel()

	// Создаём пустую последовательность
	q := Empty[int]()

	// Проверяем, что Any() → false
	if q.Any() {
		t.Error("Empty[int]() должен возвращать false для Any()")
	}

	// Проверяем, что Count() → 0
	if q.Count() != 0 {
		t.Errorf("Empty[int]().Count() = %d, want 0", q.Count())
	}

	// Проверяем, что ToSlice() → пустой слайс
	slice := q.ToSlice()
	if len(slice) != 0 {
		t.Errorf("Empty[int]().ToSlice() имеет длину %d, want 0", len(slice))
	}
	if cap(slice) == 0 {
		// Это ожидаемо, но не обязательно
	}

	// Проверяем All и Where
	if !q.All(func(x int) bool { return true }) {
		t.Error("Empty().All(...) должен быть true (vacuous truth)")
	}

	filtered := q.Where(func(x int) bool { return x > 0 })
	if filtered.Any() {
		t.Error("Empty().Where(...) должен оставаться пустым")
	}
}
