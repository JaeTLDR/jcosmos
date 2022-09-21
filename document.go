package jcosmos

import (
	"net/http"
	"strconv"
)

func (c Jcosmos) CreateDocument(pk string, body []byte, upsert bool, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	_, err := c.cosmosRequest(rl, pk, http.MethodPost, body, map[string]string{"x-ms-documentdb-is-upsert": strconv.FormatBool(upsert)}, obj)
	return err
}

func (c Jcosmos) ReadDocument(id, pk string, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	_, err := c.cosmosRequest(rl, pk, http.MethodGet, emptyByteArr, nil, obj)
	return err
}

func (c Jcosmos) UpdateDocument(id, pk string, body []byte, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	_, err := c.cosmosRequest(rl, pk, http.MethodPut, body, nil, obj)
	return err
}

func (c Jcosmos) DeleteDocument(id, pk string) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	_, err := c.cosmosRequest(rl, pk, http.MethodDelete, emptyByteArr, nil, nil)
	return err
}

func (c Jcosmos) XPartitionQueryDocument(body []byte, cont string, obj interface{}) (string, error) {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	h := map[string]string{
		"x-ms-documentdb-isquery":                    "true",
		"Content-Type":                               "application/query+json",
		"x-ms-max-item-count":                        "50",
		"x-ms-documentdb-query-enablecrosspartition": "true",
	}
	if len(cont) > 0 {
		h["x-ms-continuation"] = cont
	}
	resp, err := c.cosmosRequest(rl, "", http.MethodPost, body, h, obj)
	return resp.Header.Get("x-ms-continuation"), err
}

func (c Jcosmos) QueryDocument(pk string, body []byte, cont string, obj interface{}) (string, error) {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	h := map[string]string{
		"x-ms-documentdb-isquery": "true",
		"Content-Type":            "application/query+json",
		"x-ms-max-item-count":     "50",
		// "x-ms-consistency-level":"",
	}
	if len(cont) > 0 {
		h["x-ms-continuation"] = cont
	}
	resp, err := c.cosmosRequest(rl, pk, http.MethodPost, body, h, obj)
	return resp.Header.Get("x-ms-continuation"), err
}

func (c Jcosmos) ListDocument(pk, cont string, obj interface{}) (string, error) {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	h := map[string]string{
		"x-ms-max-item-count":                        "50",
		"x-ms-documentdb-query-enablecrosspartition": strconv.FormatBool(c.enablecrosspartition),
		// "x-ms-consistency-level":"",
	}
	if len(cont) > 0 {
		h["x-ms-continuation"] = cont
	}
	resp, err := c.cosmosRequest(rl, pk, http.MethodGet, emptyByteArr, h, obj)
	return resp.Header.Get("x-ms-continuation"), err
}
