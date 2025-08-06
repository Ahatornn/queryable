package query

import "testing"

func TestSkip(t *testing.T) {
	t.Parallel()

	t.Run("n <= 0 → возвращает оригинальную последовательность", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3})
		result := q.Skip(0).ToSlice()
		expected := []int{1, 2, 3}
		if !slicesEqual(result, expected) {
			t.Errorf("Skip(0) = %v, want %v", result, expected)
		}

		result = q.Skip(-1).ToSlice()
		if !slicesEqual(result, expected) {
			t.Errorf("Skip(-1) = %v, want %v", result, expected)
		}
	})

	t.Run("nil Queryable → возвращает nil", func(t *testing.T) {
		t.Parallel()
		var q Queryable[string]
		result := q.Skip(2)
		if result != nil {
			t.Error("ожидался nil для nil Queryable")
		}
	})

	t.Run("n = 1 → пропускает первый элемент", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{10, 20, 30})
		result := q.Skip(1).ToSlice()
		expected := []int{20, 30}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("n = 2 → пропускает два", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{"a", "b", "c", "d"})
		result := q.Skip(2).ToSlice()
		expected := []string{"c", "d"}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("n больше длины → пусто", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2})
		result := q.Skip(5).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("n равно длине → пусто", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3})
		result := q.Skip(3).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("пустой слайс → остаётся пустым", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]float64{})
		result := q.Skip(1).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("nil слайс → остаётся пустым", func(t *testing.T) {
		t.Parallel()
		var data []string = nil
		q := ToQueryable(data)
		result := q.Skip(1).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("сохраняет порядок", func(t *testing.T) {
		t.Parallel()
		type User struct{ Name string }
		q := ToQueryable([]User{{"Alice"}, {"Bob"}, {"Charlie"}, {"Diana"}})
		result := q.Skip(2).ToSlice()
		expected := []User{{"Charlie"}, {"Diana"}}
		if !slicesEqual(result, expected) {
			t.Errorf("порядок нарушен: ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("ленивая остановка", func(t *testing.T) {
		t.Parallel()
		called := 0
		q := Queryable[int](func(yield func(int) bool) {
			for i := 1; i <= 10; i++ {
				if !yield(i) {
					break
				}
				// Только если yield вернул true — считаем, что элемент был "затронут"
				called++
			}
		})

		_ = q.Skip(3).Take(2).ToSlice() // пропускаем 3, берём 2

		// Должно быть: пропущено 3, взято 2 → всего 5 элементов обработано
		if called != 5 {
			t.Errorf("ожидалось 5 вызовов, было %d", called)
		}
	})
}
