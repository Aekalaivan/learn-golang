package main

import (
	"fmt"
)

const h int = 10 //This is allowed, Its called global var whose scope is within this file and files in the same package in this case all files inside the package 'main'.
// a:= 10 Won't work in global scope, because auto type inference is won't work in global scope in go lang.

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

	const g uint = 100 //This is a constant, Constant's value must be know at compile time.
	fmt.Println(g)

	fmt.Printf("This is the const 'h' wiht value %v from gloable scope of this file.\n", h)
}
