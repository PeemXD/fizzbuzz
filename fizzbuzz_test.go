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
}
