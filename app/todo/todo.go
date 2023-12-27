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
func add_item(item string, todolist *[]Todo) {
	var todo_item = Todo{item: item}
	*todolist = append(*todolist, todo_item)
}

// List all items in the todolist
func list_todos(todolist *[]Todo) {
	for i, todo := range *todolist {
		var format string = "\t[%d] %s\n"
		if todo.done {
			var green = color.New(color.Faint).Add(color.CrossedOut)
			green.Printf(format, i+1, todo.item)
		} else {
			fmt.Printf(format, i+1, todo.item)
		}
	}
}

// Change done property to true
func complete_todo(index int, todolist *[]Todo) {
	modified_todo_item := (*todolist)[index]
	modified_todo_item.done = true

	(*todolist)[index] = modified_todo_item
}

// Read user input and format them
func read_user_input(scanner *bufio.Scanner) (bool, string, string) {
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
func command_handler(command string, item string, todolist *[]Todo) bool {
	if command == "quit" || command == "q" {
		fmt.Println("Exit program")
		return false
	}

	if command == "new" || command == "n" {
		add_item(item, todolist)
		return true
	}

	if command == "list" || command == "l" {
		list_todos(todolist)
		return true
	}

	if command == "done" || command == "d" {
		var number, err = strconv.Atoi(item)

		// Handle any errors during conversion
		if err != nil {
			fmt.Println("Invalid index received.")
		}

		if number > 0 && number <= len(*todolist) {
			complete_todo(number-1, todolist)
		}
	}

	return true
}

// Initialize the todo to listen for commands
func Initialize() {
	var todolist []Todo
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)

	for {
		success, command, item := read_user_input(scanner)
		if !success {
			break
		}

		should_continue := command_handler(command, item, &todolist)

		if !should_continue {
			break
		}
	}
}
