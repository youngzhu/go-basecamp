package basecamp

import (
	"encoding/json"
	"github.com/spf13/viper"
	"github.com/youngzhu/oauth2-apps/basecamp"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// API: https://github.com/basecamp/bc3-api

var _bc *BaseCamp

func init() {
	viper.SetEnvPrefix("BASECAMP")
	viper.AutomaticEnv() // read in environment variables that match

	accountID := viper.GetString("ACCOUNT_ID")

	accessToken, refresh := basecamp.GetAccessToken()
	if refresh {
		log.Println("refresh token")
	}

	_bc = New(accountID, accessToken)
}

type BaseCamp struct {
	accountID   string // Basecamp account ID
	accessToken string

	projectsUrl string
	projects    []Project

	dockMap map[string]docker
}

func New(accountID, accessToken string) *BaseCamp {
	bc := new(BaseCamp)
	bc.accountID = accountID
	bc.accessToken = accessToken

	bc.dockMap = make(map[string]docker)

	return bc
}

// AddSchedule adds a schedule entry
// POST /buckets/1/schedules/3/entries.json
// creates a schedule entry in the project with ID 1 and under the schedule with an ID of 3.
func AddSchedule(projectName, scheduleTitle string, scheduleEntry ScheduleEntry) error {
	return _bc.AddSchedule(projectName, scheduleTitle, scheduleEntry)
}

// AddCard creates a card
// POST /buckets/1/card_tables/lists/2/cards.json
// creates a card within the column with ID 2 in the project with id 1.
func AddCard(projectName, cardTableTitle, columnTitle string, card Card) error {
	return _bc.CreateCard(projectName, cardTableTitle, columnTitle, card)
}

// AddTodo creates a to-do
// POST /buckets/1/todolists/3/todos.json
// creates a to-do in the project with ID `1` and under the to-do list with an ID of `3`.
func AddTodo(projectName, todoSetTitle, todoListTitle string, todo Todo) error {
	project, err := GetProjectByName(projectName)
	if err != nil {
		return err
	}

	entryJson, _ := json.Marshal(todo)

	todoList := project.getTodoListByTitle(todoSetTitle, todoListTitle)
	//fmt.Printf("todo list's title: %s, id: %d\n", todoList.Title, todoList.Id)
	_, err = doRequest(todoList.TodosUrl, http.MethodPost, strings.NewReader(string(entryJson)))

	return nil
}

func parseUrl(appUrl string, ids ...int) string {
	appUrl = strings.Replace(appUrl, "$ACCOUNT_ID", _bc.accountID, 1)

	u, err := url.Parse(appUrl)
	if err != nil {
		panic(err)
	}

	arr := strings.Split(u.Path, "/")
	var idx int
	for i, s := range arr {
		if strings.HasPrefix(s, "$") {
			arr[i] = strconv.Itoa(ids[idx])
			idx++
		}
	}

	u.Path = strings.Join(arr, "/")

	return u.String()
}
