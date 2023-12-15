package basecamp

import (
	"fmt"
	"time"
)

type CardTableDock struct {
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
	CardColumns []*CardColumn `json:"lists"`
}

func (d *CardTableDock) DockType() dockType {
	return TypeCardTable
}

func (d *CardTableDock) DockTitle() string {
	return d.Title
}

type CardColumn struct {
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

// AddCard creates a card
// POST /buckets/1/card_tables/lists/2/cards.json
// creates a card within the column with ID 2 in the project with id 1.
func (bc *BaseCamp) CreateCard(projectName, cardTableTitle, columnTitle string, card Card) error {
	cardTableDock, err := bc.GetCardTableDock(projectName, cardTableTitle)
	if err != nil {
		return err
	}

	cardColumn := cardTableDock.getCardColumn(columnTitle)
	if cardColumn == nil {
		return fmt.Errorf("%w: card table: %q, card column: %q",
			ErrNotFoundCardColumn, cardTableTitle, columnTitle)
	}

	_, err = bc.doPost(cardColumn.CardsUrl, card)

	return err
}

func (bc *BaseCamp) GetCardTableDock(projectName, cardTableTitle string) (*CardTableDock, error) {
	d, err := bc.getDock(projectName, TypeCardTable, cardTableTitle)
	if err != nil {
		return nil, err
	}

	return d.(*CardTableDock), nil
}

func (d *CardTableDock) getCardColumn(cardColumnTitle string) *CardColumn {
	for _, c := range d.CardColumns {
		if cardColumnTitle == c.Title {
			return c
		}
	}

	return nil
}
