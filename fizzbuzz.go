package main

import "fmt"

func fizzBuzz(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}
	return fmt.Sprintf("%v", i)
}
