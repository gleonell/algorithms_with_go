package main

import "fmt"

func NumOnList(list []int, num int) bool {
	if list != nil {
		checked := map[int]bool{}
		
		for _, val := range list {
			checked[val] = true
		}

		for key := range checked {
			if key == num {
				return true
			}
		}
	}
	return false
}

func main() {
	fmt.Println("TRUE - " , NumOnList([]int{1, 3, 4, 6}, 1))
	fmt.Println("TRUE - " , NumOnList([]int{1, 3, 4, 6}, 6))
	fmt.Println("TRUE - " , NumOnList([]int{1, 4, 3, 6}, 3))
	fmt.Println("TRUE - " , NumOnList([]int{3, 3, 3, 3}, 3))
	fmt.Println("TRUE - " , NumOnList([]int{1, 3, 4, -6}, 1))
	fmt.Println("TRUE - " , NumOnList([]int{1, -3, 4, -6}, -6))
	fmt.Println("TRUE - " , NumOnList([]int{1}, 1))
	fmt.Println("TRUE - " , NumOnList([]int{-1,-3, -4, -6}, -1))
	fmt.Println()
	fmt.Println("FALSE - " , NumOnList([]int{1, 3, 4, 6}, 7))
	fmt.Println("FALSE - " , NumOnList([]int{1, 3, 4, 6}, -7))
	fmt.Println("FALSE - " , NumOnList([]int{1, 3, 4, 6}, -7))
	fmt.Println("FALSE - " , NumOnList([]int{}, -7))
	fmt.Println("FALSE - " , NumOnList([]int{1, 3, 4, 6}, 0))
	fmt.Println("FALSE - " , NumOnList([]int{1, 3, 4, 6}, -7))
}