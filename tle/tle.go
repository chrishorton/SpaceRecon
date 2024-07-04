package tle

import (
	"fmt"

	"github.com/chrishorton/spacerecon/config"
	"github.com/go-resty/resty/v2"
)

const (
    requestTLE       = "/class/gp_history/norad_cat_id/"
    tleParams        = "/CREATION_DATE/%3Enow-30/limit/5/format/json"

	requestIssTle    = "/class/gp_history/format/tle/NORAD_CAT_ID/25544/orderby/EPOCH%20desc/limit/22"
	uriBase          = "https://www.space-track.org"
	requestCmdAction = "/basicspacedata/query"
	requestLogin     = "/ajaxauth/login"
    configDir = ".."
)

type TLE struct {
	Ordinal             string `json:"ORDINAL"`
	Comment             string `json:"COMMENT"`
	Originator          string `json:"ORIGINATOR"`
	NoradCatID          string `json:"NORAD_CAT_ID"`
	ObjectName          string `json:"OBJECT_NAME"`
	ObjectType          string `json:"OBJECT_TYPE"`
	ClassificationType  string `json:"CLASSIFICATION_TYPE"`
	Intldes             string `json:"INTLDES"`
	Epoch               string `json:"EPOCH"`
	EpochMicroseconds   string `json:"EPOCH_MICROSECONDS"`
	MeanMotion          string `json:"MEAN_MOTION"`
	Eccentricity        string `json:"ECCENTRICITY"`
	Inclination         string `json:"INCLINATION"`
	RAOfAscNode         string `json:"RA_OF_ASC_NODE"`
	ArgOfPericenter     string `json:"ARG_OF_PERICENTER"`
	MeanAnomaly         string `json:"MEAN_ANOMALY"`
	EphemerisType       string `json:"EPHEMERIS_TYPE"`
	ElementSetNo        string `json:"ELEMENT_SET_NO"`
	RevAtEpoch          string `json:"REV_AT_EPOCH"`
	Bstar               string `json:"BSTAR"`
	MeanMotionDot       string `json:"MEAN_MOTION_DOT"`
	MeanMotionDdot      string `json:"MEAN_MOTION_DDOT"`
	File                string `json:"FILE"`
	TLELine0            string `json:"TLE_LINE0"`
	TLELine1            string `json:"TLE_LINE1"`
	TLELine2            string `json:"TLE_LINE2"`
	ObjectID            string `json:"OBJECT_ID"`
	ObjectNumber        string `json:"OBJECT_NUMBER"`
	SemimajorAxis       string `json:"SEMIMAJOR_AXIS"`
	Period              string `json:"PERIOD"`
	Apogee              string `json:"APOGEE"`
	Perigee             string `json:"PERIGEE"`
	Decayed             string `json:"DECAYED"`
}

func Get(noradId int) ([]TLE, error) {
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

    url := uriBase + requestCmdAction + requestTLE + fmt.Sprintf("%d", noradId) + tleParams

	resp, err = client.R().
		SetResult(&[]TLE{}).
		Get(url)

	if err != nil || resp.StatusCode() != 200 {
		return nil, fmt.Errorf("error fetching TLE: %v, %s", err, resp.Status())
	}

	return *resp.Result().(*[]TLE), nil
}

// Puts the data into a binary flatbuffer file
func SerializeTLE() {
    // TODO
}

func (t TLE) PrintTLE() {
	border := "====================================="
    fmt.Println()
	fmt.Println(border)
	fmt.Printf("TLE for %s (NORAD ID: %s)\n", t.ObjectName, t.NoradCatID)
	fmt.Println(border)
	fmt.Printf("%s\n", t.TLELine0)
	fmt.Printf("%s\n", t.TLELine1)
	fmt.Printf("%s\n", t.TLELine2)
	fmt.Println(border)
    fmt.Println()
}

func PrintTLEs(tles []TLE) {
    for _, tle := range tles {
        tle.PrintTLE()
    }
}


