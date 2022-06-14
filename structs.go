package jcosmos

import (
	"encoding/json"
)

type Query struct {
	Query      string      `json:"query"`
	Parameters []Parameter `json:"parameters,omitempty"`
}
type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (q Query) ToJson() ([]byte, error) {
	return json.Marshal(q)

}
