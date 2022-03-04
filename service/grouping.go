package service

import (
	"jsonParser/model"
)

type Grouping struct {
	Ready       bool
	allGroups   []model.Group
	userService User
}

func (g *Grouping) Initialize(userService User) error {
	g.userService = userService

	internalGroupMap := make(map[string]*model.Group)

	err := g.buildInternalMap(internalGroupMap)
	if err != nil {
		return err
	}
	g.buildAllGroups(internalGroupMap)

	g.Ready = true

	return nil
}

func (g *Grouping) buildInternalMap(groupMap map[string]*model.Group) error {
	for _, user := range g.userService.GetDeduplicated() {
		userIndex := user.GetIndex()

		group, keyExists := groupMap[userIndex]

		if keyExists {
			err := group.AddUser(user)
			if err != nil {
				return err
			}
		} else {
			group = &model.Group{Index: userIndex}
			err := group.AddUser(user)
			if err != nil {
				return err
			}
			groupMap[userIndex] = group
		}
	}

	return nil
}

func (g *Grouping) buildAllGroups(groupMap map[string]*model.Group) {
	for _, group := range groupMap {
		g.allGroups = append(g.allGroups, *group)
	}
}

func (g *Grouping) ExportToJson() error {
	for _, group := range g.allGroups {
		encodedGroup, err := EncodeJson(group.Json())

		if err != nil {
			return err
		}

		filename := "export/" + group.Index + ".json"
		err = WriteJsonFile(encodedGroup, filename)
		if err != nil {
			return err
		}
	}

	return nil
}
