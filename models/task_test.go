package models

import (
	"testing"

	. "github.com/otiai10/mint"
)

func TestTasks_Add(t *testing.T) {
	tasks := new(Tasks)
	err := tasks.Add(Task{Value: "おっぱい"})
	Expect(t, err).ToBe(nil)
}

func TestTasks_List(t *testing.T) {
	tasks := new(Tasks)
	err := tasks.Add(Task{Value: "おっぱい"})
	Expect(t, err).ToBe(nil)
	Expect(t, tasks.List()).ToBe([]string{"おっぱい"})
}

func TestTasks_Remove(t *testing.T) {
	tasks := new(Tasks)
	err := tasks.Add(Task{Value: "おっぱい"})
	Expect(t, err).ToBe(nil)
	Expect(t, tasks.List()).ToBe([]string{"おっぱい"})

	err = tasks.Remove(Task{Value: "おっぱい"})
	Expect(t, err).ToBe(nil)
	Expect(t, tasks.List()).ToBe([]string{})
}
