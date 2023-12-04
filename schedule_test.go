package basecamp

import "testing"

func TestProject_getSchedule(t *testing.T) {
	project, _ := GetProjectByName("MeTime")
	scheduleTitle := "Schedule"
	schedule := project.getSchedule(scheduleTitle)

	if scheduleTitle != schedule.Title {
		t.Errorf("schedule title not match, want: %q, but got: %q",
			scheduleTitle, schedule.Title)
	}
}
