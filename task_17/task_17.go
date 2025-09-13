package main

import "fmt"

func binser(list []int, target int) (int, bool) {
	low_idx := 0
	high_idx := len(list) - 1
	for low_idx <= high_idx {
		mid := (low_idx + high_idx) / 2
		guess := list[mid]

		if guess == target {
			return mid, true
		} else if guess > target {
			high_idx = mid - 1
		} else {
			low_idx = mid + 1
		}
	}
	return 0, false
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	searchVal := []int{1, 5, 10, 0, 11}

	for _, t := range searchVal {
		idx, ok := binser(arr, t)
		if ok {
			fmt.Printf("Число %d найдено на индексе %d\n", t, idx)
		} else {
			fmt.Printf("Число %d не найдено\n", t)
		}
	}
}
