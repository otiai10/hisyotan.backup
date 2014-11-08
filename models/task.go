package models

import "time"

// Task represents task
type Task struct {
	Value   string
	Created time.Time
}
