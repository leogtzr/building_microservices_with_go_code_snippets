package main

import (
	"fmt"
)

func foo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("There was a panic ...: ", err)
		}
	}()
	panic("alrpvw")
}

func main() {
	foo()
}
