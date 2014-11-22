package models

import (
	"encoding/base64"
	"fmt"
	"time"
)

// Tasks ...
type Tasks struct {
	Dict map[string]Task
}

// Task represents task
type Task struct {
	Value   string
	Created time.Time
}

// NewTask ...
func NewTask(val string) Task {
	return Task{
		Value:   val,
		Created: time.Now(),
	}
}

// Add ...
func (tasks *Tasks) Add(task Task) error {
	if tasks.Dict == nil {
		tasks.Dict = map[string]Task{}
	}
	key := base64.StdEncoding.EncodeToString([]byte(task.Value))
	if _, ok := tasks.Dict[key]; ok {
		return fmt.Errorf("`%s` is already in the tasks", task.Value)
	}
	tasks.Dict[key] = task
	return nil
}

// List ...
func (tasks *Tasks) List() []string {
	res := []string{}
	for _, task := range tasks.Dict {
		res = append(res, task.Value)
	}
	return res
}

// Remove ...
func (tasks *Tasks) Remove(task Task) error {
	key := base64.StdEncoding.EncodeToString([]byte(task.Value))
	if _, ok := tasks.Dict[key]; !ok {
		return fmt.Errorf("`%s` is not in the tasks", task.Value)
	}
	delete(tasks.Dict, key)
	return nil
}
