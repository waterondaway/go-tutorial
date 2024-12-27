package main

import (
	"fmt"
)

// Comment single-line
/*
Comment multi-line
*/

func main() {
	// var name string = "Nonpawit"             // type is string
	// var school = "Bangkok Christian College" // type is inferred
	// age := 25                                // type is inferred
	fmt.Println("Hello World!")

	/* Every variable in GO language need to have initial value
	If not value will be set to the default values of its type*/
	var a string
	var b int
	var c bool
	fmt.Println(a) // values is [empty string]
	fmt.Println(b) // values is 0
	fmt.Println(c) // values is false

	// assign value after declare variable for like. var a string
	a = "Good Morning"
	fmt.Println(a)

	/* var vs := ~ var is global variable that mean can use a both of inside-outside function and don't need
	to assign value at first | := only use inside function and need to assign value at first */
	// var hey string | OK
	// name :=  | ERROR

	// Multiple Variable Declaration
	var x, y, z, e int = 1, 3, 5, 7 // type is int
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(e)

	h, q := 7, "World!" // type is inferred
	fmt.Println(h)
	fmt.Println(q)

	var ( // Variable Declaration in a Block
		o int
		p int    = 1
		j string = "Ayo!"
	)
	fmt.Println(o)
	fmt.Println(p)
	fmt.Println(j)

	// Camel Case
	// Each word, except the first, starts with a capital letter:

	// myVariableName = "John"
	// Pascal Case
	// Each word starts with a capital letter:

	// MyVariableName = "John"
	// Snake Case
	// Each word is separated by an underscore character:

	// my_variable_name = "John"

	// Constants Varible that mean cannot be changed
	const PI float32 = 3.14
	// Typed Constants
	const NAME string = "Nonpawit"
	// Untyped Constants
	const NICKNAME = "Nam" // type is inferred same as above

	// Multiple Constants Declaration
	const (
		A int = 1
		B int = 2
		C int = 3
	)
	fmt.Println(PI, NAME, NICKNAME, A, B, C)

	// The Print() Function
	fmt.Print("Nonpawit")
	fmt.Print(" Silabumrungrad\n")

	// The Println() Function when the output are display will set new lines
	fmt.Println("Kasetsart University")

	// The Printf()
	var J string = "Hello I'm h variable"
	var w int = 10
	fmt.Printf("h has value: %v and type %T\n", J, J)
	fmt.Printf("h has value: %v and type %T\n", w, w)
}
