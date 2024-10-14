package main

import (
	"fmt"
)
func main() {
	// var str1 colour
	// str1 = "saj"
	// fmt.Println(str1)
	// cards.print()
	
	num:= []int{}
	fmt.Println(num,len(num),cap(num))
	var cards deck
	if checkFileExists("my_cards"){
		fmt.Println("cards being created from my_cards file")
		cards,_=newDeckFromFile("my_cards")
		
	} else{
		fmt.Println("cards being created locally")
		cards=newDeck()

	}

	// cards, _ = newDeckFromFile("my_cards")
	// new:=[]int{1,2}
	// fmt.Println(new[len(new)-1])
	fmt.Println(cards)
	cards.shuffle()

	fmt.Println(cards)
	cards.saveToFile("my_cards")
	cards.saveToFile("shuffle_my_cards")
	
	// d1, d2 := deal(cards, 2)
	// d1.print()
	// d2.print()
}