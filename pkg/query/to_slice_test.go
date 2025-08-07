package query

import "testing"

func TestToSlice(t *testing.T) {
	t.Parallel()

	t.Run("nil Queryable → возвращает nil", func(t *testing.T) {
		t.Parallel()
		var q Queryable[int]
		result := q.ToSlice()
		if result != nil {
			t.Errorf("ожидался nil, получено %v", result)
		}
	})

	t.Run("пустой слайс → пустой слайс", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{})
		result := q.ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получено %v", result)
		}
		if result == nil {
			t.Error("ожидался пустой слайс, не nil")
		}
	})

	t.Run("непустой слайс → все элементы", func(t *testing.T) {
		t.Parallel()
		data := []int{10, 20, 30}
		q := ToQueryable(data)
		result := q.ToSlice()
		if !slicesEqual(result, data) {
			t.Errorf("ожидалось %v, получено %v", data, result)
		}
	})

	t.Run("сохраняет порядок", func(t *testing.T) {
		t.Parallel()
		data := []string{"z", "a", "b"}
		q := ToQueryable(data)
		result := q.ToSlice()
		if !slicesEqual(result, data) {
			t.Errorf("порядок нарушен: ожидалось %v, получено %v", data, result)
		}
	})

	t.Run("с указателями", func(t *testing.T) {
		t.Parallel()
		a, b, c := 1, 2, 3
		data := []*int{&a, &b, &c}
		q := ToQueryable(data)
		result := q.ToSlice()
		if len(result) != 3 {
			t.Errorf("ожидалось 3 указателя, получено %d", len(result))
		}
		for i := range data {
			if result[i] != data[i] || *result[i] != *data[i] {
				t.Errorf("указатель %d не совпадает", i)
			}
		}
	})

	t.Run("со структурами", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name   string
			Active bool
		}
		data := []User{
			{"Alice", true},
			{"Bob", false},
		}
		q := ToQueryable(data)
		result := q.ToSlice()
		if !slicesEqual(result, data) {
			t.Errorf("данные не совпадают: %v ≠ %v", data, result)
		}
	})

	t.Run("начальная ёмкость >= 16 (если элементов мало)", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2})
		result := q.ToSlice()
		if cap(result) < 16 {
			t.Errorf("ожидалась ёмкость >= 16, получено %d", cap(result))
		}
		if len(result) != 2 {
			t.Errorf("длина должна быть 2, получено %d", len(result))
		}
	})

	t.Run("ёмкость растёт при большом количестве элементов", func(t *testing.T) {
		t.Parallel()
		data := make([]int, 100)
		for i := range data {
			data[i] = i
		}
		q := ToQueryable(data)
		result := q.ToSlice()
		if len(result) != 100 {
			t.Errorf("ожидалось 100 элементов, получено %d", len(result))
		}
		if cap(result) < 100 {
			t.Errorf("ёмкость должна быть >= 100, получено %d", cap(result))
		}
	})

	t.Run("ленивая остановка (например, с Take)", func(t *testing.T) {
		t.Parallel()
		called := 0
		q := Queryable[int](func(yield func(int) bool) {
			for i := 1; i <= 10; i++ {
				if !yield(i) {
					break
				}
				called++
			}
		})

		_ = q.Take(3).ToSlice()

		// Take(3) → yield вызван 3 раза → called == 3
		if called != 3 {
			t.Errorf("ожидалось 3 вызова, было %d", called)
		}
	})

	t.Run("интеграция с Where и Concat", func(t *testing.T) {
		t.Parallel()
		a := ToQueryable([]int{1, 2})
		b := ToQueryable([]int{3, 4})
		result := a.Concat(b).
			Where(func(x int) bool { return x%2 == 0 }).
			ToSlice()
		expected := []int{2, 4}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})
}
