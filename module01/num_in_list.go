package main

import "fmt"

func NumInList(list []int, num int) bool {
	if list != nil {
		for _, key := range list {
			if key == num {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("TRUE - " , NumInList([]int{1, 3, 4, 6}, 1))
	fmt.Println("TRUE - " , NumInList([]int{1, 3, 4, 6}, 6))
	fmt.Println("TRUE - " , NumInList([]int{1, 4, 3, 6}, 3))
	fmt.Println("TRUE - " , NumInList([]int{3, 3, 3, 3}, 3))
	fmt.Println("TRUE - " , NumInList([]int{1, 3, 4, -6}, 1))
	fmt.Println("TRUE - " , NumInList([]int{1, -3, 4, -6}, -6))
	fmt.Println("TRUE - " , NumInList([]int{1}, 1))
	fmt.Println("TRUE - " , NumInList([]int{-1,-3, -4, -6}, -1))
	fmt.Println()
	fmt.Println("FALSE - " , NumInList([]int{1, 3, 4, 6}, 7))
	fmt.Println("FALSE - " , NumInList([]int{1, 3, 4, 6}, -7))
	fmt.Println("FALSE - " , NumInList([]int{1, 3, 4, 6}, -7))
	fmt.Println("FALSE - " , NumInList([]int{}, -7))
	fmt.Println("FALSE - " , NumInList([]int{1, 3, 4, 6}, 0))
	fmt.Println("FALSE - " , NumInList([]int{1, 3, 4, 6}, -7))
}