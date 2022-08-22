package jcosmos

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const (
	databsaseIDMaxLength   = 255
	ErrorDatabaseIDTooLong = "database id is toolong, id must be between 1 and 255 characters"
)

func (c Jcosmos) CreateDatabase(db newDatabaseRequest, obj interface{}) error {
	if len(db.ID) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	body, err := json.Marshal(db)
	if err != nil {
		return err
	}
	return c.cosmosRequest("/dbs", "", http.MethodPost, body, nil, obj)
}
func (c Jcosmos) CreateDatabaseWithThroughput(db newDatabaseRequest, throughput int, obj interface{}) error {
	if len(db.ID) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	h := map[string]string{
		"x-ms-offer-throughput": strconv.Itoa(throughput),
	}
	body, err := json.Marshal(db)
	if err != nil {
		return err
	}
	return c.cosmosRequest("/dbs", "", http.MethodPost, body, h, obj)
}
func (c Jcosmos) CreateDatabaseWithAutopilot(db newDatabaseRequest, max int, obj interface{}) error {
	if len(db.ID) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	h := map[string]string{
		"x-ms-cosmos-offer-autopilot-settings": fmt.Sprintf("{\"maxThroughput\":%d}", max),
	}
	body, err := json.Marshal(db)
	if err != nil {
		return err
	}
	return c.cosmosRequest("/dbs", "", http.MethodPost, body, h, obj)
}

func (c Jcosmos) ReadDatabase(db newDatabaseRequest, resp databaseResponse) error {
	if len(db.ID) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	return c.cosmosRequest("/dbs/"+c.db, "", http.MethodGet, emptyByteArr, nil, resp)
}
func (c Jcosmos) ListDatabase(resp databaseResponse) error {
	return c.cosmosRequest("/dbs", "", http.MethodGet, emptyByteArr, nil, resp)
}
func (c Jcosmos) DeleteDatabase(db newDatabaseRequest, resp databaseResponse) error {
	if len(db.ID) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	return c.cosmosRequest("/dbs", "", http.MethodDelete, emptyByteArr, nil, resp)
}
