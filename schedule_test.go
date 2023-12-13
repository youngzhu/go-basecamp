package basecamp

import "testing"

func TestGetDock_schedule(t *testing.T) {
	scheduleTitle := "Schedule"
	d, _ := _bc.getDock("MeTime", TypeSchedule, scheduleTitle)

	scheduleDock := d.(*ScheduleDock)
	if scheduleTitle != scheduleDock.Title {
		t.Errorf("schedule title not match, want: %q, but got: %q",
			scheduleTitle, scheduleDock.Title)
	}
}
