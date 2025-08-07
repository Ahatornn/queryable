package query

import "testing"

func TestRepeat(t *testing.T) {
	t.Parallel()

	t.Run("повтор 3 раза", func(t *testing.T) {
		t.Parallel()
		slice := Repeat("hello", 3).ToSlice()
		expected := []string{"hello", "hello", "hello"}
		if len(slice) != 3 {
			t.Errorf("ожидалось 3 элемента, получено %d", len(slice))
		}
		for i, v := range expected {
			if slice[i] != v {
				t.Errorf("элемент %d: ожидалось %q, получено %q", i, v, slice[i])
			}
		}
	})

	t.Run("count = 0 → пусто", func(t *testing.T) {
		t.Parallel()
		q := Repeat(42, 0)
		if q.Any() {
			t.Error("Repeat(..., 0) должен быть пустым")
		}
	})

	t.Run("count < 0 → пусто", func(t *testing.T) {
		t.Parallel()
		q := Repeat(true, -1)
		if q.Any() {
			t.Error("Repeat с отрицательным count должен быть пустым")
		}
	})

	t.Run("ленивая остановка", func(t *testing.T) {
		t.Parallel()
		called := 0
		q := Repeat(1, 5)
		_ = q.Where(func(x int) bool {
			called++
			return x <= 1 // останавливаем цепочку
		}).Any()

		// Должен вызваться только один раз (первый элемент)
		if called != 1 {
			t.Errorf("ожидался 1 вызов Where, было %d", called)
		}
	})
}
