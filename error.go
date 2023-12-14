package basecamp

import "errors"

var (
	ErrNotFoundProject    = errors.New("not found the project")
	ErrNotFoundDock       = errors.New("not found the dock")
	ErrNotFoundSchedule   = errors.New("not found the schedule")
	ErrNotFoundCardColumn = errors.New("not found the card column")
	ErrNotFoundTodoList   = errors.New("not found the todo list")

	ErrNotSupport = errors.New("not support yet")
)
