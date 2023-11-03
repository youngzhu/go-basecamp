package basecamp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/youngzhu/go-basecamp/schedule"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var a account

type account struct {
	accountID   string // Basecamp account ID
	accessToken string
}

func init() {
	viper.SetEnvPrefix("BASECAMP")
	viper.AutomaticEnv() // read in environment variables that match

	a = account{
		accountID:   viper.GetString("ACCOUNT_ID"),
		accessToken: viper.GetString("ACCESS_TOKEN"),
	}
}

const (
	UrlProjects      = "https://3.basecampapi.com/$ACCOUNT_ID/projects.json"
	UrlScheduleEntry = "https://3.basecampapi.com/$ACCOUNT_ID/buckets/$1/schedules/$3/entries.json"
)

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

var (
	ErrNotFoundProject  = errors.New("not found the project")
	ErrNotFoundSchedule = errors.New("not found the schedule")
)

// AddScheduleEntry adds a schedule entry
// POST /buckets/1/schedules/3/entries.json
// creates a schedule entry in the project with ID 1 and under the schedule with an ID of 3.
func AddScheduleEntry(projectName, scheduleName string, scheduleEntry schedule.Entry) error {
	project, err := GetProjectByName(projectName)
	if err != nil {
		return err
	}

	var scheduleId int
	for _, dock := range project.Dock {
		if scheduleName == dock.Title {
			scheduleId = dock.Id
		}
	}
	if scheduleId == 0 {
		return fmt.Errorf("%w: %s", ErrNotFoundSchedule, scheduleName)
	}

	entryJson, _ := json.Marshal(scheduleEntry)

	url := parseUrl(UrlScheduleEntry, project.Id, scheduleId)
	_, err = doRequest(url, http.MethodPost, strings.NewReader(string(entryJson)))

	return err
}
