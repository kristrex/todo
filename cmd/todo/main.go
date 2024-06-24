package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github/kristrex/todo-app"
	"io"
	"os"
	"strings"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	delete := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Println(os.Stderr, err.Error())
		os.Exit(1)
	}
	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)

		todo.OutError(err)

		todos.Add(task)
		todos.Writing(todoFile)
	case *complete > 0:
		err := todos.Complete(*complete)

		todo.OutError(err)

		todos.Writing(todoFile)
	case *delete > 0:
		err := todos.Delete(*delete)

		todo.OutError(err)

		todos.Writing(todoFile)
	case *list:
		todos.Print()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("empty todo is not allowed")
	}
	return text, nil
}
