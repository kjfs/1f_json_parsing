package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	// PreRun
	testData := `{"status":"ok"}`
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(testData))
	}))
	apiClient := NewApiClient(svr.URL, 10*time.Second)
	resBytes, err := apiClient.GetJson()
	if err != nil {
		t.Errorf("expected err to be nil, but got %v", err)
	}
	apiClient.Unmarshall(resBytes)
	result := string(resBytes)
	result = strings.TrimSpace(string(result))
	if result != testData {
		t.Errorf("expected result to be: '%s', but got: '%s'", testData, result)
	}
	svr.Close()

	// Test1: Value of Status is 'invalid' instead of 'ok' in JSON body
	testData = `{"status":"invalid"}`
	svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(testData))
	}))
	apiClient = NewApiClient(svr.URL, 10*time.Second)
	resBytes, err = apiClient.GetJson()
	if err != nil {
		t.Errorf("expected err to be nil, but got %v", err)
	}
	apiClient.Unmarshall(resBytes)
	result = string(resBytes)
	result = strings.TrimSpace(string(result))
	if result != testData {
		t.Errorf("expected result to nil: '%s', but got: '%s'", testData, result)
	}
	svr.Close()

	// Test2: Unmarshalling should fail
	testData = `{"status":"invalid"}`
	svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(nil))
	}))
	apiClient = NewApiClient(svr.URL, 10*time.Second)
	resBytes, err = apiClient.GetJson()
	if err == nil {
		t.Errorf("Err should ne nil, but got: %v", err)
	}
	apiClient.Unmarshall(resBytes)
	result = string(resBytes)
	result = strings.TrimSpace(string(result))
	if result == testData {
		t.Errorf("expected result to be nil '%s', but got: '%s'", testData, result)
	}
	svr.Close()

	// Test3: Server is closed before GetJson is called.
	testData = `{"status":"invalid"}`
	svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	svr.Close()
	apiClient = NewApiClient(svr.URL, 10*time.Second)
	resBytes, err = apiClient.GetJson()
	if err == nil {
		t.Errorf("server should be closed, but got %v", err)
	}
	apiClient.Unmarshall(resBytes)
	result = string(resBytes)
	result = strings.TrimSpace(string(result))
	if result == testData {
		t.Errorf("expected result to be nil '%s', but got: '%s'", testData, result)
	}
	svr.Close()

	// Test4: ioutil.Readall should fail.
	testData = `{"status":"invalid"}`
	svr = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))
	apiClient = NewApiClient(svr.URL, 1*time.Millisecond)
	resBytes, err = apiClient.GetJson()
	if err == nil {
		t.Errorf("err should be nil, but got: %v", err)
	}
	apiClient.Unmarshall(resBytes)
	result = string(resBytes)
	result = strings.TrimSpace(string(result))
	if result == testData {
		t.Errorf("expected result to be nil '%s', but got: '%s'", testData, result)
	}
	svr.Close()
}
