package main

import (
	"fmt"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	shouldReturn := func(want string, give int) string {
		return fmt.Sprintf("should return %v when given %v", want, give)
	}

	tests := []struct {
		give int
		want string
	}{
		{
			give: 1,
			want: "1",
		},
		{
			give: 2,
			want: "2",
		},
		{
			give: 3,
			want: "Fizz",
		},
	}

	for _, tt := range tests {
		t.Run(shouldReturn(tt.want, tt.give), func(t *testing.T) {
			got := fizzBuzz(tt.give)

			if tt.want != got {
				t.Errorf("want %v but got %v", tt.want, got)
			}
		})
	}
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
