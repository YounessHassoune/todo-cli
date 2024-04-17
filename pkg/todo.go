package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	todos := *t
	if index <= 0 || index > len(todos) {
		return errors.New("invalid index")
	}

	todos[index-1].CompletedAt = time.Now()
	todos[index-1].Done = true

	return nil
}

func (t *Todos) Delete(index int) error {
	todos := *t
	if index <= 0 || index > len(todos) {
		return errors.New("invalid index")
	}

	*t = append(todos[:index-1], todos[index:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return nil
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) List() {
	for i, item := range *t {
		i++
		fmt.Printf("%d - %s\n", i, item.Task)
	}
}
