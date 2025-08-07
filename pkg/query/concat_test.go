package query

import "testing"

func TestConcat_Method(t *testing.T) {
	t.Parallel()

	t.Run("две непустые последовательности", func(t *testing.T) {
		t.Parallel()
		a := ToQueryable([]int{1, 2})
		b := ToQueryable([]int{3, 4, 5})
		result := a.Concat(b).ToSlice()
		expected := []int{1, 2, 3, 4, 5}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("первая пустая", func(t *testing.T) {
		t.Parallel()
		empty := Empty[int]()
		second := ToQueryable([]int{10, 20})
		result := empty.Concat(second).ToSlice()
		expected := []int{10, 20}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("вторая пустая", func(t *testing.T) {
		t.Parallel()
		first := ToQueryable([]int{1, 2})
		empty := Empty[int]()
		result := first.Concat(empty).ToSlice()
		expected := []int{1, 2}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("обе пустые", func(t *testing.T) {
		t.Parallel()
		result := Empty[int]().Concat(Empty[int]()).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("nil receiver", func(t *testing.T) {
		t.Parallel()
		var q Queryable[int] // nil
		second := ToQueryable([]int{42})
		result := q.Concat(second).ToSlice()
		expected := []int{42}
		if !slicesEqual(result, expected) {
			t.Errorf("nil.Concat(second) = %v, want %v", result, expected)
		}
	})

	t.Run("nil second", func(t *testing.T) {
		t.Parallel()
		first := ToQueryable([]int{1, 2})
		var second Queryable[int] // nil
		result := first.Concat(second).ToSlice()
		expected := []int{1, 2}
		if !slicesEqual(result, expected) {
			t.Errorf("first.Concat(nil) = %v, want %v", result, expected)
		}
	})

	t.Run("ленивая остановка: second не вызывается, если не нужно", func(t *testing.T) {
		t.Parallel()

		first := ToQueryable([]int{1, 2})

		calledInsideYield := false
		second := Queryable[int](func(yield func(int) bool) {
			// Передаём элемент
			if !yield(3) {
				return
			}
			// Только если yield вернул true — считаем, что элемент был "затронут"
			calledInsideYield = true
		})

		_ = first.Concat(second).Take(2).ToSlice()

		if calledInsideYield {
			t.Error("вторая последовательность не должна передавать элементы при Take(2)")
		}
	})
}
