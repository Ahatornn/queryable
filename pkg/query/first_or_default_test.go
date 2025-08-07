// query/first_or_default_test.go
package query

import "testing"

func TestFirstOrDefault(t *testing.T) {
	t.Parallel()

	t.Run("непустая последовательность → первый элемент", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{42, 100, 200})
		got := q.FirstOrDefault(0)
		if got != 42 {
			t.Errorf("ожидалось 42, получено %d", got)
		}
	})

	t.Run("пустой слайс → defaultValue", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{})
		got := q.FirstOrDefault("default")
		if got != "default" {
			t.Errorf("ожидалось 'default', получено %q", got)
		}
	})

	t.Run("nil Queryable → defaultValue", func(t *testing.T) {
		t.Parallel()
		var q Queryable[int] // nil
		got := q.FirstOrDefault(99)
		if got != 99 {
			t.Errorf("ожидалось 99, получено %d", got)
		}
	})

	t.Run("nil слайс → defaultValue", func(t *testing.T) {
		t.Parallel()
		var data []float64 = nil
		q := ToQueryable(data)
		got := q.FirstOrDefault(3.14)
		if got != 3.14 {
			t.Errorf("ожидалось 3.14, получено %g", got)
		}
	})

	t.Run("со структурами", func(t *testing.T) {
		t.Parallel()
		type Person struct {
			Name string
			Age  int
		}

		q := ToQueryable([]Person{
			{"Alice", 30},
			{"Bob", 25},
		})
		defaultPerson := Person{"Unknown", 0}
		got := q.FirstOrDefault(defaultPerson)

		if got != (Person{"Alice", 30}) {
			t.Errorf("ожидался {Alice 30}, получено %+v", got)
		}
	})

	t.Run("указатели на int", func(t *testing.T) {
		t.Parallel()
		a, b := 10, 20
		q := ToQueryable([]*int{&a, &b})
		var defaultPtr *int
		got := q.FirstOrDefault(defaultPtr)

		if got == nil {
			t.Fatal("ожидался указатель на 10, получил nil")
		}
		if *got != 10 {
			t.Errorf("ожидалось *got == 10, получено %d", *got)
		}
	})

	t.Run("нулевое значение как defaultValue (например, 0)", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]int{}) // пустой
		got := q.FirstOrDefault(0)
		if got != 0 {
			t.Errorf("ожидалось 0, получено %d", got)
		}
	})

	t.Run("пустая строка как defaultValue", func(t *testing.T) {
		t.Parallel()
		q := ToQueryable([]string{})
		got := q.FirstOrDefault("")
		if got != "" {
			t.Errorf("ожидалась пустая строка, получена %q", got)
		}
	})
}
