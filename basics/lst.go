package main

import (
    "fmt"
)

func lst() {
	nums := []int{1, 2, 3, 4}

	for _, num := range nums {
		fmt.Print(num, " ")
	}
}
