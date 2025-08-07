# queryable: LINQ-style Iterators for Go
[![Тесты](https://github.com/ahatornn/queryable/actions/workflows/test.yml/badge.svg)](https://github.com/ahatornn/queryable/actions/workflows/test.yml)

`queryable` — это библиотека для Go, вдохновлённая C# LINQ (Language Integrated Query), которая позволяет работать с данными в функциональном стиле с использованием **дженериков** и **ленивых итераторов**.

С её помощью ты можешь писать чистый, читаемый и выразительный код для фильтрации, преобразования, сортировки и агрегации данных — без циклов `for` и временных срезов.

> Похоже на LINQ, но на Go. Ленивые вычисления. Поддержка дженериков. Безопасно по типам.

---

## ✨ Особенности

- ✅ **LINQ-подобные методы**: `Where`, `Select`, `Take`, `Skip`, `First`, `Any`, `All`, `Count` и другие.
- ✅ **Ленивые итераторы** — данные обрабатываются по мере необходимости.
- ✅ **Дженерики (Go 1.18+)** — типобезопасность без приведения типов.
- ✅ **Цепочки операций** — красивый fluent-интерфейс.
- ✅ **Эффективность** — минимум аллокаций, подходит для больших данных.
- ✅ **Безопасная работа с nil** — методы корректно обрабатывают `nil` последовательности, не вызывая паник.

---

## 🚀 Установка

Используйте `go get`, чтобы добавить библиотеку в ваш проект:

```bash
go get github.com/ahatornn/queryable
```

## ⚖️ Сравнение: ванильный Go vs Queryable

Представим задачу: объединить два слайса, оставить числа больше 50, пропустить первые 3, взять следующие 2.

### ✅ Используя библиотеку queryable

```go
nums1 := []int{10, 20, 60, 70}
nums2 := []int{30, 40, 80, 90, 100}

result := query.ToQueryable(nums1).
    Union(query.ToQueryable(nums2)).
    Where(func(x int) bool { return x > 50 }).
    Skip(3).
    Take(2).
    ToSlice()

fmt.Println(result) // Вывод: [90 100]
```

### 🛠 Без библиотеки (ручной способ)

```go
nums1 := []int{10, 20, 60, 70}
nums2 := []int{30, 40, 80, 90, 100}

seen := make(map[int]bool)
var union []int

for _, x := range nums1 {
    if !seen[x] {
        seen[x] = true
        union = append(union, x)
    }
}
for _, x := range nums2 {
    if !seen[x] {
        seen[x] = true
        union = append(union, x)
    }
}

var filtered []int
for _, x := range union {
    if x > 50 {
        filtered = append(filtered, x)
    }
}

var result []int
countSkipped := 0
countTaken := 0
for _, x := range filtered {
    if countSkipped < 3 {
        countSkipped++
        continue
    }
    if countTaken < 2 {
        result = append(result, x)
        countTaken++
    } else {
        break
    }
}
```