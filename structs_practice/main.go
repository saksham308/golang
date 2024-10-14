package main

import "fmt"
type address struct{
	email string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact address
}

func(p person) print(){
	fmt.Printf("%+v",p)
}

func main() {
	alex := person{
	"Alex", 
	"Andersom", 
	address{email:"sak@gmail.com",zipCode:432},
}

	alex2 := person{
		firstName: "Alex",
		lastName: "new",
		contact:  address{email:"sak.com",zipCode:259},
	}
		
	fmt.Println(alex,alex2)
	alex.firstName="new_alex"
	alex2.lastName="anderson"
	alex.print()
	alex2.print()
	var alex3 person
	alex3.print()
	
}