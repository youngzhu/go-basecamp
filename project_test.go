package basecamp

import (
	"errors"
	"fmt"
	"math/rand"
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

func TestGetProjectByName_notFound(t *testing.T) {
	_, err := GetProjectByName("test123")
	if !errors.Is(err, ErrNotFoundProject) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundProject, err)
	}
}

func TestGetProjectByName(t *testing.T) {
	p1 := getRandProject()

	p2, err := GetProjectByName(p1.Name)
	fmt.Println("project name:", p1.Name)
	if err != nil {
		t.Error(err)
	}
	if p1.Name != p2.Name {
		t.Errorf("%q and %q should be the same", p1.Name, p2.Name)
	}
}

func getRandProject() *Project {
	projects, _ := GetProjects()

	size := len(projects)
	return &projects[rand.Intn(size)]
}
