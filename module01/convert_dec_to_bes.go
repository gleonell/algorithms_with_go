package main

import (
	"fmt"
	//"strconv"
)
/*
func DecToBase(dec int, base int) string {
	var res string

	for dec > 0 {
		remainder := dec % base 
		switch remainder {
			case 10: res = "A" + res
			case 11: res = "B" + res
			case 12: res = "C" + res
			case 13: res = "D" + res
			case 14: res = "E" + res
			case 15: res = "F" + res
			default: res = strconv.Itoa(remainder) + res
		}
		dec = dec / base
	}
	return res
} 
*/

						//much more elegant solution:
/*
func DecToBase(dec int, base int) string {
	var res string

	for dec > 0 {
		rem := dec % base
		res = fmt.Sprintf("%X%s", rem, res)
		dec = dec / base
	}
	return res
}
*/

						//and version without using any package:

func DecToBase(dec int, base int) string {
	const charset = "0123456789ABCDEF"
	var res string
	for dec > 0 {
		rem := dec % base
		res = string(charset[rem]) + res
		dec /= base
	}
	return res
}


func main() {
	fmt.Println(DecToBase(3735928559, 16))// "DEADBEEF"
	fmt.Println(DecToBase(1, 2))// "1"
	fmt.Println(DecToBase(2, 2)) // "10"
	fmt.Println(DecToBase(7, 3)) // "21"
	fmt.Println(DecToBase(14, 2)) // "1110"
	fmt.Println(DecToBase(14, 16) )// "E"
	fmt.Println(DecToBase(17, 16)) // "11"
}