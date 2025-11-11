package main

import (
	"todo-list-cli/cmd"
	"todo-list-cli/internal"
)

func main() {
	ToDoList := internal.InitToDoList("ToDos.csv")
	cmd.SetToDoList(ToDoList)
	cmd.Execute()
}
