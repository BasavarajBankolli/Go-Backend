package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello") // print normal print func 
	fmt.Println()

	name := "Basavaraj"
	age := 25
	fmt.Printf("%s's age is %d", name, age)	// printf used to format the o/p
	println()  

	fmt.Println("hey")  // println used to print in new line 


	var chr rune = 'r'

	fmt.Printf("Type : %T\n", chr)
	fmt.Printf("Value: %d\n", chr)
	fmt.Printf("Letter: %c\n", chr)

	var byte byte = 'c'
	fmt.Printf("Char: %c\n", byte)

	var num uint = 100
	var pie float32 = 3.14321
	val := 1000

	fmt.Printf("number %d, pie value %f and another number %d", num, pie, val)




}	