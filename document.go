package jcosmos

import (
	"net/http"
	"strconv"
)

func (c Jcosmos) CreateDocument(pk, body string, upsert bool, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	return c.cosmosRequest(rl, pk, http.MethodPost, body, map[string]string{"x-ms-documentdb-is-upsert": strconv.FormatBool(upsert)}, obj)
}

func (c Jcosmos) ReadDocument(id, pk string, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	return c.cosmosRequest(rl, pk, http.MethodGet, "", nil, obj)
}

func (c Jcosmos) UpdateDocument(id, pk, body string, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	return c.cosmosRequest(rl, pk, http.MethodPut, body, nil, obj)
}

func (c Jcosmos) DeleteDocument(id, pk string) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	return c.cosmosRequest(rl, pk, http.MethodDelete, "", nil, nil)
}

func (c Jcosmos) XPartitionQueryDocument(body, cont string, obj interface{}) error {
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
	return c.cosmosRequest(rl, "", http.MethodPost, body, h, obj)
}

func (c Jcosmos) QueryDocument(pk, body, cont string, obj interface{}) error {
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
	return c.cosmosRequest(rl, pk, http.MethodPost, body, h, obj)
}

func (c Jcosmos) ListDocument(pk, cont string, obj interface{}) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	h := map[string]string{
		"x-ms-max-item-count":                        "50",
		"x-ms-documentdb-query-enablecrosspartition": strconv.FormatBool(c.enablecrosspartition),
		// "x-ms-consistency-level":"",
	}
	if len(cont) > 0 {
		h["x-ms-continuation"] = cont
	}
	return c.cosmosRequest(rl, pk, http.MethodGet, "", h, obj)
}
