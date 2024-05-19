package main

import "strconv"

func fizzBuzz(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}
	if i%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}
