package basecamp

import (
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	atype := reflect.TypeOf(TypeKanbanBoard)

	want := "dockType"
	if atype.Name() != want {
		t.Errorf("want: %s, but got: %s", want, atype)
	}
}

func TestBuildDockKey(t *testing.T) {
	testcases := []struct {
		projectName string
		dt          dockType
		dockTitle   string
		want        string
	}{
		{"A Project", TypeSchedule, "a dock", "A Project##schedule##a dock"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			got := buildDockKey(tc.projectName, tc.dt, tc.dockTitle)
			if got != tc.want {
				t.Errorf("want: %q, but got: %q", tc.want, got)
			}
		})
	}
}

func TestGetDock_twice(t *testing.T) {
	scheduleTitle := "Schedule"
	projectName := "MeTime"

	// 1st request
	_bc.getDock(projectName, TypeSchedule, scheduleTitle)

	// 2nd request
	_bc.getDock(projectName, TypeSchedule, scheduleTitle)

}
