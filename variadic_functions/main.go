package main

import "fmt"

func sum(nums ...int) int {
	res := 0
	for _, ele := range nums {
		res += ele
	}
	fmt.Println(res)
	return res
}

func main() {
	sum(2, 4, 2, 6, 7)

	nums:= []int{1,3,5,2}
	sum(nums...)
}