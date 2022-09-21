package jcosmos

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	userIDMaxLength    = 255
	ErrorUserIDTooLong = "user id is toolong, id must be between 1 and 255 characters"
)

func (c Jcosmos) CreateUser(user newUserRequest, obj UserResponse) error {
	if len(user.ID) > userIDMaxLength || len(user.ID) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}

	_, err = c.cosmosRequest("/dbs/"+c.db+"/users", "", http.MethodPost, body, nil, obj)
	return err
}

func (c Jcosmos) ReadUser(user newUserRequest, obj UserResponse) error {
	if len(user.ID) > userIDMaxLength || len(user.ID) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	_, err := c.cosmosRequest("/dbs/"+c.db+"/users/"+user.ID, "", http.MethodGet, emptyByteArr, nil, obj)
	return err
}

func (c Jcosmos) ListUser(obj ListUserResponse) error {
	_, err := c.cosmosRequest("/dbs/"+c.db+"/users", "", http.MethodGet, emptyByteArr, nil, obj)
	return err
}

func (c Jcosmos) UpdateUser(user newUserRequest, obj UserResponse) error {
	if len(user.ID) > userIDMaxLength || len(user.ID) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	body, err := json.Marshal(user)
	if err != nil {
		return err
	}
	_, err = c.cosmosRequest("/dbs/"+c.db+"/users/"+user.ID, "", http.MethodPut, body, nil, obj)
	return err
}

func (c Jcosmos) DeleteUser(user newUserRequest) error {
	if len(user.ID) > userIDMaxLength || len(user.ID) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	_, err := c.cosmosRequest("/dbs/"+c.db+"/users/"+user.ID, "", http.MethodDelete, emptyByteArr, nil, nil)
	return err
}
