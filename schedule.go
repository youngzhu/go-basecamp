package basecamp

import (
	"encoding/json"
	"net/http"
	"time"
)

type ScheduleEntry struct {
	// for create
	Summary  string    `json:"summary"`
	AllDay   bool      `json:"all_day"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`
}

/*
type ScheduleEntry struct {
	// when create
	Summary  string    `json:"summary"`
	AllDay   bool      `json:"all_day"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt   time.Time `json:"ends_at"`

	Id               int       `json:"id"`
	Status           string    `json:"status"`
	VisibleToClients bool      `json:"visible_to_clients"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Title            string    `json:"title"`
	Description      string    `json:"description"`
	InheritsStatus   bool      `json:"inherits_status"`
	Type             string    `json:"type"`
	Url              string    `json:"url"`
	AppUrl           string    `json:"app_url"`
	BookmarkUrl      string    `json:"bookmark_url"`
	SubscriptionUrl  string    `json:"subscription_url"`
	CommentsCount    int       `json:"comments_count"`
	CommentsUrl      string    `json:"comments_url"`
	Parent           struct {
		Id     int    `json:"id"`
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
	Participants []struct {
		Id             int         `json:"id"`
		AttachableSgid string      `json:"attachable_sgid"`
		Name           string      `json:"name"`
		EmailAddress   string      `json:"email_address"`
		PersonableType string      `json:"personable_type"`
		Title          string      `json:"title"`
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
		CanManageProjects bool `json:"can_manage_projects"`
		CanManagePeople   bool `json:"can_manage_people"`
	} `json:"participants"`
}

*/

type Schedule struct {
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
	IncludeDueAssignments bool   `json:"include_due_assignments"`
	EntriesCount          int    `json:"entries_count"`
	EntriesUrl            string `json:"entries_url"`
}

func (p *Project) getSchedule(scheduleTitle string) *Schedule {
	d := p.getDock(TypeSchedule, scheduleTitle)
	if d == nil {
		return nil
	}

	resp, err := doRequest(d.Url, http.MethodGet, nil)
	if err != nil {
		return nil
	}

	var schedule *Schedule
	err = json.Unmarshal(resp, &schedule)
	if err != nil {
		return nil
	}

	return schedule
}
