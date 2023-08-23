package main

import (
	"fmt"
	"os"
	"strings"

	"pragprog.com/rggo/goprograms/todo"
)

const todoFileName = ".todo.json"

func main() {
	ls := &todo.List{}

	if err := ls.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(os.Args) == 1 {
		for _, item := range *ls {
			fmt.Println(item.Task)
		}
	} else {
		item := strings.Join(os.Args[1:], " ")
		ls.Add(item)

		if err := ls.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
