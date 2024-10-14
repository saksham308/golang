package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(a)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(("Enter a number"))
	text,_:=reader.ReadString('\n')
	new_string:=strings.TrimSpace((text))
	// integer
	num,err:=strconv.Atoi(new_string)

	// float
	numFloat,_:=strconv.ParseFloat(new_string,64)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf(" the new no. is %d,%T",num+1,num)
	fmt.Printf(" the new no. is %f,%T",numFloat+1,numFloat)

}