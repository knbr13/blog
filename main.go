package main

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Option func(*Person)

func NewPerson(options ...Option) *Person {
	person := &Person{}
	for _, option := range options {
		option(person)
	}
	return person
}
