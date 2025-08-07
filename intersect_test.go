package query

import "testing"

// queryable_test.go (дополнение)
func TestIntersect(t *testing.T) {
	t.Run("пересечение с общими элементами", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 3, 4})
		q2 := ToQueryable([]int{3, 4, 5, 6})

		result := q1.Intersect(q2).ToSlice()

		expected := []int{3, 4}
		assertEqual(t, result, expected, "общие элементы")
	})

	t.Run("нет общих элементов", func(t *testing.T) {
		q1 := ToQueryable([]string{"a", "b"})
		q2 := ToQueryable([]string{"c", "d"})

		result := q1.Intersect(q2).ToSlice()

		expected := []string{}
		assertEqual(t, result, expected, "нет пересечения")
	})

	t.Run("полное совпадение", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2})
		q2 := ToQueryable([]int{1, 2})

		result := q1.Intersect(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "полное совпадение")
	})

	t.Run("дубликаты в первой последовательности", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 2, 3})
		q2 := ToQueryable([]int{2, 3, 4})

		result := q1.Intersect(q2).ToSlice()

		expected := []int{2, 3} // дубликат 2 не повторяется
		assertEqual(t, result, expected, "без дубликатов")
	})

	t.Run("одна из последовательностей пуста", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2})
		q2 := ToQueryable([]int{})

		result := q1.Intersect(q2).ToSlice()

		expected := []int{}
		assertEqual(t, result, expected, "пустая вторая")
	})

	t.Run("ленивая остановка", func(t *testing.T) {
		q1 := ToQueryable([]int{2, 3, 4})
		q2 := ToQueryable([]int{1, 2, 3})

		result := q1.Intersect(q2).Take(1).ToSlice()

		expected := []int{2}
		assertEqual(t, result, expected, "без дубликатов")
	})
}
