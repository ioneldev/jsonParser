package service

import (
	"errors"
	"jsonParser/model"
)

type User struct {
	config       Config
	Ready        bool
	all          []model.User
	deduplicated []model.User
}

func (u *User) Initialize(config Config, apiService Api) error {
	if !config.Ready {
		return errors.New("config service is not initialized")
	}

	if !apiService.Ready {
		return errors.New("api service is not initialized")
	}

	u.config = config

	for len(u.all) < u.config.ApiLinesToRead {
		bodyBytes, err := apiService.DoGetRequest()

		if err != nil {
			return err
		}

		tempUsers, err := u.processApiResponse(bodyBytes)

		if len(u.all)+len(tempUsers) > u.config.ApiLinesToRead {
			maxTempUsers := u.config.ApiLinesToRead - len(u.all)
			tempUsers = tempUsers[0:maxTempUsers]
		}

		if err != nil {
			return err
		}

		u.all = append(u.all, tempUsers...)
	}

	u.buildDeduplicated()

	u.Ready = true

	return nil
}

func (u *User) GetDeduplicated() []model.User {
	return u.deduplicated
}

func (u *User) buildDeduplicated() {
	var dedupeMap = make(map[string]model.User)

	for _, user := range u.all {
		userId := user.GetUniqueId()
		dedupeMap[userId] = user
	}

	for _, value := range dedupeMap {
		u.deduplicated = append(u.deduplicated, value)
	}
}

func (u *User) processApiResponse(bytes []byte) ([]model.User, error) {
	var objects []model.User

	err := LoadObjectFromJson(bytes, &objects)
	if err != nil {
		return nil, err
	}

	return objects, nil
}
