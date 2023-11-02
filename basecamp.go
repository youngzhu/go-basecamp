package basecamp

import "github.com/spf13/viper"

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
