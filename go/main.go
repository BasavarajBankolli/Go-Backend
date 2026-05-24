package main

import (
	"fmt"
)


func minMax(nums []int) (mn, mx int) {
		
	mn, mx = nums[0], nums[0]
	for _, num := range nums{
		mn = max(mn, num)
		mx = min(mx, num)
	}
	return
}

func counter() func() int {
	cnt := 0
	return func() int{
		cnt++
		return cnt
	}
}


func main(){
	/* var pro string = "Basavaraj"
	cnt := 4
	fmt.Println(pro, " ",cnt)
	const (
		pend = iota
		run
		wait
		done 
		fail
	)

	fmt.Println(fail)
	fmt.Println(done)
	fmt.Println(pend)

	const users int = 20
	num := 2
	num += 2

	fmt.Println(users, "Users")

	nums := []int{1,2,3,4}
	fmt.Println(nums)
	nums = append(nums, 5)

	subSlice := nums[:2]
	fmt.Println(subSlice)
	
	s := make([]int, 11)
	fmt.Println(s)
	*/

	/*
	mp := make(map[int]int)

	mp[1] += 1
	key, ok := mp[1]
	if ok{
		fmt.Println(key)
	}
	mp[1] += 1
	fmt.Println(mp[1])
	fmt.Println()

	//c-for
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	// while
	i := 0 
	for i <= 5 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println()

	// range over slice
	sl := []int{1,2,3,4,5}
	for i, v := range sl{
		fmt.Println(i, ":", v)
	}
	fmt.Println()

	// range over map
	mp1 := map[int]int{1:2, 2:3}
	for k, v := range mp1{
		fmt.Println(k, ":", v)
	}
	fmt.Println()
	
	i = 0
	for {
		if i == 5{
			break
		}
		fmt.Println(i)
		i += 1
	}
	*/

	// minmax := func(nums []int) (mn, mx int) {
		
	// 	mn, mx = nums[0], nums[0]
	// 	for _, num := range nums{
	// 		mn = min(mn, num)
	// 		mx = max(mx, num)
	// 	}
	// 	return
	// }

	nums := []int{1,0,2,3,4,9}
	mn, mx := minMax(nums)

	fmt.Println(mn, mx)

	// 
	res := func(nums ...int) int{
		tot := 0
		for _, val := range nums {
			tot += val
		}

		return tot
	}

	sum := res(1,2,3,4,5)
	fmt.Println(sum)

	// closure func 
	
	fmt.Println(counter()())
	fmt.Println(counter()()) 

	addone := func(x int) int{
		return x+1
	}
	res1 := pt(addone,100)
	fmt.Println(res1)
	
}
func pt(f func(int) int, x int) int{
	return f(x)
}


