package main

import "fmt"

type english struct{}
type spanish struct{}
type bot interface {
	getGreeting() string
	// getSpeak()
}
// func(e english) getSpeak(){}
// func(s spanish) getSpeak(){}

func (e english) getGreeting() string {
	return "hello in english"
}
func (s spanish) getGreeting() string {
	return "hola"
}
func printGreeting(b bot){
    fmt.Println(b.getGreeting())
}
func main() {
	e := english{}
	s := spanish{}
	printGreeting(e)
    printGreeting(s)
}