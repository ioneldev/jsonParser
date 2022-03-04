package tests

import (
	"encoding/json"
	"jsonParser/model"
	"jsonParser/service"
	"testing"
)

func TestInitializeApi(t *testing.T) {
	api := service.Api{}

	err := api.Initialize("")

	if err.Error() != "invalid API url" {
		t.Errorf("TestInitializeApi(\"\"): expected error from init with empty string, got: %v", err.Error())
	}

	if api.Ready {
		t.Errorf("TestInitializeApi(\"\"): expected Ready to be false from init with empty string, got: %v", api.Ready)
	}

	err = api.Initialize("https://test.api.com")

	if err != nil {
		t.Errorf("TestInitializeApi(\"\"): didn't expect error from init with valid URL, got: %v", err.Error())
	}

	if !api.Ready {
		t.Errorf("TestInitializeApi(\"\"): expected Ready to be True from init with empty string, got: %v", api.Ready)
	}
}

func TestDoGetRequest(t *testing.T) {
	//server := mockHttpServer()
	//
	api := service.Api{}
	api.Initialize("htt://mockserver.com")

	_, err := api.DoGetRequest()
	if err.Error() != "failed to do the http GET method" {
		t.Errorf("TestDoGetRequest(): with wrong api url. Expected error(failed to do the http GET method), got %v", err)
	}

	server := mockHttpServer()

	api.Initialize(server.URL + "/error")

	_, err = api.DoGetRequest()
	if err.Error() != "Api responded with 404 status code. Request not successfull" {
		t.Errorf("TestDoGetRequest(): with valid api url. Not Successful. Expected error(Api responded with 400 status code. Request not successfull), got %v", err)
	}
	
	api.Initialize(server.URL)

	responseData, err := api.DoGetRequest()
	if err != nil {
		t.Errorf("TestDoGetRequest(): with valid api url. Successful. Did not expect error, got %v", err)
	}

	mockResponse := make([]model.User, 0)
	json.Unmarshal(responseData, &mockResponse)

	if mockResponse[0].First != "First name" {
		t.Errorf("TestDoGetRequest(): with valid api url. Successful. Failed to decode properly, got %v", err)
	}
}
