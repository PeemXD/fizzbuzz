package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when given 1", func(t *testing.T) {
		give := 1
		want := "1"

		got := fizzBuzz(give)

		if want != got {
			t.Errorf("want %v but got %v", want, got)
		}
	})
	t.Run("should return 2 when given 2", func(t *testing.T) {
		give := 2
		want := "2"

		got := fizzBuzz(give)

		if want != got {
			t.Errorf("want %v but got %v", want, got)
		}
	})

	t.Run("should return Fizz when given 3", func(t *testing.T) {
		give := 3
		want := "Fizz"

		got := fizzBuzz(give)

		if want != got {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}
