package main

import(
	"fmt"
)

func ifelse(){

	// 1.normal if 
	num := 10
	if num % 2 == 1{
		fmt.Println("Num is Odd")
	}else{
		fmt.Println("Num is Even")
	}

	// 2.with && and ||
	age := 10
	if age >= 18 && age <= 60 {
    	fmt.Println("Eligible for work")
	} else {
		fmt.Println("Not in the working age bracket")
	}
	// its single if we can use this also
	if 8 % 2 == 0 || 5 % 2 == 0{
		fmt.Println("either 8 or 5 are even")
	}


	// 3. with var initialization and use in if 
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

}