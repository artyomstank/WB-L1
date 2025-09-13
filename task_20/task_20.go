package main

import (
	"fmt"
)

func main() {
	s := "мопсы это собаки"
	runeString := []rune(s)
	for f, l := 0, len(runeString)-1; f < l; f, l = f+1, l-1 {
		runeString[f], runeString[l] = runeString[l], runeString[f]
	}

	startIdx := 0
	for i := 0; i <= len(runeString); i++ {
		if i == len(runeString) || runeString[i] == ' ' {
			for f, l := startIdx, i-1; f < l; f, l = f+1, l-1 {
				runeString[f], runeString[l] = runeString[l], runeString[f]
			}
			startIdx = i + 1
		}
	}
	fmt.Println(string(runeString))
}
