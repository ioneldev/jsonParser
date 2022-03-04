package tests

import (
	"encoding/json"
	"jsonParser/model"
	"net/http"
	"net/http/httptest"
	"strings"
)

func GetSomeMockUsers() []model.User {
	users := make([]model.User, 0, 3)

	users = append(users, model.User{First: "First name", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"})
	users = append(users, model.User{First: "Alex", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"})
	users = append(users, model.User{First: "Paula", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"})
	users = append(users, model.User{First: "Paula", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"})

	return users
}

func mockHttpServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch strings.TrimSpace(r.URL.Path) {
		case "/":
			mockFetchDataEndpointSuccess(w)
		default:
			http.NotFoundHandler().ServeHTTP(w, r)
		}
	}))

	return server
}

func mockFetchDataEndpointSuccess(w http.ResponseWriter) {
	statusCode := http.StatusOK
	response := GetSomeMockUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
