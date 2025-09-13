package main

import "fmt"

func main() {
	var a, b int = 2, 4
	var c, d int = 5, 9
	fmt.Println("до операций a,b = ", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("после a,b = ", a, b)
	fmt.Println("до операций c,d = ", c, d)
	c = c ^ d
	d = c ^ d
	c = c ^ d

	fmt.Println("после c,d = ", c, d)
}
