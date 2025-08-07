package query

import "testing"

func TestDistinct(t *testing.T) {
	t.Run("удаление дубликатов", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 2, 3, 3, 4})

		result := q.Distinct().ToSlice()

		expected := []int{1, 2, 3, 4}
		assertEqual(t, result, expected, "дубликаты удалены")
	})

	t.Run("уже уникальные элементы", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{"a", "b", "c"})

		result := q.Distinct().ToSlice()

		expected := []string{"a", "b", "c"}
		assertEqual(t, result, expected, "без изменений")
	})

	t.Run("все одинаковые", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{5, 5, 5, 5})

		result := q.Distinct().ToSlice()

		expected := []int{5}
		assertEqual(t, result, expected, "один элемент")
	})

	t.Run("пустая последовательность", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]float64{})

		result := q.Distinct().ToSlice()
		expected := []float64{}
		assertEqual(t, result, expected, "пусто")
	})

	t.Run("nil слайс → остаётся пустым", func(t *testing.T) {
		t.Parallel()
		var data []string = nil
		q := ToQueryable(data)
		result := q.Distinct().ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("nil Queryable → возвращает пустую последовательность", func(t *testing.T) {
		t.Parallel()
		var q Queryable[string]
		result := q.Distinct().ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("один элемент", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{42})

		result := q.Distinct().ToSlice()

		expected := []int{42}
		assertEqual(t, result, expected, "один элемент")
	})

	t.Run("ленивая остановка", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 1, 2, 3, 2, 4})

		result := q.Distinct().Take(2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "ленивая остановка")
	})
}
