package main

import "fmt"

func loop() {

	// 1.normal for loop 
	for i := 1; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2.while loop
	i := 0
	for i <= 5 {
		fmt.Print(i, " ")
		i += 1
	}
	fmt.Println()

	// 3.Range loop
	for i := range 5{
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 4.while True loop
	for {
		fmt.Print("True While loop")
		break
	}
	fmt.Println()

	// 5.Range with condition
	for i := range 10 {
		if i % 2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 6.Range to access nums 
	nums := []int{2,3,4,5,6,7}
	for i := 1; i < len(nums); i++ {
		print(nums[i], " ")
	}
	println()

	// Range based vals to acces the slice values 
	for _, num := range nums { // we can use nums[2:] here
		print(num, " ")
	}

}
