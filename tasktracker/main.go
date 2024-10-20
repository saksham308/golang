package main

import (
	"fmt"
)

func main() {
	var choice int
	// Create a new reader to read input

	fmt.Println("1 to add todo. 2 to remove todo. 3 to list todos. 4 to remove todos")
	fmt.Scanf("%d\n", &choice) // Read the choice

	switch choice {
	case 1:
		AddTodo()
	case 2:
		RemoveTodo()
	case 3:
		ListTodos()
	case 4:
		UpdateTodo()
	default:
		fmt.Println("Invalid option. Please choose between 1 and 4.")
	}
}
