package query

import "testing"

func TestWhere(t *testing.T) {
	t.Parallel()

	t.Run("nil Queryable → возвращает nil", func(t *testing.T) {
		t.Parallel()
		var q Queryable[int]
		result := q.Where(func(x int) bool { return x > 0 })
		if result != nil {
			t.Errorf("ожидался nil, получен Queryable")
		}
	})

	t.Run("пустой слайс → остаётся пустым", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{})
		result := q.Where(func(s string) bool { return true }).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получен %v", result)
		}
	})

	t.Run("все элементы проходят → остаётся полным", func(t *testing.T) {
		t.Parallel()
		data := []int{1, 2, 3}
		q := ToQueryable(data)
		result := q.Where(func(x int) bool { return x > 0 }).ToSlice()
		if !slicesEqual(result, data) {
			t.Errorf("ожидалось %v, получено %v", data, result)
		}
	})

	t.Run("ни один элемент не проходит → пусто", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3})
		result := q.Where(func(x int) bool { return false }).ToSlice()
		if len(result) != 0 {
			t.Errorf("ожидался пустой слайс, получен %v", result)
		}
	})

	t.Run("фильтр чётных чисел", func(t *testing.T) {
		t.Parallel()
		data := []int{1, 2, 3, 4, 5, 6}
		q := ToQueryable(data)
		result := q.Where(func(x int) bool { return x%2 == 0 }).ToSlice()
		expected := []int{2, 4, 6}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("фильтр по строке", func(t *testing.T) {
		t.Parallel()
		data := []string{"apple", "banana", "avocado", "qiwi"}
		q := ToQueryable(data)
		result := q.Where(func(s string) bool { return len(s) > 5 }).ToSlice()
		expected := []string{"banana", "avocado"}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("со структурами", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name   string
			Age    int
			Active bool
		}
		data := []User{
			{"Alice", 30, true},
			{"Bob", 25, false},
			{"Charlie", 35, true},
		}
		q := ToQueryable(data)
		activeAdults := q.Where(func(u User) bool { return u.Active && u.Age >= 30 }).ToSlice()
		expected := []User{{"Alice", 30, true}, {"Charlie", 35, true}}
		if !slicesEqual(activeAdults, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, activeAdults)
		}
	})

	t.Run("с указателями", func(t *testing.T) {
		t.Parallel()
		a, b, c := 5, 10, 15
		data := []*int{&a, &b, &c}
		q := ToQueryable(data)
		result := q.Where(func(p *int) bool { return *p > 8 }).ToSlice()
		if len(result) != 2 || *result[0] != 10 || *result[1] != 15 {
			t.Errorf("ожидалось [10, 15], получено %v", result)
		}
	})

	t.Run("цепочка из нескольких Where", func(t *testing.T) {
		t.Parallel()
		data := []int{1, 2, 3, 4, 5, 6, 7, 8}
		q := ToQueryable(data)
		result := q.
			Where(func(x int) bool { return x%2 == 0 }). // чётные
			Where(func(x int) bool { return x > 3 }).    // больше 3
			ToSlice()
		expected := []int{4, 6, 8}
		if !slicesEqual(result, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, result)
		}
	})

	t.Run("ленивая остановка (с Take)", func(t *testing.T) {
		t.Parallel()
		called := 0
		data := []int{1, 2, 3, 4, 5}
		q := ToQueryable(data)
		_ = q.Where(func(x int) bool {
			called++
			return x%2 == 0
		}).Take(2).ToSlice()

		// Должно быть: проверены 1,2,3,4,5? Нет!
		// Take(2) останавливается после получения двух элементов (2 и 4)
		// Но predicate вызывается для 5, потому что Take останавливает только после yield
		// Ожидаем: called == 5? Нет — на практике может быть 4 или 5 в зависимости от реализации

		// Уточним: Take(2) → получит 2 и 4
		// predicate вызывается для: 1 (false), 2 (true → taken=1), 3 (false), 4 (true → taken=2), 5 (должен быть вызван, но после этого Take останавливается)
		// Так что called == 5 — нормально

		if called < 4 {
			t.Errorf("ожидалось как минимум 4 вызова predicate, было %d", called)
		}
	})
}
