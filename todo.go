package basecamp

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// TodoSetDock all to-do lists under a project are children of a to-do set resource
type TodoSetDock struct {
	Id               int       `json:"id"`
	Status           string    `json:"status"`
	VisibleToClients bool      `json:"visible_to_clients"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Title            string    `json:"title"`
	InheritsStatus   bool      `json:"inherits_status"`
	Type             string    `json:"type"`
	Url              string    `json:"url"`
	AppUrl           string    `json:"app_url"`
	BookmarkUrl      string    `json:"bookmark_url"`
	Position         int       `json:"position"`
	Bucket           struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"bucket"`
	Creator struct {
		Id             int       `json:"id"`
		AttachableSgid string    `json:"attachable_sgid"`
		Name           string    `json:"name"`
		EmailAddress   string    `json:"email_address"`
		PersonableType string    `json:"personable_type"`
		Title          string    `json:"title"`
		Bio            string    `json:"bio"`
		Location       string    `json:"location"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Admin          bool      `json:"admin"`
		Owner          bool      `json:"owner"`
		Client         bool      `json:"client"`
		Employee       bool      `json:"employee"`
		TimeZone       string    `json:"time_zone"`
		AvatarUrl      string    `json:"avatar_url"`
		Company        struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"company"`
		CanManageProjects bool `json:"can_manage_projects"`
		CanManagePeople   bool `json:"can_manage_people"`
	} `json:"creator"`
	Completed        bool   `json:"completed"`
	CompletedRatio   string `json:"completed_ratio"`
	Name             string `json:"name"`
	TodolistsCount   int    `json:"todolists_count"`
	TodolistsUrl     string `json:"todolists_url"`
	AppTodoslistsUrl string `json:"app_todoslists_url"`
}

func (d *TodoSetDock) DockType() dockType {
	return TypeTodoSet
}

func (d *TodoSetDock) DockTitle() string {
	return d.Title
}

type TodoList struct {
	Id               int64     `json:"id"`
	Status           string    `json:"status"`
	VisibleToClients bool      `json:"visible_to_clients"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Title            string    `json:"title"`
	InheritsStatuS   bool      `json:"inherits_statu s"`
	Type             string    `json:"type"`
	Url              string    `json:"url"`
	AppUrl           string    `json:"app_url"`
	BookmarkUrl      string    `json:"bookmark_url"`
	SubscriptionUrl  string    `json:"subscription_url"`
	CommentsCount    int       `json:"comments_count"`
	CommentsUrl      string    `json:"comments_url"`
	Position         int       `json:"position"`
	Parent           struct {
		Id     int64  `json:"id"`
		Title  string `json:"title"`
		Type   string `json:"type"`
		Url    string `json:"url"`
		AppUrl string `json:"app_url"`
	} `json:"parent"`
	Bucket struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"bucket"`
	Creator struct {
		Id             int         `json:"id"`
		AttachableSgid string      `json:"attachable_sgid"`
		Name           string      `json:"name"`
		EmailAddress   string      `json:"email_address"`
		PersonableType string      `json:"personable_type"`
		Title          interface{} `json:"title"`
		Bio            string      `json:"bio"`
		Location       interface{} `json:"location"`
		CreatedAt      time.Time   `json:"created_at"`
		UpdatedAt      time.Time   `json:"updated_at"`
		Admin          bool        `json:"admin"`
		Owner          bool        `json:"owner"`
		Client         bool        `json:"client"`
		Employee       bool        `json:"employee"`
		TimeZone       string      `json:"time_zone"`
		AvatarUrl      string      `json:"avatar_url"`
		Company        struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"company"`
		CanPing           bool `json:"can_ping"`
		CanManageProjects bool `json:"can_manage_projects"`
		CanManagePeople   bool `json:"can_manage_people"`
	} `json:"creator"`
	Description    string `json:"description"`
	Completed      bool   `json:"completed"`
	CompletedRatio string `json:"completed_ratio"`
	Name           string `json:"name"`
	TodosUrl       string `json:"todos_url"`
	GroupsUrl      string `json:"groups_url"`
	AppTodosUrl    string `json:"app_todos_url"`
}

type Todo struct {
	Content     string `json:"content"`     // **Required parameters** for what the to-do is for
	Description string `json:"description"` // containing information about the to-do
	DueOn       string `json:"due_on"`      // a date when the to-do should be completed
	StartsOn    string `json:"starts_on"`   // allows the to-do to run from this date to the `due_on` date
}

// AddTodo creates a to-do
// POST /buckets/1/todolists/3/todos.json
// creates a to-do in the project with ID `1` and under the to-do list with an ID of `3`.
func (bc *BaseCamp) AddTodo(projectName, todoSetTitle, todoListTitle string, todo Todo) error {
	todoLists, err := bc.getTodoLists(projectName, todoSetTitle)
	if err != nil {
		return err
	}

	var todoList TodoList
	for _, list := range todoLists {
		if todoListTitle == list.Title {
			todoList = list
		}
	}

	_, err = bc.doPost(todoList.TodosUrl, todo)

	return err
}

// AddTodoList Create a to-do list
// POST /buckets/1/todosets/3/todolists.json
// creates a to-do list in the project with ID `1` and under the to-do set with an ID of `3`.
func (bc *BaseCamp) AddTodoList(projectName, todoSetTitle, todoListName string) error {
	todoSet, err := bc.getTodoSetDock(projectName, todoSetTitle)
	if err != nil {
		return err
	}

	todoList := TodoList{Name: todoListName}

	_, err = bc.doPost(todoSet.TodolistsUrl, todoList)

	return err
}

func (bc *BaseCamp) AddTodoListAndTodos(projectName, todoSetTitle, todoListAndTodos string) error {
	arr := strings.Split(todoListAndTodos, "\n")

	// 处理首尾空行的问题
	// 通过 `` 定义长字符串时会有这个问题
	firstNonblankIdx := 0
	for {
		if arr[firstNonblankIdx] != "" {
			break
		}
		firstNonblankIdx++
	}

	todoListTitle := arr[firstNonblankIdx]

	err := bc.AddTodoList(projectName, todoSetTitle, todoListTitle)
	if err != nil {
		return err
	}

	todos := arr[firstNonblankIdx+1:]
	for _, todo := range todos {
		if todo == "" {
			break
		}
		err = bc.AddTodo(projectName, todoSetTitle, todoListTitle, Todo{Content: todo})
		if err != nil {
			return err
		}
	}
	return err
}

func (bc *BaseCamp) getTodoSetDock(projectName, todoSetTitle string) (*TodoSetDock, error) {
	d, err := bc.getDock(projectName, TypeTodoSet, todoSetTitle)
	if err != nil {
		return nil, err
	}

	return d.(*TodoSetDock), nil
}

func (bc *BaseCamp) getTodoListByTitle(projectName, todoSetTitle, todoListTitle string) (TodoList, error) {
	todoList := TodoList{}

	todoLists, err := bc.getTodoLists(projectName, todoSetTitle)
	if err != nil {
		return todoList, err
	}

	for _, list := range todoLists {
		if todoListTitle == list.Title {
			return list, nil
		}
	}

	return todoList, fmt.Errorf("%w: %s", ErrNotFoundTodoList, todoListTitle)
}

func (bc *BaseCamp) getTodoLists(projectName, todoSetTitle string) ([]TodoList, error) {
	key := buildTodoListsKey(projectName, todoSetTitle)
	if v, ok := bc.todoListsMap[key]; ok {
		return v, nil
	}

	todoSet, err := bc.getTodoSetDock(projectName, todoSetTitle)
	if err != nil {
		return nil, err
	}

	resp, err := bc.doGet(todoSet.TodolistsUrl)
	if err != nil {
		return nil, err
	}

	var todoLists []TodoList
	err = json.Unmarshal(resp, &todoLists)
	if err != nil {
		return nil, err
	}

	// cached
	bc.todoListsMap[key] = todoLists

	return todoLists, nil
}

func buildTodoListsKey(projectName, todoSetTitle string) string {
	return projectName + keySplit + todoSetTitle
}
