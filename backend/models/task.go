package models

type Task struct {
	ID          string
	Title       string
	Description string
	State       TastState
	Date        string
}

type TastState int

const (
	Todo TastState = iota
	InProgress
	Done
)
