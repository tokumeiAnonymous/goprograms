package todo_test

import (
	"os"
	"testing"

	"pragprog.com/rggo/goprograms/todo"
)

func TestAdd(t *testing.T) {
	ls := todo.List{}
	taskName := "Test Task"
	ls.Add(taskName)

	if ls[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, ls[0].Task)
	}
}

func TestComplete(t *testing.T) {
	ls := todo.List{}
	taskName := "Test Task"
	ls.Add(taskName)

	if ls[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, ls[0].Task)
	}
	if ls[0].Done {
		t.Errorf("Test task should still not be completed!")
	}
	ls.Complete(1)
	if !ls[0].Done {
		t.Errorf("Test task should still already be completed!")
	}
}

func TestDelete(t *testing.T) {
	ls := todo.List{}
	tasks := []string{
		"Test task 1",
		"Test task 2",
		"Test task 3",
	}
	for _, v := range tasks {
		ls.Add(v)
	}
	if ls[0].Task != tasks[0] {
		t.Errorf("Expected list length %q, got %q instead.", tasks[0], ls[0].Task)
	}
	ls.Delete(2)
	if len(ls) != 2 {
		t.Errorf("Expected list length %d, got %d instead.", 2, len(ls))
	}
	if ls[1].Task != tasks[2] {
		t.Errorf("Expected list length %q, got %q instead.", tasks[2], ls[1].Task)
	}
}

func TestSaveGet(t *testing.T) {
	ls1 := todo.List{}
	ls2 := todo.List{}

	taskName := "Test task"
	ls1.Add(taskName)

	if ls1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, ls1[0].Task)
	}
	tf, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}
	defer os.Remove(tf.Name())

	if err := ls1.Save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}
	if err := ls2.Get(tf.Name()); err != nil {
		t.Fatalf("Error saveing list to file: %s", err)
	}
	if ls1[0].Task != ls2[0].Task {
		t.Errorf("Task %q should match %q task", ls1[0].Task, ls2[0].Task)
	}
}
