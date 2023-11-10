package basecamp

import (
	"encoding/json"
	"net/http"
	"time"
)

type CardTable struct {
	Id               int64     `json:"id"`
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
	SubscriptionUrl  string    `json:"subscription_url"`
	Position         int       `json:"position"`
	Bucket           struct {
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
	Subscribers []struct {
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
	} `json:"subscribers"`
	Lists []CardColumn `json:"lists"`
}

type CardColumn struct {
	Id               int64     `json:"id"`
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
	SubscriptionUrl  string    `json:"subscription_url"`
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
	Description interface{} `json:"description"`
	Subscribers []struct {
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
	} `json:"subscribers"`
	Color        *string `json:"color"`
	CardsCount   int     `json:"cards_count"`
	CommentCount int     `json:"comment_count"`
	CardsUrl     string  `json:"cards_url"`
	Position     int     `json:"position,omitempty"`
}

type Card struct {
	Title string `json:"title"`
	DueOn string `json:"due_on"` // due date (ISO 8601) of the card, e.g. "2021-01-01"
}

/*
type Card struct {
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
	SubscriptionUrl  string    `json:"subscription_url"`
	CommentsCount    int       `json:"comments_count"`
	CommentsUrl      string    `json:"comments_url"`
	Position         int       `json:"position"`
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
	Description           string        `json:"description"`
	Completed             bool          `json:"completed"`
	Content               string        `json:"content"`
	DueOn                 interface{}   `json:"due_on"`
	Assignees             []interface{} `json:"assignees"`
	CompletionSubscribers []interface{} `json:"completion_subscribers"`
	CompletionUrl         string        `json:"completion_url"`
	CommentCount          int           `json:"comment_count"`
	Steps                 []struct {
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
		Completed bool   `json:"completed"`
		DueOn     string `json:"due_on"`
		Assignees []struct {
			Id                int         `json:"id"`
			AttachableSgid    string      `json:"attachable_sgid"`
			Name              string      `json:"name"`
			EmailAddress      string      `json:"email_address"`
			PersonableType    string      `json:"personable_type"`
			Title             string      `json:"title"`
			Bio               interface{} `json:"bio"`
			Location          interface{} `json:"location"`
			CreatedAt         time.Time   `json:"created_at"`
			UpdatedAt         time.Time   `json:"updated_at"`
			Admin             bool        `json:"admin"`
			Owner             bool        `json:"owner"`
			Client            bool        `json:"client"`
			Employee          bool        `json:"employee"`
			TimeZone          string      `json:"time_zone"`
			AvatarUrl         string      `json:"avatar_url"`
			CanManageProjects bool        `json:"can_manage_projects"`
			CanManagePeople   bool        `json:"can_manage_people"`
		} `json:"assignees"`
		CompletionUrl string `json:"completion_url"`
	} `json:"steps"`
}

*/

func (p *Project) getCardColumn(cardTableName, cardColumnName string) *CardColumn {
	d := p.getDock(cardTableName)
	if d == nil {
		return nil
	}

	// get card table
	resp, err := doRequest(d.Url, http.MethodGet, nil)
	if err != nil {
		return nil
	}

	var cardTable *CardTable
	err = json.Unmarshal(resp, &cardTable)
	if err != nil {
		return nil
	}

	// find the card column
	for _, c := range cardTable.Lists {
		if cardColumnName == c.Title {
			return &c
		}
	}
	return nil
}
