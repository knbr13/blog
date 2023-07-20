package main

type Speed uint8

const (
	Slow Speed = iota
	Medium
	Fast
)

type Arguments struct {
	Speed Speed `default:"Medium"`
}

func main() {

}
