package main

import "fmt"

func SumListNums(list []int) int{
	sum := 0
	
	if list != nil {
		
		for _, val := range list {
			sum += val
		}
	}
	return sum
}

func main() {
	fmt.Println(15, SumListNums([]int{1, 2, 3, 4, 5}))
	fmt.Println(6, SumListNums([]int{3, 3}))
	fmt.Println(19, SumListNums([]int{3, 5, 3, 5, 3}))
	fmt.Println(26, SumListNums([]int{4, 2, 22, -10, 8}))
	fmt.Println(-15, SumListNums([]int{-1, -2, -3, -4, -5}))
	fmt.Println("", SumListNums([]int{}))
	fmt.Println("nil", SumListNums(nil))
}