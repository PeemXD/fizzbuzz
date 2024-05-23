package main

import "strconv"

func fizzBuzz(i int) string {
	if i == 15 {
		return "FizzBuzz"
	}

	return isFizz(i)
}

func isFizz(i int) string {
	if i%3 == 0 {
		return "Fizz"
	}

	return isBuzz(i)
}

func isBuzz(i int) string {
	if i%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(i)
}
