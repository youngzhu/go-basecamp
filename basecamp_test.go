package basecamp

import (
	"errors"
	"github.com/youngzhu/godate"
	"github.com/youngzhu/oauth2-apps/basecamp"
	"testing"
	"time"
)

func TestAccount(t *testing.T) {
	if _bc.accountID == "" {
		t.Error("account ID not set")
	}
	if _bc.accessToken == "" {
		t.Error("access token not set")
	}

	//println("accountID:", a.accountID)
	//println("accessToken:", a.accessToken)
}

func TestAddScheduleEntry_noProject(t *testing.T) {
	err := AddSchedule("nonproj", "", ScheduleEntry{})
	if !errors.Is(err, ErrNotFoundProject) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundProject, err)
	}
}

func TestAddScheduleEntry_noSchedule(t *testing.T) {
	proj := getRandProject()
	err := AddSchedule(proj.Name, "", ScheduleEntry{})
	if !errors.Is(err, ErrNotFoundSchedule) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundSchedule, err)
	}
}

func TestAddScheduleEntry_allDay(t *testing.T) {
	scheduleEntry := ScheduleEntry{
		Summary:  "开会一整天",
		AllDay:   true,
		StartsAt: time.Now(),
		EndsAt:   time.Now(),
	}

	err := AddSchedule("MeTime", "Schedule", scheduleEntry)
	if err != nil {
		t.Error(err)
	}
}

func TestAddScheduleEntry(t *testing.T) {
	scheduleEntry := ScheduleEntry{
		Summary: "Test",
		//StartsAt: time.Date(2023, 11, 5, 8, 0, 0, 0, time.Local),
		//EndsAt:   time.Date(2023, 11, 5, 10, 0, 0, 0, time.Local),
		StartsAt: time.Now(),
		EndsAt:   time.Now().Add(time.Minute * 30),
	}

	err := AddSchedule("MeTime", "Schedule", scheduleEntry)
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

func TestAddCard(t *testing.T) {
	today := godate.Today()
	workdays := today.Workdays()
	monday := workdays[0]
	friday := workdays[4]
	card := Card{
		Title: monday.String() + " ~ " + friday.String(),
		DueOn: friday.String(),
	}

	err := AddCard("Profession", "Card Table", "In progress", card)
	if err != nil {
		t.Error(err)
	}
}

func TestAddCard_simple(t *testing.T) {
	card := Card{
		Title: "Test Card",
	}

	err := AddCard("Profession", "Card Table", "In progress", card)
	if err != nil {
		t.Error(err)
	}
}

func TestAddTodo(t *testing.T) {
	todo := Todo{
		Content:     "test from API",
		Description: "测试",
		DueOn:       godate.Tomorrow().String(),
		StartsOn:    godate.Today().String(),
	}
	err := AddTodo("MeTime", "To-dos", "Inbox", todo)
	if err != nil {
		t.Error(err)
	}
}

func TestAddTodoList(t *testing.T) {
	AddTodoList("宝塔小学", "待办事项", "数学作业")
}

func TestAddTodoListAndTodos(t *testing.T) {
	AddTodoListAndTodos("宝塔小学", "待办事项",
		"数学作业：\n1、完成数学试卷一张；\n2、部分孩子订正数学补充习题第50--59页；\n3、读数学书第102--106页《期末复习》中的题目两遍")
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

func TestGetAccessToken(t *testing.T) {
	_, refresh := basecamp.GetAccessToken()
	if refresh != true {
		t.Error("should request a new access token by refresh")
	}
}
