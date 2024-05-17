package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when given 1", func(t *testing.T) {
		num := 1
		want := "1"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})

	t.Run("should return 2 when given 2", func(t *testing.T) {
		num := 2
		want := "2"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})
	t.Run("should return Fizz when given 3", func(t *testing.T) {
		num := 3
		want := "Fizz"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})

	t.Run("should return 4 when given 4", func(t *testing.T) {
		num := 4
		want := "4"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})

	t.Run("should return Buzz when given 5", func(t *testing.T) {
		num := 5
		want := "Buzz"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})

	t.Run("should return Fizz when given 6", func(t *testing.T) {
		num := 5
		want := "Buzz"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})

	t.Run("should return 7 when given 7", func(t *testing.T) {
		num := 7
		want := "7"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})

	t.Run("should return 8 when given 8", func(t *testing.T) {
		num := 8
		want := "8"

		got := fizzBuzz(num)

		if want != got {
			t.Errorf("want %v but got %v\n", want, got)
		}
	})
}
