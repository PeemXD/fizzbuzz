package main

import "strconv"

func fizzBuzz(i int) string {
	if i == 3 || i == 5 || i == 6 {
		return isFizz(i)
	}
	return strconv.Itoa(i)
}

func isFizz(i int) string {
	m := map[bool]string{
		true:  "Fizz",
		false: "Buzz",
	}

	return m[i == 3 || i == 6]
}
