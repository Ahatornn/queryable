package query

import "testing"

func TestAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     []int
		predicate func(int) bool
		want      bool
	}{
		{
			name:      "пустой слайс → true",
			input:     []int{},
			predicate: func(x int) bool { return x > 0 },
			want:      true,
		},
		{
			name:      "nil Queryable → true",
			input:     nil,
			predicate: func(x int) bool { return x > 0 },
			want:      true,
		},
		{
			name:      "все элементы положительные",
			input:     []int{1, 2, 3, 4},
			predicate: func(x int) bool { return x > 0 },
			want:      true,
		},
		{
			name:      "все чётные",
			input:     []int{2, 4, 6},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      true,
		},
		{
			name:      "один отрицательный → false",
			input:     []int{1, 2, -3, 4},
			predicate: func(x int) bool { return x > 0 },
			want:      false,
		},
		{
			name:      "последний элемент не проходит → false",
			input:     []int{2, 4, 6, 7},
			predicate: func(x int) bool { return x%2 == 0 },
			want:      false,
		},
		{
			name:      "первый элемент не проходит → false (и досрочный выход)",
			input:     []int{-1, 2, 3, 4},
			predicate: func(x int) bool { return x > 0 },
			want:      false,
		},
		{
			name:      "все элементы удовлетворяют сложному условию",
			input:     []int{10, 20, 30},
			predicate: func(x int) bool { return x >= 10 && x <= 30 },
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := ToQueryable(tt.input).All(tt.predicate)

			if got != tt.want {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAll_LazyEvaluation(t *testing.T) {
	t.Parallel()

	called := 0
	predicate := func(x int) bool {
		called++
		return x < 3
	}

	input := []int{1, 2, 5, 10, 15} // 5 — первый, кто не проходит
	result := ToQueryable(input).All(predicate)

	if result {
		t.Error("ожидалось false, так как 5 >= 3")
	}
	if called != 3 {
		t.Errorf("предикат должен быть вызван 3 раза (на 1, 2, 5), но вызван %d раз", called)
	}
}

func TestAll_NilQueryable(t *testing.T) {
	t.Parallel()

	var q Queryable[int] // нулевая функция → nil
	got := q.All(func(x int) bool { return true })
	if !got {
		t.Error("ожидалось true для nil Queryable")
	}
}
