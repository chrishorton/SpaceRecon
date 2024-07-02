package tle

import (
	"fmt"

	"github.com/chrishorton/spacerecon/config"
	"github.com/go-resty/resty/v2"
)

const (
	requestTLE = "/class/gp_history/format/tle/NORAD_CAT_ID/25544/orderby/EPOCH%20desc/limit/22"
	uriBase          = "https://www.space-track.org"
	requestCmdAction = "/basicspacedata/query"
	requestLogin     = "/ajaxauth/login"
    configDir = ".."
)

type TLE struct {
	TLE string `json:"TLE"`
}

func Get() ([]TLE, error) {
	username := config.Cfg.Username
	password := config.Cfg.Password

	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"identity": username,
			"password": password,
		}).
		Post(uriBase + requestLogin)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error logging in: %v, %s", err, resp.Status())
	}

	resp, err = client.R().
		SetResult(&[]TLE{}).
		Get(uriBase + requestCmdAction + requestTLE)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error fetching TLE: %v, %s", err, resp.Status())
	}

    fmt.Println(resp.Result().(*[]TLE))

	return *resp.Result().(*[]TLE), nil
}
