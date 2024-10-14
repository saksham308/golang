package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func checkFileExists(filename string)bool{
	_, error := os.Stat(filename)
	return !errors.Is(error,os.ErrNotExist)
}
func newDeck() deck{
	cards:=deck{}
	cardSuits:=[]string{"spades","hearts","diamonds","clubs"}
	cardValues:=[]string{"ace","two","three","four","five"}
	for _,suit :=range(cardSuits){
		for _,value :=range(cardValues){
			cards=append(cards,value+" of "+suit)
		}
	}
	return cards
}
func (d deck) toString() string{
	str1:=[]string(d)
	result:= strings.Join(str1,",")
	return result
}
func (d deck) shuffle()  {
	source:=rand.NewSource(time.Now().UnixNano())
	rand:=rand.New(source)
	for i,_ :=range d{
		rand_idx:=rand.Intn(len(d)-1)
		d[i],d[rand_idx]=d[rand_idx],d[i]

	}
}
func (d deck) print()  {
	for i, ele := range d {
		fmt.Println(i,ele)

	}
}
func newDeckFromFile(filename string)(deck,error){
	byte_slice,err:=os.ReadFile(filename)
	if err !=nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	str1:=string(byte_slice)
	string2:=strings.Split(str1,",")
	deck:=deck(string2)
	return deck,err

}
func deal(d deck,handSize int) (deck,deck) {
	return d[:handSize],d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	byte_slice:=[]byte(d.toString())

	err:=os.WriteFile(filename,byte_slice,0666)
	return err
}