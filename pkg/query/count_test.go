package query

import "testing"

func TestCount(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []string // используем string для разнообразия типов
		want  int
	}{
		{
			name:  "nil Queryable → 0",
			input: nil,
			want:  0,
		},
		{
			name:  "пустой слайс → 0",
			input: []string{},
			want:  0,
		},
		{
			name:  "один элемент → 1",
			input: []string{"hello"},
			want:  1,
		},
		{
			name:  "три элемента → 3",
			input: []string{"a", "b", "c"},
			want:  3,
		},
		{
			name:  "много элементов → 5",
			input: []string{"1", "2", "3", "4", "5"},
			want:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var q Queryable[string]
			if tt.input != nil {
				q = ToQueryable(tt.input)
			}

			got := q.Count()

			if got != tt.want {
				t.Errorf("Count() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCount_FullIteration(t *testing.T) {
	t.Parallel()

	called := 0
	q := Queryable[int](func(yield func(int) bool) {
		for _, x := range []int{10, 20, 30, 40} {
			called++
			yield(x)
		}
	})

	result := q.Count()

	if result != 4 {
		t.Errorf("ожидалось 4, получено %d", result)
	}
	if called != 4 {
		t.Errorf("ожидалось 4 вызова, было %d", called)
	}
}

func TestCount_NilQueryable(t *testing.T) {
	t.Parallel()

	var q Queryable[int] // nil-функция
	got := q.Count()

	if got != 0 {
		t.Errorf("ожидалось 0 для nil Queryable, получено %d", got)
	}
}

func TestCount_LargeSlice(t *testing.T) {
	t.Parallel()

	// Создаём большой слайс
	items := make([]int, 1000)
	for i := range items {
		items[i] = i
	}

	q := ToQueryable(items)
	got := q.Count()

	if got != 1000 {
		t.Errorf("ожидалось 1000, получено %d", got)
	}
}
