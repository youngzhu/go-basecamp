package basecamp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Project struct {
	Id             int       `json:"id"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Purpose        string    `json:"purpose"`
	ClientsEnabled bool      `json:"clients_enabled"`
	BookmarkUrl    string    `json:"bookmark_url"`
	Url            string    `json:"url"`
	AppUrl         string    `json:"app_url"`
	Dock           []struct {
		Id       int    `json:"id"`
		Title    string `json:"title"`
		Name     string `json:"name"`
		Enabled  bool   `json:"enabled"`
		Position *int   `json:"position"`
		Url      string `json:"url"`
		AppUrl   string `json:"app_url"`
	} `json:"dock"`
}

func GetProjects() ([]Project, error) {
	url := parseUrl(UrlProjects)
	jsonProjects, err := doRequest(url, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = json.Unmarshal(jsonProjects, &projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

var ErrNotFoundProject = errors.New("not found the project")

func GetProjectByName(name string) (*Project, error) {
	projects, err := GetProjects()
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if name == project.Name {
			return &project, nil
		}
	}

	return nil, fmt.Errorf("%w: %s", ErrNotFoundProject, name)
}
