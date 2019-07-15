package main

import "fmt"

type I2 interface {
	m1()
	m2()
}

type s1 struct {
	v1 int
	v2 int
}

func NewS1() *s1 {
	return &s1{
		v1:1,
		v2:2,
	}
}

func (s *s1) m1() {
	fmt.Println("m1 called by s1")
}

func (s *s1) m2() {
	fmt.Println("m2 called by s2")
}

func main() {
	s1 := NewS1()
	fmt.Println("v1:", s1.v1)
	fmt.Println("v2:", s1.v2)

	s1.m1()
	s1.m2()

}



