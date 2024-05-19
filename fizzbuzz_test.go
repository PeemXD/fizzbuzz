package main

import (
	"fmt"
	"testing"
)

func shouldReturn(want string, given int) string {
	return fmt.Sprintf("should return %v when given %v", want, given)
}

func TestFizzBuzz(t *testing.T) {

	tests := []struct {
		given int
		want  string
	}{
		{
			given: 1,
			want:  "1",
		},
		{
			given: 2,
			want:  "2",
		},
		{
			given: 3,
			want:  "Fizz",
		},
		{
			given: 4,
			want:  "4",
		},
		{
			given: 5,
			want:  "Buzz",
		},
		{
			given: 6,
			want:  "Fizz",
		},
		{
			given: 7,
			want:  "7",
		},
		{
			given: 9,
			want:  "Fizz",
		},
		{
			given: 10,
			want:  "Buzz",
		},
		{
			given: 12,
			want:  "Fizz",
		},
		{
			given: 15,
			want:  "FizzBuzz",
		},
	}

	for _, tt := range tests {
		t.Run(shouldReturn(tt.want, tt.given), func(t *testing.T) {
			got := fizzBuzz(tt.given)

			if tt.want != got {
				t.Errorf("want %v but got %v\n", tt.want, got)
			}
		})
	}
}
