package main

import (
	"fmt"
)

func intersec(a, b []int) []int {
	if len(a) > len(b) {
		a, b = b, a
	}
	mins := len(a)
	result := make([]int, 0, mins)
	m := make(map[int]struct{}, mins)
	for _, x := range a {
		m[x] = struct{}{}
	}
	for _, x := range b {
		if _, ok := m[x]; ok {
			result = append(result, x)
			delete(m, x)
		}
	}
	return result
}

func main() {
	A, B := []int{1, 2, 3, 5, 7, 7, 21, 33}, []int{2, 3, 4, 7, 9, 21, 23}
	C := intersec(A, B)

	fmt.Printf("Слайс A: %v\n", A)
	fmt.Printf("Слайс B: %v\n", B)
	fmt.Printf("Пересечение: %v\n", C) //вывод: Пересечение: [2 3 7 21]

}
