package main

import (
    "fmt"
    "maps"
)

func mps() {

    m := make(map[string]int)

    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)

    
    fmt.Println("v1:", m["k1"])
    fmt.Println("v3:", m["k2"])

    fmt.Println("len:", len(m))

    delete(m, "k2") // delete remove the last val of the map
    fmt.Println("map:", m)
 
    clear(m)
    fmt.Println("map:", m)

    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	// 2. check they r equal or not
	fmt.Println("// 2.")
    n2 := map[string]int{"foo": 1, "bar": 2}
    if maps.Equal(n, n2) {
        fmt.Println("n == n2")
    }

	// 3. make copy of the map
	n3 := make(map[string]int)
	maps.Copy(n3, n2) // copy(dst, src)

	fmt.Println(n3)

	mp := map[string]int{"foo": 1, "bar": 2}

	for k, v := range mp {
		fmt.Println(k, v)
	}

}
