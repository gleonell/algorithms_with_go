//Given a number N, print out all the numbers from 1 to N but when a number is 
//divisible by3 print out Fizz instead og the number, and when number is
//divisible by 5 print out Buzz instead of the number. For numbers divisible 
//by both 3 and 5 print Fizz Buzz

package main 

import "fmt"

func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Print("Fizz Buzz")
		} else if i % 3 == 0 {
			fmt.Print("Fizz")
		} else if i % 5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		if i != n {
			fmt.Print(", ")
		}	
	}
	fmt.Println()
}

func main() {
	FizzBuzz(1)
	FizzBuzz(5)
	FizzBuzz(15)
}