package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should return 1 when given 1", func(t *testing.T) {
		given := 1
		want := "1"

		got := fizzBuzz(given)

		if want != got {
			t.Errorf("want %v but got %v", want, got)
		}
	})
}
