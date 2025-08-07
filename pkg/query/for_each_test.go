package query

import "testing"

func TestForEach(t *testing.T) {
	t.Parallel()

	t.Run("nil Queryable → не паникует", func(t *testing.T) {
		t.Parallel()
		var q Queryable[int]
		called := false
		q.ForEach(func(item int) {
			called = true
		})
		if called {
			t.Error("action не должен вызываться для nil Queryable")
		}
	})

	t.Run("пустой слайс → action не вызывается", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{})
		called := false
		q.ForEach(func(item string) {
			called = true
		})
		if called {
			t.Error("action не должен вызываться для пустой последовательности")
		}
	})

	t.Run("nil слайс → action не вызывается", func(t *testing.T) {
		t.Parallel()
		var data []float64 = nil
		q := ToQueryable(data)
		called := false
		q.ForEach(func(item float64) {
			called = true
		})
		if called {
			t.Error("action не должен вызываться для nil слайса")
		}
	})

	t.Run("один элемент → action вызывается один раз", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{42})
		var calledWith int
		q.ForEach(func(item int) {
			calledWith = item
		})
		if calledWith != 42 {
			t.Errorf("ожидалось 42, получено %d", calledWith)
		}
	})

	t.Run("несколько элементов → action вызывается для каждого", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{1, 2, 3})
		var log []int
		q.ForEach(func(item int) {
			log = append(log, item)
		})
		expected := []int{1, 2, 3}
		if !slicesEqual(log, expected) {
			t.Errorf("ожидалось %v, получено %v", expected, log)
		}
	})

	t.Run("в правильном порядке", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{"a", "b", "c"})
		var order []string
		q.ForEach(func(item string) {
			order = append(order, item)
		})
		expected := []string{"a", "b", "c"}
		if !slicesEqual(order, expected) {
			t.Errorf("порядок нарушен: ожидалось %v, получено %v", expected, order)
		}
	})

	t.Run("с указателями", func(t *testing.T) {
		t.Parallel()
		a, b := 10, 20
		q := ToQueryable([]*int{&a, &b})
		var sum int
		q.ForEach(func(item *int) {
			if item != nil {
				sum += *item
			}
		})
		if sum != 30 {
			t.Errorf("ожидалось 30, получено %d", sum)
		}
	})

	t.Run("со структурами", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name   string
			Active bool
		}
		q := ToQueryable([]User{
			{"Alice", true},
			{"Bob", false},
		})
		var activeCount int
		q.ForEach(func(u User) {
			if u.Active {
				activeCount++
			}
		})
		if activeCount != 1 {
			t.Errorf("ожидался 1 активный, получено %d", activeCount)
		}
	})

	t.Run("action с побочным эффектом (например, изменение внешней переменной)", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]bool{true, false, true})
		count := 0
		q.ForEach(func(b bool) {
			if b {
				count++
			}
		})
		if count != 2 {
			t.Errorf("ожидалось 2 true, получено %d", count)
		}
	})
}
