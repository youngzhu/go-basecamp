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
	err := AddScheduleEntry("nonproj", "")
	if !errors.Is(err, ErrNotFoundProject) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundProject, err)
	}
}

func TestAddScheduleEntry_noSchedule(t *testing.T) {
	proj := getRandProject()
	err := AddScheduleEntry(proj.Name, "")
	if !errors.Is(err, ErrNotFoundSchedule) {
		t.Errorf("Expected error %q, got %q instead", ErrNotFoundSchedule, err)
	}
}

func TestAddScheduleEntry(t *testing.T) {
	err := AddScheduleEntry("MeTime", "Schedule")
	if err != nil {
		t.Error(err)
	}
}

func TestParseUrl(t *testing.T) {
	testcases := []struct {
		url    string
		params []int
		want   string
	}{
		{"http://youngzy.com", nil, "http://youngzy.com"},
		{"http://youngzy.com/x/y/z", nil, "http://youngzy.com/x/y/z"},
		{"http://youngzy.com/x/y/z", []int{123, 456}, "http://youngzy.com/x/y/z"},
		{"http://youngzy.com/$1/y/$2", []int{4546, 7568679}, "http://youngzy.com/4546/y/7568679"},
	}

	for _, tc := range testcases {
		t.Run("", func(t *testing.T) {
			got := parseUrl(tc.url, tc.params...)
			if got != tc.want {
				t.Errorf("want: %q, but got: %q", tc.want, got)
			}
		})
	}

}
