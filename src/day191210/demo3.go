package main

import "fmt"

type Dog struct {
	name string
}

func (dog *Dog) setName(name string) {
	dog.name = name
}

func New(name string) Dog {
	return Dog{
		name: name,
	}
}

func main()  {
	d := New("dog1")
	d.setName("monster")
	fmt.Println(d.name)
}
