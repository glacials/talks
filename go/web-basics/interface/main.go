package main

import (
	"fmt"

	"github.com/glacials/talks/go/web-basics/interface/user"
)

// main.go
func main() {
	u := user.New("Ben", "Carlsson")
	fmt.Printf("Registered %s", u.FullName())
}
