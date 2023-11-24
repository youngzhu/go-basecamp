package basecamp

import "testing"

func TestProject_getTodoSet(t *testing.T) {
	project, _ := GetProjectByName("MeTime")
	todoSetTitle := "To-dos"
	project.getTodoSet(todoSetTitle)

	//if cardColumn.Title != columnName {
	//	t.Errorf("card column title not match, want: %s, but got: %s",
	//		columnName, cardColumn.Title)
	//}
}

func TestProject_getTodoLists(t *testing.T) {
	project, _ := GetProjectByName("MeTime")
	todoSetTitle := "To-dos"
	todoSet := project.getTodoSet(todoSetTitle)
	todoLists := project.getTodoLists(todoSetTitle)

	if todoSet.TodolistsCount != len(todoLists) {
		t.Errorf("count of todo list, want: %d, but got: %d",
			todoSet.TodolistsCount, len(todoLists))
	}
}

func TestProject_getTodoListByTitle(t *testing.T) {
	project, _ := GetProjectByName("MeTime")
	todoSetTitle := "To-dos"
	todoListTitle := "券"
	todoList := project.getTodoListByTitle(todoSetTitle, todoListTitle)

	if todoListTitle != todoList.Title {
		t.Errorf("todo list title not match, want: %q, but got: %q",
			todoListTitle, todoList.Title)
	}
}
