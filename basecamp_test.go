package basecamp

import (
	"errors"
	"testing"
)

func TestAccount(t *testing.T) {
	if a.accountID == "" {
		t.Error("account ID not set")
	}
	if a.accessToken == "" {
		t.Error("access token not set")
	}

	//println("accountID:", a.accountID)
	//println("accessToken:", a.accessToken)
}

func TestAddScheduleEntry_noProject(t *testing.T) {
	err := AddScheduleEntry("nonproj")
	if !errors.Is(err, ErrNotFoundProject) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundProject, err)
	}
}
