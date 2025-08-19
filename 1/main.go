package main

import (
	"fmt"
)

type Human struct {
	name string
	age  int
}

func (h Human) SayHello() {
	fmt.Printf("Hello! My name is %s, im %d years old.\n", h.name, h.age)
}

func (h Human) GetName() string {
	return h.name
}

func (h Human) GetAge() int {
	return h.age
}

type Action struct {
	Human
	action string
}

func (a Action) DoAction() {
	fmt.Printf("%s, who is %d years old, doing %s\n", a.GetName(), a.GetAge(), a.action)
}

func main() {
	a := &Action{
		Human: Human{
			name: "Andrey",
			age:  20,
		},
		action: "wb-tech L1.1 task",
	}

	a.SayHello()

	a.DoAction()
}
