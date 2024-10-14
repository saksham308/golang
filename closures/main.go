package main

import "fmt"

func counter() func() int {
	count:=0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	num := counter()
	ans1 := num()
	fmt.Println(ans1)
	ans2:=num()
	fmt.Println(ans2)
}
