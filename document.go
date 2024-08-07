package jcosmos

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (c Jcosmos) CreateDocument(pk []string, upsert bool, obj any) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	_, err = c.cosmosRequest(rl, pk, http.MethodPost, body, map[string]string{"x-ms-documentdb-is-upsert": strconv.FormatBool(upsert)}, obj)
	return err
}

func (c Jcosmos) ReadDocument(id string, pk []string, obj any) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	_, err := c.cosmosRequest(rl, pk, http.MethodGet, emptyByteArr, nil, obj)
	return err
}

func (c Jcosmos) UpdateDocument(id string, pk []string, obj any) error {
	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	_, err = c.cosmosRequest(rl, pk, http.MethodPut, body, nil, obj)
	return err
}

func (c Jcosmos) DeleteDocument(id string, pk []string) error {
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs/" + id
	_, err := c.cosmosRequest(rl, pk, http.MethodDelete, emptyByteArr, nil, nil)
	return err
}

func (c Jcosmos) XPartitionQueryDocument(query Query, cont string, obj any) (string, error) {
	body, err := query.ToJson()
	if err != nil {
		return "", err
	}
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
	resp, err := c.cosmosRequest(rl, nil, http.MethodPost, body, h, obj)
	return resp.Header.Get("x-ms-continuation"), err
}

func (c Jcosmos) QueryDocument(pk []string, query Query, cont string, obj any) (string, error) {
	body, err := query.ToJson()
	if err != nil {
		return "", err
	}
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

func (c Jcosmos) ListDocument(pk []string, cont string, obj any) (string, error) {
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

func (c Jcosmos) PatchDocument(id string, pk []string, p Patch, obj any) error {
	var patchOpErr error
	for _, po := range p.Operations {
		patchOpErr = po.validate()
		if patchOpErr != nil {
			break
		}
	}
	if patchOpErr != nil {
		return patchOpErr
	}
	rl := "dbs/" + c.db + "/colls/" + c.coll + "/docs"
	_, err := c.cosmosRequest(rl, pk, http.MethodPatch, emptyByteArr, nil, obj)
	return err
}
