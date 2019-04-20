package main

import "fmt"

// Reference and discard any package fields to avoid complaints from the compiler
// if the package is not being used.
var _ = fmt.Println

func main() {
}
