package main

import "fmt"

type animal interface {
	NumLegs() int
}

type dog struct{}

func (d dog) NumLegs() string { return "4" }

func main() {
	takeCutePhoto(dog{})
}

func takeCutePhoto(a animal) {
	fmt.Println("Taking a cute animal photo")
}
