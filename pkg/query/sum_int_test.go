package query

import "testing"

func TestSumInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []int         // используем []int, так как SumInt суммирует int
		selector func(int) int // функция извлечения (в данном случае — может быть identity или другая)
		want     int
	}{
		{
			name:     "nil Queryable → 0",
			input:    nil,
			selector: func(x int) int { return x },
			want:     0,
		},
		{
			name:     "пустой слайс → 0",
			input:    []int{},
			selector: func(x int) int { return x },
			want:     0,
		},
		{
			name:     "один элемент → значение этого элемента",
			input:    []int{42},
			selector: func(x int) int { return x },
			want:     42,
		},
		{
			name:     "три элемента → сумма",
			input:    []int{10, 20, 30},
			selector: func(x int) int { return x },
			want:     60,
		},
		{
			name:     "отрицательные числа → корректная сумма",
			input:    []int{-5, 0, 5},
			selector: func(x int) int { return x },
			want:     0,
		},
		{
			name:     "суммирование по полю: например, квадраты",
			input:    []int{1, 2, 3},
			selector: func(x int) int { return x * x }, // суммируем квадраты: 1 + 4 + 9 = 14
			want:     14,
		},
		{
			name:     "все нули → 0",
			input:    []int{0, 0, 0},
			selector: func(x int) int { return x },
			want:     0,
		},
		{
			name:     "большой список → корректная сумма",
			input:    []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, // 10 единиц
			selector: func(x int) int { return x },
			want:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var q Queryable[int]
			if tt.input != nil {
				q = ToQueryable(tt.input)
			}

			got := q.SumInt(tt.selector)

			if got != tt.want {
				t.Errorf("SumInt() = %d, want %d", got, tt.want)
			}
		})
	}
}

type Person struct {
	Age  int
	Name string
}

func TestSumInt_WithStruct(t *testing.T) {
	t.Parallel()

	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	q := ToQueryable(people)
	totalAge := q.SumInt(func(p Person) int { return p.Age })

	if totalAge != 90 {
		t.Errorf("SumInt(Age) = %d, want 90", totalAge)
	}
}
