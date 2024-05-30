package main

import "strconv"

type fbI interface {
	isFizzBuzz(int) string
}
type fb struct {
	s int
}

func newFb() fb {
	return fb{s: 1}
}

func (f *fb) isFizzBuzz() string {
	return strconv.Itoa(f.s)
}

func fizzBuzz(n int) string {
	fb := newFb()
	return fb.isFizzBuzz()
}
