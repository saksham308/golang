package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type todo struct {
	Description string
	Id          string
	Status      string 
	CreatedAt   time.Time
	UpdatedAt 	time.Time
}

func AddTodo() {
	var task string
	fmt.Println("Enter todo description:")
	reader := bufio.NewReader(os.Stdin) 
	task, _ = reader.ReadString('\n')
	task = strings.TrimSpace(task)	
	newTodo:=todo{ Description: task,Id:uuid.New().String(),Status: "not done",CreatedAt: time.Now(),UpdatedAt: time.Now()}
	jsonData, err := json.Marshal(newTodo)
	if err!=nil{
		fmt.Println("Not able to create json data",err)
		os.Exit(2)
	}
	file,err:=os.OpenFile("todo.jsonl",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("Error creating file:", err)
		os.Exit(2)
	}
	defer file.Close()
	if _, err := file.WriteString("\n"); err != nil {
		fmt.Println("Error writing newline to file:", err)
		os.Exit(2)
	}
	_,err=file.Write(jsonData)
	if err !=nil{
		fmt.Println("Error writing to file:", err)
		os.Exit(2)
	}
}
		

func UpdateTodo() {
	var taskID string
	fmt.Println("update todo:")
	reader := bufio.NewReader(os.Stdin) 
	taskID, _ = reader.ReadString('\n')
	taskID = strings.TrimSpace(taskID)	
	file,err:=os.ReadFile("todo.jsonl")
	if err!=nil{
		fmt.Println("Couldnot open file",err)
		os.Exit(2)
	}
	var todos []todo
	var todo todo
	lines:=strings.Split(string(file),"\n")
	for _,line:=range lines{
		if strings.TrimSpace(line)!=""{
		err=json.Unmarshal([]byte(line),&todo)
		if err !=nil{
			fmt.Println("cannot unmarshal",err)
			os.Exit(2)
		}
		if todo.Id==taskID{
			var choice int
			fmt.Println("Enter 1 to change description. enter 2 to change the status to done")
			fmt.Scanf("%d",&choice)
			reader.ReadString('\n')
			switch choice{
			case 1:
				fmt.Println("Update description:")
					task, err := reader.ReadString('\n')
					if err != nil {
						fmt.Println("Error reading input:", err)
						return
					}
					todo.Description = strings.TrimSpace(task)
			case 2 :
				if todo.Status=="not done"{
					todo.Status="done"
				}else{
					todo.Status="not done"
				}
			}
			todo.UpdatedAt=time.Now()
		}
	}
		todos=append(todos, todo)
	}
		
		

	json_file, err := os.OpenFile("todo.jsonl", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(2)
	}
	defer json_file.Close()
	for idx,todo:=range todos{
		jsonData, err := json.Marshal(todo)
		if err!=nil{
			fmt.Println("Not able to create json data",err)
			os.Exit(2)
		}
		if idx==len(todos)-1{
			_, err = json_file.Write(jsonData)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(2)
			}
			break
		}
		_, err = json_file.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(2)
		}
	}
}
func RemoveTodo(){
	var taskID string
	fmt.Println("remove todo:")
	reader := bufio.NewReader(os.Stdin) 
	taskID, _ = reader.ReadString('\n')
	taskID = strings.TrimSpace(taskID)	
	file,err:=os.ReadFile("todo.jsonl")
	if err!=nil{
		fmt.Println("Couldnot open file",err)
		os.Exit(2)
	}
	var todos []todo
	var todo todo
	lines:=strings.Split(string(file),"\n")
	for _,line:=range lines{
		if strings.TrimSpace(line)!=""{
		err=json.Unmarshal([]byte(line),&todo)
		if err !=nil{
			fmt.Println("cannot unmarshal",err)
			os.Exit(2)
		}
		if todo.Id!=taskID{
			todos=append(todos, todo)
		}}}
	

	json_file, err := os.OpenFile("todo.jsonl", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(2)
	}
	defer json_file.Close()
	for idx,todo:=range todos{
		jsonData, err := json.Marshal(todo)
		if err!=nil{
			fmt.Println("Not able to create json data",err)
			os.Exit(2)
		}
		if idx==len(todos)-1{
			_, err = json_file.Write(jsonData)
			if err != nil {
				fmt.Println("Error writing to file:", err)
				os.Exit(2)
			}
			break
		}
		_, err = json_file.WriteString(string(jsonData) + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(2)
		}
	}
}
func ListTodosNotDone(){
	file,err:=os.ReadFile("todo.jsonl")
	if err!=nil{
		fmt.Println("Error opening file",err)
	}
	stringData:=strings.Split(string(file),"\n")
	var todo todo
	for _,line :=range stringData{
		json.Unmarshal([]byte(line),&todo)
		if todo.Status=="not done"{
		fmt.Println(todo.Description)}
	}
}
func ListComplete(){
	file,err:=os.ReadFile("todo.jsonl")
	if err!=nil{
		fmt.Println("Error opening file",err)
	}
	stringData:=strings.Split(string(file),"\n")
	var todo todo
	for _,line :=range stringData{
		json.Unmarshal([]byte(line),&todo)
		if todo.Status=="done"{
		fmt.Println(todo.Description)}
	}
}