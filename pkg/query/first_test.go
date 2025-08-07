package query

import "testing"

func TestFirst(t *testing.T) {
	t.Parallel()

	t.Run("непустая последовательность → указатель на первый элемент", func(t *testing.T) {
		t.Parallel()
		result := ToQueryable([]int{10, 20, 30}).First()
		if result == nil {
			t.Fatal("ожидался указатель, получил nil")
		}
		if *result != 10 {
			t.Errorf("ожидалось *result == 10, получено %d", *result)
		}
	})

	t.Run("пустой слайс → nil", func(t *testing.T) {
		t.Parallel()
		result := ToQueryable([]int{}).First()
		if result != nil {
			t.Errorf("пустая последовательность должна возвращать nil, получено %v", *result)
		}
	})

	t.Run("nil Queryable → nil", func(t *testing.T) {
		t.Parallel()
		var q Queryable[string]
		result := q.First()
		if result != nil {
			t.Errorf("nil Queryable должен возвращать nil")
		}
	})

	t.Run("nil слайс → nil", func(t *testing.T) {
		t.Parallel()
		var data []float64 = nil
		result := ToQueryable(data).First()
		if result != nil {
			t.Errorf("nil слайс должен возвращать nil")
		}
	})

	t.Run("со структурами", func(t *testing.T) {
		t.Parallel()
		type User struct {
			Name string
			Age  int
		}
		result := ToQueryable([]User{
			{"Alice", 30},
			{"Bob", 25},
		}).First()
		if result == nil {
			t.Fatal("ожидался указатель, получил nil")
		}
		if *result != (User{"Alice", 30}) {
			t.Errorf("ожидался {Alice 30}, получено %+v", *result)
		}
	})

	t.Run("с одним элементом", func(t *testing.T) {
		t.Parallel()
		result := ToQueryable([]bool{true}).First()
		if result == nil || !*result {
			t.Error("ожидался указатель на true")
		}
	})

	t.Run("указатели", func(t *testing.T) {
		t.Parallel()
		a, b := 1, 2
		q := ToQueryable([]*int{&a, &b})
		result := q.First()
		if result == nil || *result == nil {
			t.Fatal("ожидался указатель на указатель")
		}
		if **result != 1 {
			t.Errorf("ожидалось **result == 1, получено %d", **result)
		}
	})
}
