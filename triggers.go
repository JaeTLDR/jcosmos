package jcosmos

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	maxTriggerIDLen = 255

	TriggerTypePre  = "Pre"
	TriggerTypePost = "Post"

	TriggerOperationAll     = "All"
	TriggerOperationCreate  = "Create"
	TriggerOperationReplace = "Replace"
	TriggerOperationDelete  = "Delete"
)

var (
	allowedTriggerOps   map[string]struct{} // [4]string{"All", "Create", "Replace", "Delete"}
	allowedTriggerTypes map[string]struct{} // [2]string{"Pre", "Post"}
)

func init() {
	allowedTriggerOps = make(map[string]struct{})   // [4]string{"All", "Create", "Replace", "Delete"}
	allowedTriggerTypes = make(map[string]struct{}) // [2]string{"Pre", "Post"}
	allowedTriggerOps["All"] = struct{}{}
	allowedTriggerOps["Create"] = struct{}{}
	allowedTriggerOps["Replace"] = struct{}{}
	allowedTriggerOps["Delete"] = struct{}{}
	allowedTriggerTypes["Pre"] = struct{}{}
	allowedTriggerTypes["Post"] = struct{}{}
}

func (c Jcosmos) CreateTrigger(t Trigger, obj TriggerResponse) error {
	if len(t.ID) > maxTriggerIDLen {
		return errors.New("id must be less than 255 characters")
	}
	if _, ok := allowedTriggerOps[t.TriggerOperation]; !ok {
		return errors.New("triggerOperation must be one of All,Create,Replace,Delete")
	}
	if _, ok := allowedTriggerTypes[t.TriggerType]; !ok {
		return errors.New("triggerOperation must be one of Pre,Post")
	}
	body, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = c.cosmosRequest("dbs/"+c.db+"/colls/"+c.coll+"/triggers", "", http.MethodPost, body, nil, obj)
	return err
}

func (c Jcosmos) ListTrigger(obj TriggerList) error {
	_, err := c.cosmosRequest("dbs/"+c.db+"/colls/"+c.coll+"/triggers", "", http.MethodGet, emptyByteArr, nil, obj)
	return err
}

func (c Jcosmos) ReplaceTrigger(t Trigger, obj TriggerResponse) error {
	if len(t.ID) > maxTriggerIDLen {
		return errors.New("id must be less than 255 characters")
	}
	if _, ok := allowedTriggerOps[t.TriggerOperation]; !ok {
		return errors.New("triggerOperation must be one of All,Create,Replace,Delete")
	}
	if _, ok := allowedTriggerTypes[t.TriggerType]; !ok {
		return errors.New("triggerOperation must be one of Pre,Post")
	}
	body, err := json.Marshal(t)
	if err != nil {
		return err
	}
	_, err = c.cosmosRequest("dbs/"+c.db+"/colls/"+c.coll+"/triggers/"+t.ID, "", http.MethodPut, body, nil, obj)
	return err
}

func (c Jcosmos) DeleteTrigger(id string) error {
	_, err := c.cosmosRequest("dbs/"+c.db+"/colls/"+c.coll+"/triggers/"+id, "", http.MethodDelete, emptyByteArr, nil, nil)
	return err
}
