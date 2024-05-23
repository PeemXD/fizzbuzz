package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name  string
		given int
		want  string
	}{
		{
			name:  "should return 1 when given 1",
			given: 1,
			want:  "1",
		},
		{
			name:  "should return 2 when given 2",
			given: 2,
			want:  "2",
		},
		{
			name:  "should return Fizz when given 3",
			given: 3,
			want:  "Fizz",
		},
		{
			name:  "should return 4 when given 4",
			given: 4,
			want:  "4",
		},
		{
			name:  "should return Buzz when given 5",
			given: 5,
			want:  "Buzz",
		},
		{
			name:  "should return Fizz when given 6",
			given: 6,
			want:  "Fizz",
		},
		{
			name:  "should return 7 when given 7",
			given: 7,
			want:  "7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fizzBuzz(tt.given)

			if tt.want != got {
				t.Errorf("want %v but got %v", tt.want, got)
			}
		})
	}
}
