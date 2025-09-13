package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var l, eq, gr []int //l-less, gr- greater, eq-equal
	for _, item := range arr {
		if item < pivot {
			l = append(l, item)
		} else if item == pivot {
			eq = append(eq, item)
		} else {
			gr = append(gr, item)
		}
	}
	lessSorted := quickSort(l)
	greaterSorted := quickSort(gr)
	return append(append(lessSorted, eq...), greaterSorted...)
}

func main() {
	s1 := []int{8, 2, 8, 10, 1, 4, 3, 9, 4, 2, 2}
	fmt.Println("исходные данные слайса 1:", s1)
	sortS1 := quickSort(s1)
	fmt.Println("отсортированный слайс 1:", sortS1)

	s2 := []int{0, 1, 24, 16, 2, 13, 53, 62, 75, 34, 94}
	fmt.Println("исходные данные слайса 2:", s2)
	sortS2 := quickSort(s2)
	fmt.Println("отсортированный слайс 2:", sortS2)
}
