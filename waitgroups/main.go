package main

import (
	"fmt"
	"net/http"
	"sync"
)
var wg sync.WaitGroup //pointers  if a WaitGroup is explicitly passed into functions, it should be done by pointer.

func main() {
	//when we run a go program it automatically creates a go routine. This go routine executes the program line by line.
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	//if we add wg.Add(1) here than only once statement will be logged 
	for _, web := range websitelist {
		wg.Add(1)
		go getStatusCode(web)
		 //increases the waitgroup counter by 1
		
		//if we add "wg.Wait()" statement here than the program execution will be synchronous because it won't let the loop execute further until and unless the waitgroup counter again returns to 0. 
	}
	wg.Wait() //makes the main method wait until the waitgroup counter becomes 0

	
}
func getStatusCode(endpoint string) {
	defer wg.Done() //decreases the waitgroup counter by 1 at the end of the execution of this function, that this go routine has finished it's execution

	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

	}
}