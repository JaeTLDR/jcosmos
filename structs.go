package jcosmos

import (
	"encoding/json"
	"errors"
	"strings"
)

var emptyByteArr = []byte{}

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
	Documents []any  `json:"Documents"`
	Count     int    `json:"_count"`
	Rid       string `json:"_rid"`
}

type PatchOp string

const (
	PatchOpAdd       PatchOp = "add"
	PatchOpSet       PatchOp = "set"
	PatchOpReplace   PatchOp = "replace"
	PatchOpRemove    PatchOp = "remove"
	PatchOpIncrement PatchOp = "increment"
)

var allowedPatchOperations []PatchOp = []PatchOp{
	PatchOpAdd,
	PatchOpSet,
	PatchOpReplace,
	PatchOpRemove,
	PatchOpIncrement,
}

type Patch struct {
	Condition  string           `json:"condition,omitempty"`
	Operations []PatchOperation `json:"operations"`
}
type PatchOperation struct {
	Op    PatchOp `json:"op"`
	Path  string  `json:"path"`
	Value any     `json:"value"`
}

func (po PatchOperation) validate() error {
	if inArray(allowedPatchOperations, po.Op) {
		return errors.New("invalid operation")
	}
	if !strings.HasPrefix(po.Path, "/") {
		return errors.New("path does nto start at root of document")
	}
	return nil
}

// stub to be added later
// func (qr QueryResponse) ToStruct(obj any) {
// 	t := reflect.TypeOf(obj)
// 	tArr := reflect.MakeSlice(t, 0, 1000)

// 	for _, i := range qr.Documents {
// 		tArr = reflect.Append(tArr, i)
// 	}
// 	obj = tArr
// }

// attachments
// sprocs
// functions
// trigers
type Trigger struct {
	ID               string `json:"id"`
	Body             string `json:"body"`
	TriggerOperation string `json:"triggerOperation"`
	TriggerType      string `json:"triggerType"`
}
type TriggerResponse struct {
	ID               string `json:"id"`
	Body             string `json:"body"`
	TriggerOperation string `json:"triggerOperation"`
	TriggerType      string `json:"triggerType"`
	Rid              string `json:"_rid"`
	Ts               int64  `json:"_ts"`
	Self             string `json:"_self"`
	Etag             string `json:"_etag"`
}
type TriggerList struct {
	Triggers []TriggerResponse `json:"Users"`
	Count    int               `json:"_count"`
	Rid      string            `json:"_rid"`
}

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
type OfferRequest struct {
	OfferVersion    string       `json:"offerVersion"`
	Content         OfferContent `json:"content"`
	Resource        string       `json:"resource"`
	OfferResourceID string       `json:"offerResourceID"`
	ID              string       `json:"id"`
	Rid             string       `json:"_rid"`
}
type OfferResponse struct {
	OfferVersion    string       `json:"offerVersion"`
	OfferType       string       `json:"offerType"`
	Content         OfferContent `json:"content"`
	Resource        string       `json:"resource"`
	OfferResourceID string       `json:"offerResourceID"`
	ID              string       `json:"id"`
	Rid             string       `json:"_rid"`
	Self            string       `json:"_self"`
	Etag            string       `json:"_etag"`
	Ts              int64        `json:"_ts"`
}
type OfferContent struct {
	OfferThroughput int `json:"offerThroughput"`
}

type ListOfferResponse struct {
	Offers []OfferResponse `json:"Offers"`
	Count  int             `json:"_count"`
	Rid    string          `json:"_rid"`
}
