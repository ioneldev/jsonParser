package tests

import (
	"jsonParser/service"
	"testing"
)

func TestInitializeUsers(t *testing.T) {
	server := mockHttpServer()

	apiService := service.Api{}
	userService := service.User{}

	config := service.Config{ApiUrl: server.URL, ApiLinesToRead: 4}

	apiService.Initialize(server.URL)

	userService.Initialize(config, apiService)

}
