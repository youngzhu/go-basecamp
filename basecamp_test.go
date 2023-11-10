package basecamp

import (
	"errors"
	"fmt"
	"github.com/youngzhu/go-basecamp/schedule"
	"github.com/youngzhu/godate"
	"net/http"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	if a.accountID == "" {
		t.Error("account ID not set")
	}
	if a.accessToken == "" {
		t.Error("access token not set")
	}

	//println("accountID:", a.accountID)
	//println("accessToken:", a.accessToken)
}

func TestAddScheduleEntry_noProject(t *testing.T) {
	err := AddScheduleEntry("nonproj", "", schedule.Entry{})
	if !errors.Is(err, ErrNotFoundProject) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundProject, err)
	}
}

func TestAddScheduleEntry_noSchedule(t *testing.T) {
	proj := getRandProject()
	err := AddScheduleEntry(proj.Name, "", schedule.Entry{})
	if !errors.Is(err, ErrNotFoundSchedule) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundSchedule, err)
	}
}

func TestAddScheduleEntry_allDay(t *testing.T) {
	scheduleEntry := schedule.Entry{
		Summary:  "开会一整天",
		AllDay:   true,
		StartsAt: time.Now(),
		EndsAt:   time.Now(),
	}

	err := AddScheduleEntry("MeTime", "Schedule", scheduleEntry)
	if err != nil {
		t.Error(err)
	}
}

func TestAddScheduleEntry(t *testing.T) {
	scheduleEntry := schedule.Entry{
		Summary:  "Test",
		StartsAt: time.Date(2023, 11, 5, 8, 0, 0, 0, time.Local),
		EndsAt:   time.Date(2023, 11, 5, 10, 0, 0, 0, time.Local),
	}

	err := AddScheduleEntry("MeTime", "Schedule", scheduleEntry)
	if err != nil {
		t.Error(err)
	}
}

//func TestGetCardTable(t *testing.T) {
//	_, err := GetCardTable("Profession", "Card Table")
//	if err != nil {
//		t.Error(err)
//	}
//}

func TestCreateCard(t *testing.T) {
	today := godate.Today()
	workdays := today.Workdays()
	monday := workdays[0]
	friday := workdays[4]
	card := Card{
		Title: monday.String() + " ~ " + friday.String(),
		DueOn: friday.String(),
	}

	err := CreateCard("Profession", "Card Table", "In progress", card)
	if err != nil {
		t.Error(err)
	}
}

func TestParseUrl(t *testing.T) {
	testcases := []struct {
		url    string
		params []int
		want   string
	}{
		{"http://youngzy.com", nil, "http://youngzy.com"},
		{"http://youngzy.com/x/y/z", nil, "http://youngzy.com/x/y/z"},
		{"http://youngzy.com/x/y/z", []int{123, 456}, "http://youngzy.com/x/y/z"},
		{"http://youngzy.com/$1/y/$2", []int{4546, 7568679}, "http://youngzy.com/4546/y/7568679"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			got := parseUrl(tc.url, tc.params...)
			if got != tc.want {
				t.Errorf("want: %q, but got: %q", tc.want, got)
			}
		})
	}

}

func TestDoRequest(t *testing.T) {
	url := "https://3.basecampapi.com/5161745/buckets/28933112/card_tables/5878088123.json"
	jsonResp, err := doRequest(url, http.MethodGet, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Responese: %s", jsonResp)
}
