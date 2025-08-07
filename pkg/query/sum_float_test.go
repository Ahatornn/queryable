package query

import "testing"

func TestSumFloat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []float32             // используем []float32 для совместимости с SumFloat
		selector func(float32) float32 // функция извлечения значения типа float32
		want     float32
	}{
		{
			name:     "nil Queryable → 0",
			input:    nil,
			selector: func(x float32) float32 { return x },
			want:     0.0,
		},
		{
			name:     "пустой слайс → 0",
			input:    []float32{},
			selector: func(x float32) float32 { return x },
			want:     0.0,
		},
		{
			name:     "один элемент → сам элемент",
			input:    []float32{3.14},
			selector: func(x float32) float32 { return x },
			want:     3.14,
		},
		{
			name:     "сумма положительных чисел",
			input:    []float32{1.5, 2.5, 3.0},
			selector: func(x float32) float32 { return x },
			want:     7.0, // 1.5 + 2.5 + 3.0
		},
		{
			name:     "отрицательные и положительные",
			input:    []float32{-1.5, 2.0, -0.5},
			selector: func(x float32) float32 { return x },
			want:     0.0, // -1.5 + 2.0 - 0.5 = 0
		},
		{
			name:     "все нули → 0",
			input:    []float32{0.0, 0.0, 0.0},
			selector: func(x float32) float32 { return x },
			want:     0.0,
		},
		{
			name:     "с округлением (проверка float32)",
			input:    []float32{0.1, 0.2, 0.3},
			selector: func(x float32) float32 { return x },
			want:     0.6,
		},
		{
			name:     "преобразование: квадраты значений",
			input:    []float32{2.0, 3.0},
			selector: func(x float32) float32 { return x * x }, // 4 + 9 = 13
			want:     13.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var q Queryable[float32]
			if tt.input != nil {
				q = ToQueryable(tt.input)
			}

			got := q.SumFloat(tt.selector)

			// Используем сравнение с погрешностью для float
			const epsilon = 1e-6
			if got < tt.want-epsilon || got > tt.want+epsilon {
				t.Errorf("SumFloat() = %v, want %v (с точностью до %v)", got, tt.want, epsilon)
			}
		})
	}
}

type Product struct {
	Price float32
	Name  string
}

func TestSumFloat_WithStruct(t *testing.T) {
	t.Parallel()

	products := []Product{
		{Name: "Milk", Price: 1.99},
		{Name: "Bread", Price: 2.50},
		{Name: "Butter", Price: 3.25},
	}

	q := ToQueryable(products)
	total := q.SumFloat(func(p Product) float32 { return p.Price })

	const epsilon = 1e-6
	if total < 7.74-epsilon || total > 7.74+epsilon {
		t.Errorf("SumFloat(Price) = %v, want ~7.74", total)
	}
}
