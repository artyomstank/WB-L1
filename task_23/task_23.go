package main

import "fmt"

func delSliceElem[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) { //валидация
		return s
	}
	temp := make([]T, len(s)-1)
	copy(temp, s[:i])
	copy(temp[i:], s[i+1:])
	return temp
}

func main() {
	slice1 := []int{1, 2, 45, 4, 5}
	slice2 := []string{"d", "q", "t", "r", "w"}
	fmt.Println("Исходный:", slice1)
	fmt.Println("Исходный:", slice2)

	cutSlice1 := delSliceElem(slice1, 3)
	fmt.Println("После удаления:", cutSlice1)
	cutSlice2 := delSliceElem(slice2, 1)
	fmt.Println("После удаления:", cutSlice2)
}
