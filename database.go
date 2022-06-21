package jcosmos

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const (
	databsaseIDMaxLength   = 255
	ErrorDatabaseIDTooLong = "database id is toolong, id must be between 1 and 255 characters"
)

func (c Jcosmos) CreateDatabase(id string, obj interface{}) error {
	if len(id) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	body := fmt.Sprintf("{\"id\":\"%s\"}", id)
	return c.cosmosRequest("/dbs", "", http.MethodPost, body, nil, obj)
}
func (c Jcosmos) CreateDatabaseWithThroughput(id string, throughput int, obj interface{}) error {
	if len(id) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	h := map[string]string{
		"x-ms-offer-throughput": strconv.Itoa(throughput),
	}
	body := fmt.Sprintf("{\"id\":\"%s\"}", id)
	return c.cosmosRequest("/dbs", "", http.MethodPost, body, h, obj)
}
func (c Jcosmos) CreateDatabaseWithAutopilot(id string, max int, obj interface{}) error {
	if len(id) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	h := map[string]string{
		"x-ms-cosmos-offer-autopilot-settings": fmt.Sprintf("{\"maxThroughput\":%d}", max),
	}
	body := fmt.Sprintf("{\"id\":\"%s\"}", id)
	return c.cosmosRequest("/dbs", "", http.MethodPost, body, h, obj)
}

func (c Jcosmos) ReadDatabase(id string, resp databaseResponse) error {
	if len(id) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	return c.cosmosRequest("/dbs/"+c.db, "", http.MethodPost, "", nil, resp)
}
func (c Jcosmos) ListDatabase(resp databaseResponse) error {
	return c.cosmosRequest("/dbs", "", http.MethodPost, "", nil, resp)
}
func (c Jcosmos) DeleteDatabase(id string, resp databaseResponse) error {
	if len(id) > databsaseIDMaxLength {
		return errors.New(ErrorDatabaseIDTooLong)
	}
	return c.cosmosRequest("/dbs", "", http.MethodPost, "", nil, resp)
}
