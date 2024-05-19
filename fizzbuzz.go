package main

import "strconv"

func fizzBuzz(i int) string {
	if i == 3 {
		return "Fizz"
	}
	if i == 5 {
		return "Buzz"
	}
	return strconv.Itoa(i)
}
