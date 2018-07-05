package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_helloWorld_Method_Get_Should_Be_hello_world(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	helloworld(w, req)
	expected := "hello world"

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if expected != string(body) {
		t.Errorf("should be %s but is %s", expected, body)
	}
}

func Test_helloWorld_Method_Post_Should_Be_statuscode_500(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	w := httptest.NewRecorder()
	helloworld(w, req)
	expectedStatusCode := 500
	expectedResponseText := "don't use post method"

	response := w.Result()
	actualResponseText, _ := ioutil.ReadAll(response.Body)

	if expectedStatusCode != response.StatusCode {
		t.Errorf("should be %d but is %d", expectedStatusCode, response.StatusCode)
	}
	if expectedResponseText != string(actualResponseText) {
		t.Errorf("should be \n%s' but is \n%s'", expectedResponseText, actualResponseText)
	}
}

func Test_helloWorld_Method_Get_Should_Be_hello_golf(t *testing.T) {
	req := httptest.NewRequest("GET", "/?text=golf", nil)
	w := httptest.NewRecorder()
	helloworld(w, req)
	expectedStatusCode := 200
	expectedResponseText := "hello golf"

	response := w.Result()
	actualResponseText, _ := ioutil.ReadAll(response.Body)

	if expectedStatusCode != response.StatusCode {
		t.Errorf("should be %d but is %d", expectedStatusCode, response.StatusCode)
	}
	if expectedResponseText != string(actualResponseText) {
		t.Errorf("should be \n%s' but is \n%s'", expectedResponseText, actualResponseText)
	}
}

func Test_HelloWithName_Method_Post_Should_Be_hello_Apipol_Sukgler(t *testing.T) {
	body := RequestBody{Id: "1", Name: "Apipol Sukgler"}
	bodyByte, _ := json.Marshal(body)

	req := httptest.NewRequest("POST", "/?text=golf", bytes.NewBuffer(bodyByte))
	w := httptest.NewRecorder()
	HelloWithName(w, req)
	expectedStatusCode := 200
	expectedResponseText := "hello Apipol Sukgler"

	response := w.Result()
	actualResponseText, _ := ioutil.ReadAll(response.Body)

	if expectedStatusCode != response.StatusCode {
		t.Errorf("should be %d but is %d", expectedStatusCode, response.StatusCode)
	}
	if expectedResponseText != string(actualResponseText) {
		t.Errorf("should be \n%s' but is \n%s'", expectedResponseText, actualResponseText)
	}
}
