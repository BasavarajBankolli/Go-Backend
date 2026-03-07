package main

import "fmt"

func main() {

	lst := []int{1, 2, 2, 3}

    for _, num := range lst {
        fmt.Println(num)
    }
}
