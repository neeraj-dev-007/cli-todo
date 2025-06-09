package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Task struct {
	Description string
}

const banner = `
 _____ ___________ _____    ___  ____________ 
|_   _|  _  |  _  \  _  |  / _ \ | ___ \ ___ \
  | | | | | | | | | | | | / /_\ \| |_/ / |_/ /
  | | | | | | | | | | | | |  _  ||  __/|  __/ 
  | | \ \_/ / |/ /\ \_/ / | | | || |   | |    
  \_/  \___/|___/  \___/  \_| |_/\_|   \_|    
                                              
  `

func printBanner() {
	fmt.Println(banner)
}

func main() {
	printBanner()
	fmt.Println("Welcome to  CLI to-do application!")
	fmt.Println("Type a command: add <task>, list, done <task> or exit")

	tasks := loadTasks()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			saveTasks(tasks)
			fmt.Println("Exiting the application. Goodbye!")
			break
		}

		args := strings.SplitN(input, " ", 2)
		command := args[0]

		switch argsLen := len(args); command {
		case "add":
			if argsLen < 2 {
				fmt.Println("Please provide a task to add.")
				continue
			}
			task := args[1]
			tasks = append(tasks, Task{Description: task})
			fmt.Printf("Task added: %s\n", task)
		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				fmt.Println("Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d: %s\n", i+1, task.Description)
				}
			}
		case "done":
			if argsLen < 2 {
				fmt.Println("Please provide a task to mark as done.")
				continue
			}
			task := args[1]
			for i, t := range tasks {
				if t.Description == task {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Printf("Task marked as done: %s\n", task)
					break
				}
			}
		default:
			fmt.Println("Unknown command. Please use add, list, done or exit.")
		}
	}
}

func saveTasks(tasks []Task) {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	filePath := filepath.Join(dir, "tasks.json")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		fmt.Println("Error encoding tasks:", err)
		return
	}
	fmt.Println("Tasks saved successfully.")
}

func loadTasks() []Task {
	exePath, _ := os.Executable()
	dir := filepath.Dir(exePath)
	filePath := filepath.Join(dir, "tasks.json")

	tasks := []Task{}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return tasks // No tasks file found, return empty slice
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return tasks
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Println("Error decoding tasks:", err)
		return []Task{}
	}
	fmt.Println("Tasks loaded successfully.")
	return tasks
}
