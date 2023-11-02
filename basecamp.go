package basecamp

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

var a account

type account struct {
	accountID   string // Basecamp account ID
	accessToken string
}

func init() {
	viper.SetEnvPrefix("BASECAMP")
	viper.AutomaticEnv() // read in environment variables that match

	a = account{
		accountID:   viper.GetString("ACCOUNT_ID"),
		accessToken: viper.GetString("ACCESS_TOKEN"),
	}
}

const (
	UrlProjects = "https://3.basecampapi.com/$ACCOUNT_ID/projects.json"
)

func parseUrl(url string, ids ...string) string {
	url = strings.Replace(url, "$ACCOUNT_ID", a.accountID, 1)

	return url
}

func AddScheduleEntry(projectName string) error {
	project, err := GetProjectByName(projectName)
	if err != nil {
		return err
	}

	fmt.Println("project ID:", project.Id)

	return nil
}
