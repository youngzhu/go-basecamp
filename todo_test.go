package basecamp

import (
	"fmt"
	"github.com/youngzhu/godate"
	"testing"
)

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

func TestCreateTodo_xingye(t *testing.T) {
	dateStart := godate.MustDate(2023, 11, 1)
	dateEnd := godate.MustDate(2024, 4, 1)

	d := dateStart.Time
	for d.Before(dateEnd.Time) {
		content := fmt.Sprintf("兴业银行100-25（%d月）", int(d.Month()))
		dueOn := godate.MustDate(d.Year(), int(d.Month()), 25)
		//fmt.Println(content, dueOn)
		createCouponTodo(content, dueOn.String())
		d = d.AddDate(0, 1, 0)
	}
}

func createCouponTodo(content, dueOn string) {
	todo := Todo{
		Content: content,
		DueOn:   dueOn,
	}
	err := CreateTodo("MeTime", "To-dos", "券", todo)
	if err != nil {
		panic(err)
	}
}
