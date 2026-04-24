package main

import "fmt"

func loop() {

	for i := 1; i < 9; i++ {

		fmt.Printf("%d\n", i)

	}

	fmt.Print()

	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

}
