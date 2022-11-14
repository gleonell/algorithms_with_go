package main

import "fmt"

func Reverse(word string) string {
	var reversed string

	for _, r := range word {
		reversed = string(r) + reversed
	}
	return reversed
}

func main() {
	//word := "люси"
	word := "日本語"
	fmt.Println(Reverse(word))
}