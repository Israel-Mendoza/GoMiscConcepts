package main

import "fmt"

const a = 10

func main() {
	// Printing value and type of a
	fmt.Printf("a = %v (%T)\n", a, a)

	// Declaring a byte local variable
	var myByte byte = 128
	// Printing value and type of a
	fmt.Printf("myByte = %v (%T)\n", myByte, myByte)

	/* Adding a + myByte. Operation succeeds because the
	constant is replaced by a literal value at compile time */
	fmt.Printf("a + myByte = %v (%T)", a+myByte, a+myByte)
}
