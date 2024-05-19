package main

import (
	"strconv"
)

func fizzBuzz(i int) string {
	return isFizz(i)
}

func isFizz(i int) string {
	m := map[bool]string{
		true:  "Fizz",
		false: strconv.Itoa(i),
	}

	return m[i%3 == 0]
}
