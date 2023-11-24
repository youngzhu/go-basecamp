package basecamp

import (
	"encoding/json"
	"net/http"
	"time"
)

// TodoSet all to-do lists under a project are children of a to-do set resource
type TodoSet struct {
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

func (p *Project) getTodoSet(todoSetTitle string) *TodoSet {
	d := p.getDock(todoSetTitle)
	if d == nil {
		return nil
	}

	resp, err := doRequest(d.Url, http.MethodGet, nil)
	if err != nil {
		return nil
	}

	var todoSet *TodoSet
	err = json.Unmarshal(resp, &todoSet)
	if err != nil {
		return nil
	}

	return todoSet
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

func (p *Project) getTodoLists(todoSetTitle string) []TodoList {
	todoSet := p.getTodoSet(todoSetTitle)

	resp, err := doRequest(todoSet.TodolistsUrl, http.MethodGet, nil)
	if err != nil {
		return nil
	}

	var todoLists []TodoList
	err = json.Unmarshal(resp, &todoLists)
	if err != nil {
		return nil
	}

	return todoLists
}

func (p *Project) getTodoListByTitle(todoSetTitle, todoListTitle string) TodoList {
	todoLists := p.getTodoLists(todoSetTitle)

	for _, list := range todoLists {
		if todoListTitle == list.Title {
			return list
		}
	}

	return TodoList{}
}
