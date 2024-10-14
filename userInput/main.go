package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
	welcome:="userInput"
	reader:= bufio.NewReader((os.Stdin))

// 	In Go, when using ReadString with a delimiter, the delimiter should be a byte, not a string. The correct form is to use single quotes (') for a character (byte) literal. The newline character ('\n') is represented as a byte in Go.
// If you use double quotes (") around the delimiter, you are creating a string, not a byte.
	text,_:=reader.ReadString('\n')

	fmt.Println('a', text)
	fmt.Println((welcome));
}
// Byte:
// Represents a single byte (8 bits).
// Commonly used for binary data and individual characters.
// Example: var b byte = 'A'.

// String:
// Represents a sequence of characters.
// Immutable, meaning you cannot modify the individual characters directly.
// Example: var s string = "Hello, World!".