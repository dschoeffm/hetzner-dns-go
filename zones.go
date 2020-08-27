package hdns

import (
	"encoding/json"
	"errors"
)

type Zone struct {
	ID string `json:"id"`
	/* Hetzner uses a weird time format. Dont use time.Time here */
	Created         string   `json:"created"`
	Modified        string   `json:"modified"`
	LegacyDnsHost   string   `json:"legacy_dns_host"`
	LegacyNs        []string `json:"legacy_ns"`
	Name            string   `json:"name"`
	Ns              []string `json:"ns"`
	Owner           string   `json:"owner"`
	Paused          bool     `json:"paused"`
	Permission      string   `json:"permission"`
	Project         string   `json:"project"`
	Registrar       string   `json:"registrar"`
	Status          string   `json:"status"`
	Ttl             int
	Verified        string `json:"verified"`
	RecordsCount    int
	IsSecondaryDns  bool
	TxtVerification struct {
		Name  string `json:"name"`
		Token string `json:"token"`
	} `json:"txt_verification"`
}

type GetAllZonesResponse struct {
	Zones []Zone `json:"zones"`
	Meta  struct {
		Pagination Pagination `json:"pagination"`
	} `json:"meta"`
}

func GetAllZones(caller *apiCaller) ([]Zone, error) {
	respBody, code, err := caller.Call("GET", "zones", nil, nil)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, errors.New("GetAllZones failed! Status Code was not 200")
	}

	var respZones GetAllZonesResponse
	err = json.Unmarshal(respBody, &respZones)
	if err != nil {
		panic(err)
	}

	return respZones.Zones, err
}

func IdOfZone(zone string, zones []Zone) (string, error) {
	for _, z := range zones {
		if zone == z.Name {
			return z.ID, nil
		}
	}
	return "", errors.New("Zone not found")
}
