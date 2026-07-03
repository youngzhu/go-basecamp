package basecamp

import (
	"fmt"
	"github.com/youngzhu/godate"
	"testing"
)

func TestBaseCamp_getTodoSet(t *testing.T) {
	todoSetTitle := "To-dos"
	todoSet, err := _bc.getTodoSetDock("MeTime", todoSetTitle)
	if err != nil {
		t.Fatal(err)
	}

	if todoSet.Title != todoSetTitle {
		t.Errorf("todoset title not match, want: %s, but got: %s",
			todoSetTitle, todoSet.Title)
	}
}

func TestBaseCamp_getTodoLists(t *testing.T) {
	projectName := "MeTime"
	todoSetTitle := "To-dos"

	todoSet, err := _bc.getTodoSetDock(projectName, todoSetTitle)
	if err != nil {
		t.Fatal(err)
	}

	// cache test
	_bc.getTodoLists(projectName, todoSetTitle)
	todoLists, err := _bc.getTodoLists(projectName, todoSetTitle)

	if err != nil {
		t.Fatal(err)
	}

	if todoSet.TodolistsCount != len(todoLists) {
		t.Errorf("count of todo list, want: %d, but got: %d",
			todoSet.TodolistsCount, len(todoLists))
	}
}

func TestBaseCamp_getTodoListByTitle(t *testing.T) {
	todoListTitle := "券"
	todoList, err := _bc.getTodoListByTitle("MeTime", "To-dos", todoListTitle)

	if err != nil {
		t.Fatal(err)
	}

	if todoListTitle != todoList.Title {
		t.Errorf("todo list title not match, want: %q, but got: %q",
			todoListTitle, todoList.Title)
	}
}

func TestBaseCamp_getTodoByTitle(t *testing.T) {
	todoTitle := "指派测试"
	todo, err := _bc.getTodoByTitle("MeTime", "To-dos", "券", todoTitle)

	if err != nil {
		t.Fatal(err)
	}

	if todoTitle != todo.Content {
		t.Errorf("todo list title not match, want: %q, but got: %q",
			todoTitle, todo.Content)
	}
}

func TestCreateTodo(t *testing.T) {
	createCouponTodo("测试", "", "", 36278984)
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
	content = "光大 微信 20-1.7"
	dueOn = "2025-11-14"

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

func TestCreateTodo_jianshe(t *testing.T) {
	dueOn := "2024-11-15"

	//content = "建设银行6.01-6"
	//for i := 0; i < 3; i++ {
	//	createCouponTodo(content, dueOn, "")
	//}

	titles := []string{"500-5", "-3", "66-6", "-5", "-2", "2000-10", "-5", "100-20"}
	for _, title := range titles {
		createCouponTodo("建设银行 "+title, dueOn, "")
	}
}

func TestCreateTodo_jiaotong(t *testing.T) {
	var content, dueOn string

	content = "交通银行储蓄卡 微信 随机减"
	dueOn = "2025-11-19"

	for i := 0; i < 5; i++ {
		createCouponTodo(content, dueOn, "")
	}
}

func createCouponTodo(content, dueOn, startsOn string, assigneeIds ...int64) {
	todo := Todo{
		Content:     content,
		DueOn:       dueOn,
		StartsOn:    startsOn,
		AssigneeIds: assigneeIds,
	}
	err := AddTodo("MeTime", "To-dos", "券", todo)
	if err != nil {
		panic(err)
	}
}
