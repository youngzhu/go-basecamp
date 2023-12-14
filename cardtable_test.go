package basecamp

import "testing"

func TestBaseCamp_GetCardTableDock(t *testing.T) {
	cardTableTitle := "Card Table"
	cardTableDock, err := _bc.GetCardTableDock("Profession", cardTableTitle)
	if err != nil {
		t.Error(err)
	}
	if cardTableDock.Title != cardTableTitle {
		t.Errorf("card table title not match, want: %s, but got: %s",
			cardTableTitle, cardTableDock.Title)
	}
}
