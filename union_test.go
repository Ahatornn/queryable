package query

import "testing"

func TestUnion(t *testing.T) {
	t.Run("объединение двух слайсов с дубликатами", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 3})
		q2 := ToQueryable([]int{3, 4, 5})

		result := q1.Union(q2).ToSlice()

		expected := []int{1, 2, 3, 4, 5}
		assertEqual(t, result, expected, "объединение с дубликатами")
	})

	t.Run("объединение с пустой второй последовательностью", func(t *testing.T) {
		q1 := ToQueryable([]string{"a", "b"})
		q2 := ToQueryable([]string{})

		result := q1.Union(q2).ToSlice()

		expected := []string{"a", "b"}
		assertEqual(t, result, expected, "пустая вторая")
	})

	t.Run("объединение с пустой первой последовательностью", func(t *testing.T) {
		q1 := ToQueryable([]int{})
		q2 := ToQueryable([]int{1, 2})

		result := q1.Union(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "пустая первая")
	})

	t.Run("обе последовательности пустые", func(t *testing.T) {
		q1 := ToQueryable([]float64{})
		q2 := ToQueryable([]float64{})

		result := q1.Union(q2).ToSlice()
		expected := []float64{}
		assertEqual(t, result, expected, "обе пустые")
	})

	t.Run("полное пересечение", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2})
		q2 := ToQueryable([]int{1, 2})

		result := q1.Union(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "полное пересечение")
	})

	t.Run("ленивая остановка (например, Take)", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 3})
		q2 := ToQueryable([]int{3, 4, 5})

		result := q1.Union(q2).Take(2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "получили только первые 2")
	})
}
