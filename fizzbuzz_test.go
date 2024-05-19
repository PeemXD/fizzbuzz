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
