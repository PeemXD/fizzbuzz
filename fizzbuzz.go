package main

import "strconv"

func fizzBuzz(i int) string {
	if i%3 == 0 || i%5 == 0 {
		return isFizzBuzz(i)
	}
	return strconv.Itoa(i)
}

func isFizzBuzz(i int) string {
	m := map[bool]string{
		true:  "FizzBuzz",
		false: isFizz(i),
	}

	return m[i == 15]
}

func isFizz(i int) string {
	m := map[bool]string{
		true:  "Fizz",
		false: "Buzz",
	}

	return m[i%3 == 0]
}
