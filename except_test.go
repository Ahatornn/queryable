package query

import "testing"

func TestExcept(t *testing.T) {
	t.Run("обычное исключение", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 3, 4})
		q2 := ToQueryable([]int{3, 4, 5})

		result := q1.Except(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "Except: {1,2,3,4} \\ {3,4,5}")
	})

	t.Run("все элементы исключаются", func(t *testing.T) {
		q1 := ToQueryable([]string{"a", "b"})
		q2 := ToQueryable([]string{"a", "b", "c"})

		result := q1.Except(q2).ToSlice()

		expected := []string{}
		assertEqual(t, result, expected, "все исключены")
	})

	t.Run("ничего не исключается", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2})
		q2 := ToQueryable([]int{3, 4})

		result := q1.Except(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "нет пересечения")
	})

	t.Run("дубликаты в первой последовательности", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 2, 3})
		q2 := ToQueryable([]int{2})

		result := q1.Except(q2).ToSlice()

		expected := []int{1, 3} // 2 исключён, дубликаты не важны
		assertEqual(t, result, expected, "без дубликатов и без 2")
	})

	t.Run("пустая первая", func(t *testing.T) {
		q1 := ToQueryable([]int{})
		q2 := ToQueryable([]int{1, 2})

		result := q1.Except(q2).ToSlice()

		expected := []int{}
		assertEqual(t, result, expected, "пустая первая")
	})

	t.Run("пустая вторая", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2})
		q2 := ToQueryable([]int{})

		result := q1.Except(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "пустая вторая")
	})

	t.Run("ленивая остановка", func(t *testing.T) {
		q1 := ToQueryable([]int{1, 2, 3})
		q2 := ToQueryable([]int{1, 4, 5})

		result := q1.Except(q2).Take(1).ToSlice()
		expected := []int{2}
		assertEqual(t, result, expected, "ленивая остановка")
	})
}
