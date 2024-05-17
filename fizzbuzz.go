package main

import "fmt"

func fizzBuzz(num int) string {
	if num%3 == 0 {
		return "Fizz"
	}
	if num%5 == 0 {
		return "Buzz"
	}

	return fmt.Sprintf("%v", num)
}
