package basecamp

import (
	"encoding/json"
	"fmt"
)

type dockType string

const (
	TypeMessageBoard  dockType = "message_board"
	TypeTodoSet       dockType = "todoset"
	TypeVault         dockType = "vault"
	TypeChat          dockType = "chat"
	TypeSchedule      dockType = "schedule"
	TypeQuestionnaire dockType = "questionnaire"
	TypeInbox         dockType = "inbox"
	TypeKanbanBoard   dockType = "kanban_board"
	TypeCardTable              = TypeKanbanBoard
)

type dock struct {
	Id       int      `json:"id"`
	Title    string   `json:"title"`
	Name     dockType `json:"name"`
	Enabled  bool     `json:"enabled"`
	Position int      `json:"position"`
	Url      string   `json:"url"`
	AppUrl   string   `json:"app_url"`
}

type docker interface {
	DockType() dockType
	DockTitle() string
}

func (bc *BaseCamp) putDock(projectName string, adock docker) {
	key := buildDockKey(projectName, adock.DockType(), adock.DockTitle())
	bc.dockMap[key] = adock
}

func (bc *BaseCamp) getDock(projectName string, dt dockType, dockTitle string) (docker, error) {
	key := buildDockKey(projectName, dt, dockTitle)
	if v, ok := bc.dockMap[key]; ok {
		return v, nil
	}

	project, err := bc.GetProjectByName(projectName)
	if err != nil {
		return nil, err
	}

	var dockUrl string
	for _, d := range project.Dock {
		if dt == d.Name && dockTitle == d.Title {
			dockUrl = d.Url
		}
	}
	if dockUrl == "" {
		return nil, fmt.Errorf("%w: type: %s, title: %q", ErrNotFoundDock, dt, dockTitle)
	}

	resp, err := bc.doGet(dockUrl)
	if err != nil {
		return nil, err
	}

	var d docker

	switch dt {
	case TypeSchedule:
		var schedule *ScheduleDock
		err = json.Unmarshal(resp, &schedule)
		if err != nil {
			return nil, err
		}
		d = schedule
	case TypeCardTable:
		var cardTable *CardTableDock
		err = json.Unmarshal(resp, &cardTable)
		if err != nil {
			return nil, err
		}
		d = cardTable
	case TypeTodoSet:
		var todoSet *TodoSetDock
		err = json.Unmarshal(resp, &todoSet)
		if err != nil {
			return nil, err
		}
		d = todoSet
	default:
		return nil, ErrNotSupport
	}

	// cached
	bc.dockMap[key] = d

	return d, nil
}

func buildDockKey(projectName string, dt dockType, dockTitle string) string {
	const split = "##"
	return projectName + split + string(dt) + split + dockTitle
}

// todo remove
func (p *Project) getDock(dt dockType, dockTitle string) *dock {
	for _, d := range p.Dock {
		if dt == d.Name && dockTitle == d.Title {
			return &d
		}
	}
	return nil
}
