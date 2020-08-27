package hdns

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Pagination struct {
	Page         int `json:"page"`
	PerPage      int `json:"per_page"`
	LastPage     int `json:"last_page"`
	TotalEntries int `json:"total_entries"`
}

type apiCaller struct {
	token string
}

func GetApiCaller(token string) *apiCaller {
	return &apiCaller{token: token}
}

func (caller *apiCaller) Call(method string, suffix string, body []byte,
	params map[string]string) ([]byte, int, error) {
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest(method, "https://dns.hetzner.com/api/v1/"+
		suffix, bytes.NewReader(body))

	// Headers
	req.Header.Add("Auth-API-Token", caller.token)

	// Add query params
	if params != nil {
		for param, val := range params {
			req.URL.Query().Add(param, val)
		}
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return respBody, resp.StatusCode, nil
}
