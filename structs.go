package jcosmos

import (
	"encoding/json"
	"reflect"
)

// database
type newDatabaseRequest struct {
	ID string `json:"id"`
}
type databaseResponse struct {
	ID    string `json:"id"`
	Rid   string `json:"_rid"`
	TS    int64  `json:"_ts"`
	Self  string `json:"_self"`
	Etag  string `json:"_etag"`
	Colls string `json:"_colls"`
	Users string `json:"_users"`
}

// collection

// documents

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
type Query struct {
	Query      string      `json:"query"`
	Parameters []Parameter `json:"parameters,omitempty"`
}

func (q Query) ToJson() ([]byte, error) {
	return json.Marshal(q)

}

type QueryResponse struct {
	Documents []interface{} `json:"Documents"`
	Count     int           `json:"_count"`
	Rid       string        `json:"_rid"`
}

// stub to be added later
func (qr QueryResponse) ToStruct(t reflect.Type, obj interface{}) {}

// attachments
// sprocs
// functions
// trigers
// users
type newUserRequest struct {
	ID string `json:"id"`
}
type ListUserResponse struct {
	Users []UserResponse `json:"Users"`
	Count int            `json:"_count"`
	Rid   string         `json:"_rid"`
}
type UserResponse struct {
	ID          string `json:"id"`
	Rid         string `json:"_rid"`
	Ts          int64  `json:"_ts"`
	Self        string `json:"_self"`
	Etag        string `json:"_etag"`
	Permissions string `json:"_permissions"`
}

// permissions
// offers
