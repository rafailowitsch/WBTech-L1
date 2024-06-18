package main

import "fmt"

type Gender int

const (
	Male Gender = iota
	Female
)

func (g Gender) String() string {
	return [2]string{"Male", "Female"}[g]
}

type Human struct {
	Name   string
	Age    int
	Gender Gender
}

func (h Human) String() string {
	return fmt.Sprintf("%s (%s) is %d years old", h.Name, h.Gender.String(), h.Age)
}

type Action struct {
	Actor Human
}

func (a Action) Act() {
	fmt.Println(a.Actor.String())
}

func taskN1() {
	human := Human{
		"John",
		25,
		0,
	}

	fmt.Println("Human struct: ", human.String())

	action := Action{
		human,
	}

	fmt.Print("Action struct: ")
	action.Act()
}
