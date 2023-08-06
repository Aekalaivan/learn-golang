package main

import (
	"fmt"
)

func main() {
	a := 10
	var b = 10
	var c int = 10
	fmt.Println(a, b, c)

	a, b, c = 1, 2, 3 // Assigning values to multiple variables in single line
	fmt.Println(a, b, c)

	var defaultValue uint64
	fmt.Printf("Default value of %T -> %v\n", defaultValue, defaultValue) // Prints zero, Since default value of Integer is zero.

	d := -a
	d += d - a
	fmt.Println(d)

	e := &a // Address of the var 'a' is assigned to var 'e'
	f := *e // The value at the memory address held by var 'e' is assigned to var 'f'
	fmt.Println(e, *e, f)
}
