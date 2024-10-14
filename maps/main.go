package main

import (
	"fmt"
	"maps"
)

func main() {
	//map=> dictionary/hashmaps
	//create map
	m := map[string]int{"phone":50,"price":30}
	m2 := map[string]int{"phone":50,"price":30}
	fmt.Println(m["sak"])
	// new_m:=m

	fmt.Println(maps.Equal(m,m2))
	
	//if key is not present in map than map returns 0 or empty str or false bool
	fmt.Println(m["sa"])
	fmt.Println(len(m))

	// deleting an element
	delete(m,"price")
	fmt.Println(m)


	ele1,ele2:=m["phone"]
	fmt.Println(ele1,ele2)

	// clearing a map
	clear(m)
	fmt.Println(m)
	for key,val:=range m2{
		fmt.Println("key " ,key, " value ",val)
	}
}