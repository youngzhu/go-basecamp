package basecamp

import (
	"encoding/json"
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
	Dock           []dock    `json:"dock"`
}

func (bc *BaseCamp) GetProjects() ([]Project, error) {
	if bc.projects == nil {
		const urlProjects = "https://3.basecampapi.com/$ACCOUNT_ID/projects.json"

		bc.projectsUrl = parseUrl(urlProjects)
		jsonProjects, err := bc.doRequest(bc.projectsUrl, http.MethodGet, nil)
		if err != nil {
			return nil, err
		}

		var projects []Project
		err = json.Unmarshal(jsonProjects, &projects)
		if err != nil {
			return nil, err
		}

		bc.projects = projects
	}

	return bc.projects, nil
}

func GetProjects() ([]Project, error) {
	return _bc.GetProjects()
}

func (bc *BaseCamp) GetProjectByName(name string) (*Project, error) {
	projects, err := bc.GetProjects()
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

func GetProjectByName(name string) (*Project, error) {
	return _bc.GetProjectByName(name)
}
