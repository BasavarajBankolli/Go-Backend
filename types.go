package main

import "fmt"

func main() {
	
	var chr rune = 'r'

	fmt.Printf("Type : %T\n", chr)
	fmt.Printf("Value: %d\n", chr)
	fmt.Printf("Letter: %c\n", chr)

	var name string = "Raj"
	fmt.Printf("Name: %s\n", name)

	var byte byte = 'c'
	fmt.Printf("Char: %c\n", byte)

	var num uint = 100
	var pie float32 = 3.14321
	val := 1000

	fmt.Printf("number %d, pie value %f and another number %d", num, pie, val)



}