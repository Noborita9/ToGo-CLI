package domain

import "time"

type Task struct {
	Id          int
	Completed   bool
	Priority    string
	Description string
	CreatedAt   time.Time
}
