package query

import "testing"

func TestTake(t *testing.T) {
	t.Parallel()

	t.Run("n <= 0 → возвращает пустую последовательность", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3})
		result := q.Take(0).ToSlice()
		if len(result) != 0 {
			t.Errorf("Take(0) должен быть пустым, получено %v", result)
		}

		result = q.Take(-1).ToSlice()
		if len(result) != 0 {
			t.Errorf("Take(-1) должен быть пустым, получено %v", result)
		}
	})

	t.Run("nil Queryable → возвращает пустую последовательность", func(t *testing.T) {
		t.Parallel()
		var q Queryable[string]
		result := q.Take(2).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("n = 1 → берёт первый элемент", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{10, 20, 30})
		result := q.Take(1).ToSlice()
		expected := []int{10}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("n = 2 → берёт два", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{"a", "b", "c", "d"})
		result := q.Take(2).ToSlice()
		expected := []string{"a", "b"}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("n больше длины → берёт все", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2})
		result := q.Take(5).ToSlice()
		expected := []int{1, 2}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("n равно длине → берёт все", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3})
		result := q.Take(3).ToSlice()
		expected := []int{1, 2, 3}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("пустой слайс → остаётся пустым", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]float64{})
		result := q.Take(1).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("nil слайс → остаётся пустым", func(t *testing.T) {
		t.Parallel()
		var data []string = nil
		q := ToQueryable(data)
		result := q.Take(1).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
	})

	t.Run("сохраняет порядок", func(t *testing.T) {
		t.Parallel()
		type User struct{ Name string }
		q := ToQueryable([]User{{"Alice"}, {"Bob"}, {"Charlie"}})
		result := q.Take(2).ToSlice()
		expected := []User{{"Alice"}, {"Bob"}}
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

		_ = q.Take(3).ToSlice() // берём только 3

		// Должно быть: yield вызван 3 раза → called == 3
		if called != 3 {
			t.Errorf("ожидалось 3 вызова, было %d", called)
		}
	})

	t.Run("в комбинации с Skip", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3, 4, 5})
		result := q.Skip(2).Take(2).ToSlice() // пропустить 2, взять 2
		expected := []int{3, 4}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})
}
