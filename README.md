# queryable: LINQ-style Iterators for Go
[![Тесты](https://github.com/ahatornn/queryable/actions/workflows/test.yml/badge.svg)](https://github.com/ahatornn/queryable/actions/workflows/test.yml)

`queryable` is a Go library inspired by **C# LINQ** (Language Integrated Query) that enables functional-style data processing using **generics** and **lazy iterators**.

With it, you can write clean, readable, and expressive code for filtering, transforming, and aggregating data — without `for` loops or temporary slices.

> Feels like LINQ, but in Go. Lazy evaluation. Generic. Type-safe.

---

## ✨ Features

- ✅ **LINQ-like methods**: `Where`, `Select`, `Take`, `Skip`, `First`, `Any`, `All`, `Count`, and more.
- ✅ **Lazy iterators** — data is processed on-demand.
- ✅ **Generics (Go 1.18+)** — type safety without type casting.
- ✅ **Method chaining** — fluent, readable APIs.
- ✅ **Efficiency** — minimal allocations, suitable for large datasets.
- ✅ **Nil-safe operations** — methods safely handle `nil` sequences without panics.

---

## 🚀 Installation

Use `go get` to add the library to your project:

```bash
go get github.com/ahatornn/queryable
```

## ⚖️ Comparison: Vanilla Go vs Queryable

Let’s say you want to: Merge two slices, keep numbers greater than 50, skip the first 3, and take the next 2.

### ✅ Using queryable

```go
nums1 := []int{10, 20, 60, 70}
nums2 := []int{30, 40, 80, 90, 100}

result := query.ToQueryable(nums1).
    Union(query.ToQueryable(nums2)).
    Where(func(x int) bool { return x > 50 }).
    Skip(3).
    Take(2).
    ToSlice()

fmt.Println(result) // result: [90 100]
```

### 🛠 Without the library (manual)

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
