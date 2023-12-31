package todo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// The todo data type
type Todo struct {
	item string
	done bool
}

// Add an item to the todolist
func addItem(item string, todolist *[]Todo) {
	var todoItem = Todo{item: item}
	*todolist = append(*todolist, todoItem)
}

// List all items in the todolist
func listTodos(todolist *[]Todo) {
	for i, todo := range *todolist {
		var format string = "\t[%d] %s\n"

		if todo.done {
			var crossedOutText = color.New(color.Faint).Add(color.CrossedOut)
			crossedOutText.Printf(format, i+1, todo.item)
		} else {
			fmt.Printf(format, i+1, todo.item)
		}
	}
}

// Change done property to true
func completeTodo(index int, todolist *[]Todo) {
	modifiedTodoItem := (*todolist)[index]
	modifiedTodoItem.done = true

	(*todolist)[index] = modifiedTodoItem
}

// Change done property for fall todos
func completeAllTodos(todolist *[]Todo) {
	for index, todoItem := range *todolist {
		todoItem.done = true
		(*todolist)[index] = todoItem
	}
}

// Remove todo item from todolist
func removeTodo(index int, todolist *[]Todo) {
	(*todolist) = append((*todolist)[:index], (*todolist)[index+1:]...)
}

// Read user input and format them
func readUserInput(scanner *bufio.Scanner) (bool, string, string) {
	var command, item string

	fmt.Print("todo > ")
	scanned := scanner.Scan()

	if !scanned {
		return false, "", ""
	}

	line := scanner.Text()
	parts := strings.SplitN(line, ":", 2)

	command = parts[0]

	if len(parts) > 1 {
		item = strings.TrimSpace(parts[1])
	}

	// remove any leading or trailing whitespace and store as lowercase
	command = strings.TrimSpace(command)
	command = strings.ToLower(command)

	return true, command, item
}

// Handle commands, returns true to continue, false to quit
func commandHandler(command string, item string, todolist *[]Todo) bool {
	if command == "quit" || command == "q" {
		fmt.Println("Exit program")
		return false
	}

	if command == "new" || command == "n" {
		addItem(item, todolist)
		return true
	}

	if command == "list" || command == "l" {
		listTodos(todolist)
		return true
	}

	if command == "done" || command == "d" {
		if item == "all" {
			completeAllTodos(todolist)
		} else {
			var number, err = strconv.Atoi(item)

			// Handle any errors during conversion
			if err != nil {
				fmt.Println("Invalid index received.")
			}

			if number > 0 && number <= len(*todolist) {
				completeTodo(number-1, todolist)
			}
		}
	}

	if command == "clear" || command == "c" {
		*todolist = []Todo{}
	}

	if command == "remove" || command == "r" {
		var number, err = strconv.Atoi(item)

		// Handle any errors during conversion
		if err != nil {
			fmt.Println("Invalid number received.")
		}

		if number > 0 && number <= len(*todolist) {
			removeTodo(number-1, todolist)
		}
	}

	return true
}

// Initialize the todo to listen for commands
func Initialize() {
	var todolist []Todo
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		success, command, item := readUserInput(scanner)
		if !success {
			break
		}

		shouldContinue := commandHandler(command, item, &todolist)

		if !shouldContinue {
			break
		}
	}
}
