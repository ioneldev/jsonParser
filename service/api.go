package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Api struct {
	url   string
	Ready bool
}

func (api *Api) Initialize(url string) error {
	if url == "" {
		return errors.New("invalid API url")
	}

	api.url = url
	api.Ready = true
	return nil
}

func (api Api) checkResponseStatus(statusCode int) error {
	if statusCode < 200 || statusCode > 299 {
		return errors.New(fmt.Sprintf("Api responded with %d status code. Request not successfull", statusCode))
	}

	return nil
}

func (api Api) DoGetRequest() ([]byte, error) {
	resp, err := http.Get(api.url)

	if err != nil {
		return nil, errors.New("failed to do the http GET method")
	}

	err = api.checkResponseStatus(resp.StatusCode)

	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}
