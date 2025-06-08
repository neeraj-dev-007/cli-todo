package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to  CLI to-do application!")
	fmt.Println("Type a command: add <task>, list, done <task> or exit")

	tasks := []string{}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
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
			tasks = append(tasks, task)
			fmt.Printf("Task added: %s\n", task)
		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks available.")
			} else {
				fmt.Println("Tasks:")
				for i, task := range tasks {
					fmt.Printf("%d: %s\n", i+1, task)
				}
			}
		case "done":
			if argsLen < 2 {
				fmt.Println("Please provide a task to mark as done.")
				continue
			}
			task := args[1]
			for i, t := range tasks {
				if t == task {
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
