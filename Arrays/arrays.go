package main

import "fmt"

func modify(a [5]int) {
	a[0] = 100
}

func arrays() {
	// 1.basic
	var a [5]int
	fmt.Println(a)

	a[1] = 2
	fmt.Println(a)

	// 2.
	fmt.Println("// 2.")
	var b = [5]int{1, 2, 3, 4, 5}

	fmt.Println(b)

	// clear and reasing vals
	b = [...]int{5, 4, 3, 2, 1}
	fmt.Println("dcl:", b)

	// exceptional
	b = [...]int{100, 3: 400, 500} // for idx:3 = 400  simply o/p: idx: [100, 0, 0, 400, 500]
	fmt.Println("idx:", b)

	// 3.
	fmt.Println("// 3.")
	c := [4]int{1, 2, 3} // here c[3] = 0
	for i := range len(c) {
		fmt.Println(c[i], " ")
	}

	// 4.this is also accepted clear and reassing
	fmt.Println("// 4.")
	var nums = [...]int{1, 2, 3, 4, 5}

	for idx, num := range nums {
		fmt.Println(idx, ":", num)
	}

	// It passes a copy, not reference.
	modify(nums)
	fmt.Println(nums)

	// 5. 2-D array
	fmt.Println("// 5.")
	var mat [2][3]int // need to remmenber this if i use = means assing the val their in-line else this

	for r := range 2 {
		for c := range 3 {
			mat[r][c] = r + c
		}
	}

	fmt.Println(mat)
	mat = [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(mat)
}
