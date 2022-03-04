package processor

import (
	"jsonParser/service"
)

type Worker struct {
	config       service.Config
	apiService   service.Api
	userService  service.User
	groupService service.Grouping
}

func (w *Worker) Initialize(config service.Config) *Worker {
	w.config = config

	return w
}

func (w *Worker) Start() error {
	apiService := service.Api{}
	userService := service.User{}
	groupService := service.Grouping{}

	err := apiService.Initialize(w.config.ApiUrl)
	if err != nil {
		return err
	}

	err = userService.Initialize(w.config, apiService)
	if err != nil {
		return err
	}

	err = groupService.Initialize(userService)
	if err != nil {
		return err
	}

	err = groupService.ExportToJson()
	if err != nil {
		return err
	}

	return nil
}

func (w *Worker) HasApiService() bool {
	if w.apiService.Ready {
		return true
	}
	return false
}

func (w *Worker) HasUserService() bool {
	if w.userService.Ready {
		return true
	}
	return false
}

func (w *Worker) HasGroupingService() bool {
	if w.groupService.Ready {
		return true
	}
	return false
}
