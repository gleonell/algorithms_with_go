package main

import (
	"fmt"
	"strconv"
)

func BaseToDec(value string, base int) int {
	rem := 0
	for i := range value {
		next := 0
		switch value[i] {
		case 'A': next = 10
		case 'B': next = 11
		case 'C': next = 12
		case 'D': next = 13
		case 'E': next = 14
		case 'F': next = 15
		default: next, _ = strconv.Atoi(string((value[i])))
		}
		rem = (rem * base) + next
	}
	return rem
}

func main() {
	fmt.Println(BaseToDec("E", 16))	
}