package main

import "strconv"

type fbI interface {
	isFizzBuzz(int) string
}
type fb struct {
	s int
}

func newFb(n int) fb {
	return fb{s: n}
}

func (f *fb) isFizzBuzz() string {
	return strconv.Itoa(f.s)
}

func fizzBuzz(n int) string {
	fb := newFb(n)
	return fb.isFizzBuzz()
}
