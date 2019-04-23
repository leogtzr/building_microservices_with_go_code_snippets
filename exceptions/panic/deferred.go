package main

import (
	"fmt"
)

func foo() {
	defer fmt.Println("World")
	defer fmt.Println("Hello")
	panic("vale dick la vida")
}

func main() {
	foo()
}
