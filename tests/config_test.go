package tests

import (
	"jsonParser/service"
	"testing"
)

func TestInitializeConfig(t *testing.T) {
	config := service.Config{}

	if config.Ready {
		t.Errorf("TestInitializeConfig(): expected config to not be ready")
	}

	err := config.Initialize("mock_configg.json")
	if err.Error() != "failed to open config file" {
		t.Errorf("TestInitializeConfig(): expected error from init with invalid config file, got: %v", err.Error())
	}

	config.Initialize("mock_config.json")
	if config.ApiUrl != "" {
		t.Errorf("TestInitializeConfig(): struct is not nitialized, got: %v", config)
	}

	config.Initialize("../config.json")
	if config.ApiUrl == "" {
		t.Errorf("TestInitializeConfig(): struct is initialized, got: %v", config)
	}
}
