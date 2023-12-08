package basecamp

import (
	"fmt"
	"github.com/youngzhu/godate"
	"testing"
)

func TestProject_getTodoSet(t *testing.T) {
	project, _ := GetProjectByName("MeTime")
	todoSetTitle := "To-dos"
	todoSet := project.getTodoSet(todoSetTitle)

	if todoSet.Title != todoSetTitle {
		t.Errorf("todoset title not match, want: %s, but got: %s",
			todoSetTitle, todoSet.Title)
	}
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
		createCouponTodo(content, dueOn.String(), "")
		d = d.AddDate(0, 1, 0)
	}
}

func TestCreateTodo_guangda(t *testing.T) {
	var content, dueOn string

	//content = "光大银行30-15"
	//dueOn = "2023-11-30"
	content = "光大银行20-10"
	dueOn = "2023-12-25"

	for i := 0; i < 2; i++ {
		createCouponTodo(content, dueOn, "")
	}
}

func TestCreateTodo_pufa(t *testing.T) {
	var content, dueOn, startsOn string

	content = "浦发银行26-10"
	//dueOn = "2023-11-30"
	//
	startsOn = "2023-12-27"
	dueOn = "2024-01-31"

	for i := 0; i < 4; i++ {
		createCouponTodo(content, dueOn, startsOn)
	}
}

func createCouponTodo(content, dueOn, startsOn string) {
	todo := Todo{
		Content:  content,
		DueOn:    dueOn,
		StartsOn: startsOn,
	}
	err := CreateTodo("MeTime", "To-dos", "券", todo)
	if err != nil {
		panic(err)
	}
}
