package jcosmos

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	userIDMaxLength    = 255
	ErrorUserIDTooLong = "user id is toolong, id must be between 1 and 255 characters"
)

func (c Jcosmos) CreateUser(id string, user UserResponse) error {
	if len(id) > userIDMaxLength || len(id) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	body := fmt.Sprintf("{\"id\":\"%s\"}", id)

	return c.cosmosRequest("/dbs/"+c.db+"/users", "", http.MethodPost, body, nil, user)
}

func (c Jcosmos) ReadUser(id string, user UserResponse) error {
	if len(id) > userIDMaxLength || len(id) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	return c.cosmosRequest("/dbs/"+c.db+"/users/"+id, "", http.MethodGet, "", nil, user)
}

func (c Jcosmos) ListUser(users ListUserResponse) error {
	return c.cosmosRequest("/dbs/"+c.db+"/users", "", http.MethodGet, "", nil, users)
}

func (c Jcosmos) UpdateUser(id string, user UserResponse) error {
	if len(id) > userIDMaxLength || len(id) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	body := fmt.Sprintf("{\"id\":\"%s\"}", id)
	return c.cosmosRequest("/dbs/"+c.db+"/users/"+id, "", http.MethodPut, body, nil, user)
}

func (c Jcosmos) DeleteUser(id string) error {
	if len(id) > userIDMaxLength || len(id) < 1 {
		return errors.New(ErrorUserIDTooLong)
	}
	return c.cosmosRequest("/dbs/"+c.db+"/users/"+id, "", http.MethodDelete, "", nil, nil)
}
