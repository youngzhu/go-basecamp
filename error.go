package basecamp

import "errors"

var (
	ErrNotFoundProject   = errors.New("not found the project")
	ErrNotFoundSchedule  = errors.New("not found the schedule")
	ErrNotFoundCardTable = errors.New("not found the card table")
)
