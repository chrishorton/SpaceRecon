package spacerecon

import (
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

const (
	uriBase          = "https://www.space-track.org"
	requestCmdAction = "/basicspacedata/query"
	requestLogin     = "/ajaxauth/login"
	requestTLE       = "/class/gp_history/format/tle/NORAD_CAT_ID/25544/orderby/EPOCH%20desc/limit/22"
)

type TLE struct {
	TLE string `json:"TLE"`
}

func GetTLE() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	username := viper.GetString("configuration.username")
	password := viper.GetString("configuration.password")

	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"identity": username,
			"password": password,
		}).
		Post(uriBase + requestLogin)

	if err != nil || resp.StatusCode() != 200 {
		log.Fatalf("Error logging in: %v, %s", err, resp.Status())
	}

	resp, err = client.R().
		SetResult(&[]TLE{}).
		Get(uriBase + requestCmdAction + requestTLE)

}
