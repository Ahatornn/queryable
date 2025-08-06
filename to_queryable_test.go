package query

import "testing"

func TestToQueryable(t *testing.T) {
	t.Parallel()

	t.Run("пустой слайс → пустая последовательность", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{})
		if q.Any() {
			t.Error("пустой слайс не должен содержать элементы")
		}
		if q.Count() != 0 {
			t.Errorf("Count() = %d, want 0", q.Count())
		}
	})

	t.Run("nil слайс → пустая последовательность", func(t *testing.T) {
		t.Parallel()
		var data []string = nil
		q := ToQueryable(data)
		if q.Any() {
			t.Error("nil слайс не должен содержать элементы")
		}
		if q.Count() != 0 {
			t.Errorf("Count() = %d, want 0", q.Count())
		}
	})

	t.Run("непустой слайс → все элементы доступны", func(t *testing.T) {
		t.Parallel()
		data := []int{10, 20, 30}
		q := ToQueryable(data)
		result := q.ToSlice()
		if len(result) != 3 {
			t.Errorf("ожидалось 3 элемента, получено %d", len(result))
		}
		for i, v := range data {
			if result[i] != v {
				t.Errorf("элемент %d: ожидалось %d, получено %d", i, v, result[i])
			}
		}
	})

	t.Run("сохраняет порядок", func(t *testing.T) {
		t.Parallel()
		data := []string{"first", "second", "third"}
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
			if *result[i] != *data[i] {
				t.Errorf("указатель %d не совпадает", i)
			}
		}
	})

	t.Run("со структурами", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name string
			Age  int
		}
		data := []User{
			{"Alice", 30},
			{"Bob", 25},
		}
		q := ToQueryable(data)
		result := q.ToSlice()
		if len(result) != 2 {
			t.Errorf("ожидалось 2 пользователя, получено %d", len(result))
		}
		if result[0] != data[0] || result[1] != data[1] {
			t.Errorf("данные не совпадают: %v ≠ %v", result, data)
		}
	})

	t.Run("ленивая остановка (например, с Take)", func(t *testing.T) {
		t.Parallel()
		called := 0
		data := []int{1, 2, 3, 4, 5}
		q := ToQueryable(data)
		_ = q.Where(func(x int) bool {
			called++
			return true
		}).Take(2).ToSlice()

		// Должно быть: Take(2) → остановка после 2 элементов
		if called != 3 {
			t.Errorf("ожидалось 2 вызова Where, было %d", called)
		}
	})

	t.Run("ленивая остановка (с First)", func(t *testing.T) {
		t.Parallel()
		called := 0
		data := []string{"a", "b", "c"}
		q := ToQueryable(data)
		_ = q.Where(func(s string) bool {
			called++
			return true
		}).First()

		// First() останавливается после первого элемента
		if called != 1 {
			t.Errorf("ожидался 1 вызов, было %d", called)
		}
	})

	t.Run("интеграция с Concat", func(t *testing.T) {
		t.Parallel()
		a := ToQueryable([]int{1, 2})
		b := ToQueryable([]int{3, 4})
		result := a.Concat(b).ToSlice()
		expected := []int{1, 2, 3, 4}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})
}
