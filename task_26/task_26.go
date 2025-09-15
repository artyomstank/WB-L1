package main

import "fmt"

func uniqueOrNot(s string) bool {
	unq := make(map[rune]struct{})
	for _, i := range s {
		if _, ok := unq[i]; ok {
			return false
		}
		unq[i] = struct{}{}
	}
	return true
}
func main() {
	fmt.Printf("'abcd' -> %t\n", uniqueOrNot("abcd"))
	fmt.Printf("'abCdefAaf' -> %t\n", uniqueOrNot("abCdefAaf"))
	fmt.Printf("'aabcd' -> %t\n", uniqueOrNot("aabcd"))
}
