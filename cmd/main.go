package main

import (
	"flag"
	"fmt"
	"os"

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
		todos.Add("Sample todo")
		err := todos.Store(todosFile)
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
