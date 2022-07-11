package main

import (
		"fmt"
		"unicode/utf8"
	)



func Reverse(word string) string {
	reversed := []rune{}
	for i := utf8.RuneCountInString(word) - 1; i >= 0; i-- {
		reversed = append(reversed, rune(word[i]))
	}
	return string(reversed)
}

func main() {
	//word := "люси"
	word := "日本語"
	fmt.Println(Reverse(word))
}