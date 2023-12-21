package main

type Person struct {
	Name   string
	Age    int
	Gender string
}

type Option func(*Person)
