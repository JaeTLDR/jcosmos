package jcosmos

import (
	"encoding/json"
	"reflect"
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
