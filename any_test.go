package query

import "testing"

func TestAny(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input []int
		want  bool
	}{
		{
			name:  "пустой слайс → false",
			input: []int{},
			want:  false,
		},
		{
			name:  "один элемент → true",
			input: []int{42},
			want:  true,
		},
		{
			name:  "несколько элементов → true",
			input: []int{1, 2, 3},
			want:  true,
		},
		{
			name:  "nil слайс → false",
			input: nil,
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			q := ToQueryable(tt.input)
			got := q.Any()

			if got != tt.want {
				t.Errorf("Any() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAny_LazyEvaluation(t *testing.T) {
	t.Parallel()

	called := 0
	q := Queryable[int](func(yield func(int) bool) {
		for _, x := range []int{1, 2, 3, 4, 5} {
			called++
			if !yield(x) {
				break
			}
		}
	})

	result := q.Any()

	if !result {
		t.Error("ожидалось true — есть элементы")
	}
	if called != 1 {
		t.Errorf("ожидался 1 вызов (ленивость), но было: %d", called)
	}
}

func TestAny_NilQueryable(t *testing.T) {
	t.Parallel()

	var q Queryable[string] // nil-функция
	got := q.Any()

	if got {
		t.Error("ожидалось false для nil Queryable")
	}
}

func TestAny_WithStruct(t *testing.T) {
	t.Parallel()

	type User struct {
		Name string
		Age  int
	}

	input := []User{{"Alice", 30}, {"Bob", 25}}
	q := ToQueryable(input)

	got := q.Any()

	if !got {
		t.Error("ожидалось true — есть пользователи")
	}
}
