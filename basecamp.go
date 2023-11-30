package basecamp

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"github.com/youngzhu/oauth2-apps/basecamp"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// API: https://github.com/basecamp/bc3-api

var a account

type account struct {
	accountID   string // Basecamp account ID
	accessToken string
}

func init() {
	viper.SetEnvPrefix("BASECAMP")
	viper.AutomaticEnv() // read in environment variables that match

	accessToken, refresh := basecamp.GetAccessToken()
	if refresh {
		log.Println("refresh token")
	}

	a = account{
		accountID:   viper.GetString("ACCOUNT_ID"),
		accessToken: accessToken,
	}
}

// AddScheduleEntry adds a schedule entry
// POST /buckets/1/schedules/3/entries.json
// creates a schedule entry in the project with ID 1 and under the schedule with an ID of 3.
func AddScheduleEntry(projectName, scheduleName string, scheduleEntry ScheduleEntry) error {
	project, err := GetProjectByName(projectName)
	if err != nil {
		return err
	}

	schedule := project.getDock(scheduleName)
	if schedule == nil {
		return fmt.Errorf("%w: %s", ErrNotFoundSchedule, scheduleName)
	}

	entryJson, _ := json.Marshal(scheduleEntry)

	url := parseUrl(urlScheduleEntry, project.Id, schedule.Id)
	_, err = doRequest(url, http.MethodPost, strings.NewReader(string(entryJson)))

	return err
}

// CreateCard creates a card
// POST /buckets/1/card_tables/lists/2/cards.json
// creates a card within the column with ID 2 in the project with id 1.
func CreateCard(projectName, cardTableName, columnName string, card Card) error {
	project, err := GetProjectByName(projectName)
	if err != nil {
		return err
	}

	cardColumn := project.getCardColumn(cardTableName, columnName)
	if cardColumn == nil {
		return fmt.Errorf("%w: card table: %q, card column: %q",
			ErrNotFoundCardColumn, cardTableName, columnName)
	}

	entryJson, _ := json.Marshal(card)

	url := parseUrl(urlCard, project.Id, cardColumn.Id)
	_, err = doRequest(url, http.MethodPost, strings.NewReader(string(entryJson)))

	return err
}

// CreateTodo creates a to-do
// POST /buckets/1/todolists/3/todos.json
// creates a to-do in the project with ID `1` and under the to-do list with an ID of `3`.
func CreateTodo(projectName, todoSetTitle, todoListTitle string, todo Todo) error {
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
	appUrl = strings.Replace(appUrl, "$ACCOUNT_ID", a.accountID, 1)

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
