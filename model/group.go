package model

import "errors"

type Group struct {
	Index           string
	Records         []User
	NumberOfRecords int
}

func (g *Group) AddUser(user User) error {
	if user.GetIndex() != g.Index {
		return errors.New("invalid user being added to this Group")
	}

	g.Records = append(g.Records, user)
	return nil
}

func (g *Group) Json() Group {
	g.NumberOfRecords = len(g.Records)
	return *g
}
