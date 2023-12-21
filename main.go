package main

import "fmt"

func main() {
	p := NewPerson(WithAge(10), WithGender("Male"), WithName("Abdullah"))
	fmt.Printf("person name: %s | person gender: %s | person age: %d\n", p.Name, p.Gender, p.Age)
}

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

func WithName(name string) Option {
	return func(person *Person) {
		person.Name = name
	}
}

func WithAge(age int) Option {
	return func(person *Person) {
		person.Age = age
	}
}

func WithGender(gender string) Option {
	return func(person *Person) {
		person.Gender = gender
	}
}
