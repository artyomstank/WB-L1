package main

import "fmt"

func main() {
	arr := [5]string{"cat", "cat", "dog", "cat", "tree"}
	set := make([]string, 0, len(arr))
	m := make(map[string]struct{}, len(arr)) //реализуем через уникальные ключи в мапе
	for _, x := range arr {
		m[x] = struct{}{}
	}
	for i := range m {
		set = append(set, i)
	}
	fmt.Println("множество =", set)
}
