package spacerecon

import (
	"log"

	"github.com/chrishorton/spacerecon/config"
	"github.com/go-resty/resty/v2"
)

const (
	requestFindStarlinks = "/class/tle_latest/NORAD_CAT_ID/>40000/ORDINAL/1/OBJECT_NAME/STARLINK~~/format/json/orderby/NORAD_CAT_ID%20asc"

	requestOMMStarlink1 = "/class/omm/NORAD_CAT_ID/"
	requestOMMStarlink2 = "/orderby/EPOCH%20asc/format/json"
	GM                  = 398600441800000.0
	MRAD                = 6378.137
	PI                  = 3.14159265358979
	TPI86               = 2.0 * PI / 86400.0
)

type Satellite struct {
	NORAD_CAT_ID      string `json:"NORAD_CAT_ID"`
	OBJECT_NAME       string `json:"OBJECT_NAME"`
	EPOCH             string `json:"EPOCH"`
	REV_AT_EPOCH      string `json:"REV_AT_EPOCH"`
	INCLINATION       string `json:"INCLINATION"`
	ECCENTRICITY      string `json:"ECCENTRICITY"`
	MEAN_MOTION       string `json:"MEAN_MOTION"`
	RA_OF_ASC_NODE    string `json:"RA_OF_ASC_NODE"`
	ARG_OF_PERICENTER string `json:"ARG_OF_PERICENTER"`
	MEAN_ANOMALY      string `json:"MEAN_ANOMALY"`
}

func GetStarlinkSatellites() ([]Satellite, error) {
	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"identity": config.Cfg.Username,
			"password": config.Cfg.Password,
		}).
		Post(uriBase + requestLogin)

	if err != nil || resp.StatusCode() != 200 {
		log.Fatalf("Error logging in: %v, %s", err, resp.Status())
	}

	resp, err = client.R().
		SetResult(&[]Satellite{}).
		Get(uriBase + requestCmdAction + requestFindStarlinks)

	if err != nil || resp.StatusCode() != 200 {
		return nil, err
	}

	return *resp.Result().(*[]Satellite), nil
}
