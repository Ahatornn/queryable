package query

import "testing"

// helper: упрощённый способ сообщить об ошибке с контекстом
func assertEqual[T comparable](t *testing.T, got, expected []T, msg string) {
	t.Helper()
	if !slicesEqual(got, expected) {
		t.Errorf("%s: ожидается %v, получено %v", msg, expected, got)
	}
}

// Вспомогательная функция для сравнения слайсов
func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
