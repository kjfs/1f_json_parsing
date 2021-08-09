package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ApiClient is a struct for the api clients, which will be used to call api server.
type ApiClient struct {
	URL        string
	httpClient *http.Client
}

// function NewApiClient creates an api client.
func NewApiClient(url string, timeout time.Duration) ApiClient {
	return ApiClient{
		URL: url,
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Method "GetJson" does three things:
// 1. calls server API,
// 2. reads JSON response body,
// 3. returns []bytes of JSON response.
func (apiClient ApiClient) GetJson() ([]byte, error) {
	resp, err := http.Get(apiClient.URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil || len(respBytes) == 0 {
		return nil, fmt.Errorf("error with response body: %s", err)
	}
	return respBytes, nil
}

// Method "Unmarshall" does two things:
// 1. unmarshalls the byte data (JSON-encoded data),
// 2. stores the result into a variable pointed by '&tmp'.
func (apiClient ApiClient) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return fmt.Errorf("error with unmarshalling: %s", err)
	}
	return nil
}
