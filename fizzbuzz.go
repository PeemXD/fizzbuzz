package main

import "fmt"

func fizzBuzz(i int) string {
	if i%3 == 0 || i%5 == 0 {
		return isFizzBuzz(i)
	}

	return fmt.Sprintf("%v", i)
}

func isFizzBuzz(i int) string {
	m := map[bool]string{
		true:  "FizzBuzz",
		false: isFizz(i),
	}

	return m[i%3 == 0 && i%5 == 0]
}

func isFizz(i int) string {
	m := map[bool]string{
		true:  "Fizz",
		false: "Buzz",
	}

	return m[i%3 == 0]
}
