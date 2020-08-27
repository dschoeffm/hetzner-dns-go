package hdns

import (
	"encoding/json"
	"errors"
)

type Record struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	/* Hetzner uses a weird time format. Dont use time.Time here */
	Created  string `json:"created"`
	Modified string `json:"modified"`
	ZoneId   string `json:"zone_id"`
	Name     string `json:"name"`
	Value    string `json:"value"`
	Ttl      int    `json:"ttl"`
}

type GetAllRecordsResponse struct {
	Records []Record `json:"records"`
	Meta    struct {
		Pagination Pagination `json:"pagination"`
	} `json:"meta"`
}

type CreateRecordRequest struct {
	Type   string `json:"type"`
	ZoneId string `json:"zone_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Ttl    int    `json:"ttl"`
}

type CreateRecordResponse struct {
	Record Record `json:"record"`
}

type UpdateRecordRequest struct {
	Type   string `json:"type"`
	ZoneId string `json:"zone_id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	Ttl    int    `json:"ttl"`
}

type UpdateRecordResponse struct {
	Record Record `json:"record"`
}

func GetAllRecords(caller *apiCaller, zoneId string) ([]Record, error) {
	resp, code, err := caller.Call("GET", "records", nil,
		map[string]string{"zone_id": zoneId})
	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, errors.New("GetAllRecords failed! Status code was not 200")
	}

	var respRecords GetAllRecordsResponse
	err = json.Unmarshal(resp, &respRecords)
	if err != nil {
		panic(err)
	}

	return respRecords.Records, nil
}

func FilterRecords(records []Record, name string, Type string) []Record {
	ret := make([]Record, 0)
	for _, r := range records {
		if r.Name == name && r.Type == Type {
			ret = append(ret, r)
		}
	}

	return ret
}

func UpdateRecord(caller *apiCaller, req *UpdateRecordRequest, recordId string) error {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	_, code, err := caller.Call("PUT", "records/"+recordId, b, nil)
	if err != nil {
		return nil
	}

	if code != 200 {
		return errors.New("UpdateRecord failed! Status code was not 200")
	}

	return nil
}

func CreateRecord(caller *apiCaller, req *CreateRecordRequest) error {
	b, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	_, code, err := caller.Call("POST", "records", b, nil)
	if err != nil {
		return nil
	}

	if code != 200 {
		return errors.New("CreateRecord failed! Status code was not 200")
	}

	return nil
}
