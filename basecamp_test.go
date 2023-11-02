package basecamp

import "testing"

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
