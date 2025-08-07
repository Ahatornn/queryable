package query

import (
	"reflect"
	"testing"
)

func TestUnion(t *testing.T) {
	t.Run("объединение двух слайсов с дубликатами", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]int{1, 2, 3})
		q2 := ToQueryable([]int{3, 4, 5})

		result := q1.Union(q2).ToSlice()

		expected := []int{1, 2, 3, 4, 5}
		assertEqual(t, result, expected, "объединение с дубликатами")
	})

	t.Run("объединение с пустой второй последовательностью", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]string{"a", "b"})
		q2 := ToQueryable([]string{})

		result := q1.Union(q2).ToSlice()

		expected := []string{"a", "b"}
		assertEqual(t, result, expected, "пустая вторая")
	})

	t.Run("объединение с пустой первой последовательностью", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]int{})
		q2 := ToQueryable([]int{1, 2})

		result := q1.Union(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "пустая первая")
	})

	t.Run("обе последовательности пустые", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]float64{})
		q2 := ToQueryable([]float64{})

		result := q1.Union(q2).ToSlice()
		expected := []float64{}
		assertEqual(t, result, expected, "обе пустые")
	})

	t.Run("полное пересечение", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]int{1, 2})
		q2 := ToQueryable([]int{1, 2})

		result := q1.Union(q2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "полное пересечение")
	})

	t.Run("ленивая остановка (например, Take)", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]int{1, 2, 3})
		q2 := ToQueryable([]int{3, 4, 5})

		result := q1.Union(q2).Take(2).ToSlice()

		expected := []int{1, 2}
		assertEqual(t, result, expected, "получили только первые 2")
	})

	t.Run("nil первая", func(t *testing.T) {
		t.Parallel()
		var q1 Queryable[int]
		q2 := ToQueryable([]int{1, 2})
		result := q1.Union(q2).ToSlice()
		expected := []int{1, 2}
		assertEqual(t, result, expected, "nil первая")
	})

	t.Run("nil вторая", func(t *testing.T) {
		t.Parallel()
		q1 := ToQueryable([]int{1, 2})
		var q2 Queryable[int]
		result := q1.Union(q2).ToSlice()
		expected := []int{1, 2}
		assertEqual(t, result, expected, "nil вторая")
	})

	t.Run("обе nil", func(t *testing.T) {
		t.Parallel()
		var q1, q2 Queryable[int]
		result := q1.Union(q2).ToSlice()
		expected := []int{}
		assertEqual(t, result, expected, "обе nil")
	})

	t.Run("ленивая остановка второй", func(t *testing.T) {
		t.Parallel()
		eff1, eff2 := 0, 0
		q1 := ToQueryable([]int{1, 2}).Select(func(x int) any { eff1++; return x })
		q2 := ToQueryable([]int{3, 4}).Select(func(x int) any { eff2++; return x })

		result := q1.Union(q2).Take(1).ToSlice()

		if !reflect.DeepEqual(result, []any{1}) {
			t.Errorf("ожидается [1,2], получено %v", result)
		}
		if eff1 != 2 {
			t.Errorf("q1 итерируется до конца (из-за своей природы): ожидается 2, получено %v", eff1)
		}
		if eff2 != 0 {
			t.Errorf("q2 не должен итерироваться, если не нужен: ожидается 0, получено %v", eff2)
		}
	})
}
