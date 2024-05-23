package main

import "strconv"

func fizzBuzz(i int) string {
	if i == 3 {
		return "Fizz"
	}

	return strconv.Itoa(i)
}
