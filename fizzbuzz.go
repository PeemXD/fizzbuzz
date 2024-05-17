package main

import "fmt"

func fizzBuzz(num int) string {
	if num == 3 {
		return "Fizz"
	}

	return fmt.Sprintf("%v", num)
}
