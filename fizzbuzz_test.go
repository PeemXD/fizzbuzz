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
		{
			give: 4,
			want: "4",
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
}
