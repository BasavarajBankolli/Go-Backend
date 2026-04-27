package main

import (
	"fmt"
	"slices"
)

func slice() {
	var s []string
	fmt.Println("uninitialized :", s, s == nil, len(s) == 0)
	
	// 2. use make to dcl sz  
	fmt.Println("// 2.") 
	s = make([]string, 3)
	s[0], s[1],s[2] = "a", "b", "c"
	fmt.Println("after updating s:", s)

	
	// 3. copy method use 
	fmt.Println("// 3.")
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy of s: ", c)

	l, r := s[:2], s[1:]
	fmt.Println("Sub str l:", l, "Sub str r:", r)
	
	// 4.dcl -> declare	
	fmt.Println("// 4.")

	t := []string {"a","b","c"}
	fmt.Println("Strign t:", t)

	t = []string {"d", "c", "b", "a"}
	fmt.Println("after modification of t:", t)

	// 5. slice.equal method 
	fmt.Println("// 5.")
	t1 := []string {"d", "c", "b", "a"}
	fmt.Println("Is t1 and t are equal:", slices.Equal(t, t1))
	
	t2 := []string {"d", "c", "b"}
	fmt.Println("Is t2 and t are equal:", slices.Equal(t, t2))

	// 6. 2-D
	fmt.Println("// 6.")

	twoD := make([][]int, 3)
    for i := range 3 {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := range innerLen {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
	rows, cols := 3, 4
	
	data := make([]int, rows*cols)
	twoD = make([][]int, rows)
	for i := 0; i < rows; i++ {
		twoD[i] = data[i*cols : (i+1)*cols]
	}
}