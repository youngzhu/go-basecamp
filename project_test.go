package basecamp

import (
	"fmt"
	"testing"
)

func TestGetProjects(t *testing.T) {
	projects, err := GetProjects()
	if err != nil {
		t.Error(err)
	}
	if len(projects) == 0 {
		t.Error("there's no project")
	}

	for i, project := range projects {
		fmt.Printf("project[%d] name: %s\n", i, project.Name)
	}
}
