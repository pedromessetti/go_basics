package main

import "fmt"

func main() {
	
	num := 42
	var ptr *int = &num // ptr is a pointer to the memory address of num
	
	fmt.Printf("num: %d\n", num)
	fmt.Printf("ptr: %p\n", ptr)
	fmt.Printf("*ptr: %d\n", *ptr) // *ptr is the value of the memory address that ptr points to

	fmt.Printf("\n-=-=-=-=\tType Casting\t-=-=-=-=\n\n")
	
	// var f float32 = num // Compilation error: cannot use num (type int) as type float32 in assignment
	var f float32 = float32(num) // Type conversion of int to float32
	fmt.Printf("int to float conversion\nf: %.2f\n", f)

	var pi float32 = 3.141592
	intPi := int(pi) // Type conversion of float32 to int
	fmt.Printf("\npi: %f\n\nfloat to int conversion\nintPi: %d\n", pi, intPi)

	fmt.Printf("\n-=-=-=-=\tType Aliases\t-=-=-=-=\n\n")
	type Celsius float32
	type Fahrenheit float32

	var c Celsius = 100
	var fh Fahrenheit = 212

	fmt.Printf("c: %.2f\n", c)

	// c = fh // Compilation error: cannot use fh (type Fahrenheit) as type Celsius in assignment
	c = Celsius(fh) // Type conversion of Fahrenheit to Celsius
	fmt.Printf("Celsius to Fahrenheit conversion\nc: %.2f\n", c)
}
