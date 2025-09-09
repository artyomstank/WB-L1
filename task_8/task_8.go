package main

import "fmt"

func setBit(n int64, i uint, value bool) int64 {
	if i >= 64 {
		panic("bit index out of range")
	}
	mask := int64(1 << i)
	if value {
		return n | mask // установить в 1
	}
	return n &^ mask // установить в 0
}

func assertEqual(got, want int64, msg string) {
	if got != want {
		fmt.Printf("ошибка: %s: получено %d (%b), ожидаем %d (%b)\n", msg, got, got, want, want)
	} else {
		fmt.Printf("успешно: %s: %d (%b)\n", msg, got, got)
	}
}

func main() {
	// Тест 1: 5 (0101) → установить 1-й бит в 0 → 4 (0100) -- ошибка он уже равен 0 в числе 5
	assertEqual(setBit(5, 1, false), 4, "установить 1-й бит числа 5 в 0 ")

	// Тест 2: 5 (0101) → установить 1-й бит в 1 → 7 (0111)
	assertEqual(setBit(5, 1, true), 7, "установить 1-й бит числа 5 в 1")

	// Тест 3: 6 (0110) → установить 1-й бит в 0 → 4 (0100)
	assertEqual(setBit(6, 1, false), 4, "установить 1-й бит числа 6 в 0")

	// Тест 4: 0 (0000) → установить 3-й бит в 1 → 8 (1000)
	assertEqual(setBit(0, 3, true), 8, "установить 3-й бит числа 0 в 1")

	// Тест 5: 15 (1111) → установить 2-й бит в 0 → 11 (1011)
	assertEqual(setBit(15, 2, false), 11, "установить 2-й бит числа 15 в 0")
}
