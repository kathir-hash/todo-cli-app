package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string `json:"title"`
	Completed   bool	`json:"Completed`
	CreatedAt   time.Time `json:"CreatedAt"`
	CompletedAt *time.Time
}

type Todos []Todo



func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, todo)
}

func (todo *Todos) validateindex(index int) error {
	if index < 0 || index >= len(*todo) {
		err := errors.New("INvalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos
	if err := t.validateindex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}
func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := t.validateindex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed
	if !isCompleted {
		completeTime := time.Now()
		t[index].CompletedAt = &completeTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) update(index int, title string) error {
	t := *todos
	if err := t.validateindex(index); err != nil {
		return err
	}
	t[index].Title = title
	return nil
}


func (todos Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("sno", "Tile", "Completed", "Created At", "Completed At")
	for i, t := range todos {
		Completed := "no"
		CompletedAt := ""
		if t.Completed {
			Completed = "yes"
			if t.CompletedAt != nil {
				CompletedAt = t.CompletedAt.Format(time.RFC1123)
			} else {
				CompletedAt = "not Completed"
			}
		}

		table.AddRow(strconv.Itoa(i), t.Title, Completed, t.CreatedAt.Format(time.RFC1123), CompletedAt)
	}
	table.Render()
}