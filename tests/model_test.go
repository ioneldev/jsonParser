package tests

import (
	"encoding/base64"
	"fmt"
	"jsonParser/model"
	"testing"
)

func TestUserGetIndex(t *testing.T) {
	GetIndexTests := []struct {
		user     model.User
		expected string
	}{
		{
			model.User{First: "First name"},
			"F",
		},
		{
			model.User{First: "Alex"},
			"A",
		},
		{
			model.User{First: "Paula"},
			"P",
		},
	}

	for _, tt := range GetIndexTests {
		testUser := tt.user
		if testUser.GetIndex() != tt.expected {
			t.Errorf("GetIndex(): expected %s, got %s, user %v", tt.expected, testUser.GetIndex(), testUser)
		}
	}
}

func TestUserGetUniqueId(t *testing.T) {
	mockUsers := GetSomeMockUsers()

	for _, tt := range mockUsers {
		testUser := tt
		expected := encodeUser(testUser)
		if testUser.GetUniqueId() != expected {
			t.Errorf("GetUniqueId(): expected %s, got %s, user %v", expected, testUser.GetUniqueId(), testUser)
		}
	}
}

func TestGroupAddRecordsAndJson(t *testing.T) {
	type expected struct {
		Type  string
		Value interface{}
	}

	AddRecordsTests := []struct {
		user     model.User
		expected expected
	}{
		{
			model.User{First: "First name", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"},
			expected{Type: "Value", Value: 1},
		},
		{
			model.User{First: "Ff", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"},
			expected{Type: "Value", Value: 2},
		},
		{
			model.User{First: "Fff", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"},
			expected{Type: "Value", Value: 3},
		},
		{
			model.User{First: "Eff", Last: "Last name", Email: "first.last@gmail.com", Address: "1st Street, Bla", Created: "25.02.2020", Balance: "120$"},
			expected{Type: "Error", Value: "invalid user being added to this Group"},
		},
	}

	group := model.Group{Index: "F"}

	if len(group.Records) != 0 {
		t.Errorf("TestAddRecords(): expected the len of records to be %d, got %d, group %v", 0, len(group.Records), group)
	}

	for _, tt := range AddRecordsTests {
		testUser := tt.user
		expected := tt.expected

		if expected.Type == "Value" {
			err := group.AddUser(testUser)

			if err != nil {
				t.Errorf("AddUser(): Got error instead of adding user %v, group %v", testUser, group)
			}

			if len(group.Records) != expected.Value {
				t.Errorf("AddUser(): expected records %s, got %d to group %v", expected.Value, len(group.Records), group)
			}

			group.Json()
			if group.NumberOfRecords != expected.Value {
				t.Errorf("AddUser(): expected records %s, got %d to group %v", expected.Value, group.NumberOfRecords, group)
			}
		}

		if expected.Type == "Error" {
			err := group.AddUser(testUser)

			if err == nil {
				t.Errorf("AddUser(): Didn't get error when adding user %s to group %v", testUser, group)
			}

			if err.Error() != expected.Value {

				t.Errorf("AddUser(): Didn't get expected error when adding user %s to group %v, got %v , expected %v", testUser, group, err.Error(), expected.Value)
			}
		}
	}
}

func encodeUser(u model.User) string {
	allUserFields := fmt.Sprintf("%s_%s_%s_%s_%s_%s", u.First, u.Last, u.Email, u.Address, u.Created, u.Balance)
	encodedText := base64.StdEncoding.EncodeToString([]byte(allUserFields))
	return encodedText
}
