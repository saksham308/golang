package main

import (
	"fmt"
	"net/http"
)

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	c:=make(chan string) //this is an unbuffered channel (always blocking and can only send the message to the channel only if some other channel is listening to it). the channel can only send strings and not any other dtype

	// c<-"sak"  limitation of unbuffered channel the send operation c <- "sak" cannot proceed until something receives that value from the channel

	// msg:=<-c     (However, since the receiving operation msg := <-c happens after the send in the same goroutine, the program never reaches the receive step because it's waiting at the send step indefinitely. This results in all goroutines being blocked (or "asleep"), which Go detects and throws a deadlock error.)

	// fmt.Println(msg)  This causes a deadlock: both sending and receiving operations need to happen in different goroutines for the communication to succeed, but in your code, both the send and receive happen in the main goroutine.
	for _, web := range websitelist {
		go getStatusCode(web,c)

	}
	for {
		msg_ch:=<-c //receiving a message from channel
		go getStatusCode(msg_ch ,c) //blocking code-doesn't let the loop progress further until and unless a value is returned from the channel and printed(buffered channel)
	}
}

func getStatusCode(endpoint string,c chan string) {

	_, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS DOWN")
		c<-endpoint //sending a message to a channel
		return
	}
	fmt.Println("LINK IS UP")
	c<-endpoint

}
