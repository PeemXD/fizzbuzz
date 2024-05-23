package main

import "strconv"

func fizzBuzz(i int) string {
	m := map[bool]string{
		true:  "Fizz",
		false: isBuzz(i),
	}

	return m[i%3 == 0]
}

func isBuzz(i int) string {
	m := map[bool]string{
		true:  "Buzz",
		false: strconv.Itoa(i),
	}

	return m[i == 5]
}
