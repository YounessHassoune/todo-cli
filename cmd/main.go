package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/YounessHassoune/todo-go/pkg"
)

const (
	todosFile = ".todos.json"
)

func check(e error) {
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
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

func main() {

	add := flag.Bool("add", false, "add new todo")
	complete := flag.Int("complete", 0, "mark todo as complete")
	dele := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todosFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		check(err)
		todos.Add(task)
		err = todos.Store(todosFile)
		check(err)
	case *complete > 0:
		err := todos.Complete(*complete)
		check(err)
		err = todos.Store(todosFile)
		check(err)
	case *dele > 0:
		err := todos.Delete(*dele)
		check(err)
		err = todos.Store(todosFile)
		check(err)
	case *list:
		todos.List()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
