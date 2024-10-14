package main

import "fmt"

// numUsers:=444--> can't declare variables in this way outside functions because variable declared outside function
// are global variable and global variables should always be declared using "var" keyword4

// value types in go- int,float,string,bool,structs passed by value to functions
// reference types in go-slices,maps,channels,pointers, functions --> pass by reference to functions

const Num int=4 //Public variable

func main() {
	fmt.Println(("variables"))
	//ways of declaring variable
	var str = "salsham"
	var num int=47
	numUsers:=444

	fmt.Println(str,num,numUsers,Num)

	fmt.Printf("Variable is of type %T",str)

}
