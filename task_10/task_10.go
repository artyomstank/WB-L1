package main

import (
	"fmt"
)

func main() {
	tempSlice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := make(map[int][]float64)
	for _, x := range tempSlice {
		groupKey := int(x/10) * 10
		groups[groupKey] = append(groups[groupKey], x)
	}
	fmt.Println(groups)
}
