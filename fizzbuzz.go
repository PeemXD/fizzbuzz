package main

import "fmt"

func fizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	}
	if num == 5 {
		return "Buzz"
	}

	return fmt.Sprintf("%v", num)
}
