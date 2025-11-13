package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"todo/cmd"
	"todo/internal"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Unable to get current working directory: %v\n", err)
		os.Exit(1)
	}

	dirName := filepath.Base(wd)
	dirName = strings.ReplaceAll(dirName, " ", "_")
	dirName = strings.ReplaceAll(dirName, "/", "_")
	dirName = strings.ReplaceAll(dirName, "\\", "_")

	filename := dirName + "_ToDoList.csv"

	ToDoList := internal.InitToDoList(filename)
	cmd.SetToDoList(ToDoList)
	cmd.Execute()
}
