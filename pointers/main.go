package main

import "fmt"
func changeInt(ptr *int,originalNum int){
	*ptr=*ptr*2
	fmt.Println("the copied vairiable 'originalNum' passed as 2nd param to our function with diff memory location",&originalNum,originalNum)
	fmt.Println("the changed variable *ptr referencing to 'num' in main function ",ptr,*ptr)

}
func(p *person) changeName(name string){
	p.name=name
}
type person struct{
	name string
}
func main() {
	num := 8
	// to get the memory location of a variable
	ptrToNums := &num

	// to 
	fmt.Println(num,ptrToNums)
	changeInt(ptrToNums,num)
	p1:=person{name:"saksham"}

	// go shortcut to directly use the type of object instead of using pointer 
	p1.changeName("bhatnagar")

	// this is the intutive way to do the change the name using receiver function
	p1_ptr:=&p1
	p1_ptr.changeName("bhatnagar")
}