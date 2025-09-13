package main

import "fmt"

func main() {
	s := "сопм"
	runeString := []rune(s)
	for f, l := 0, len(runeString)-1; f < l; f, l = f+1, l-1 {
		runeString[f], runeString[l] = runeString[l], runeString[f]
	}
	fmt.Println(string(runeString))
}
