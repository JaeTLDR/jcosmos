package jcosmos

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	maxThroughput = 5000
)

func (c Jcosmos) ReadOffer(rid string, obj OfferResponse) error {
	_, err := c.cosmosRequest("/dbs/offers/"+rid, "", http.MethodGet, emptyByteArr, nil, obj)
	return err
}

func (c Jcosmos) ListOffer(q Query, obj ListOfferResponse) error {
	_, err := c.cosmosRequest("/dbs/offers", "", http.MethodPost, emptyByteArr, nil, obj)
	return err
}
func (c Jcosmos) QueryOffer(query Query, obj ListOfferResponse) error {
	body, err := json.Marshal(query)
	if err != nil {
		return err
	}
	_, err = c.cosmosRequest("/dbs/offers", "", http.MethodPost, body, nil, obj)
	return err
}
func (c Jcosmos) ReplaceOffer(offer OfferRequest, obj OfferResponse) error {
	if offer.OfferVersion != "v2" {
		return errors.New("offerVersion invalid. jcomos only supports version 2")
	}
	if offer.Content.OfferThroughput < 400 || offer.Content.OfferThroughput > maxThroughput {
		return errors.New("invalid offerThrroughput value")
	}
	offer.Rid = offer.ID
	body, err := json.Marshal(offer)
	if err != nil {
		return err
	}

	_, err = c.cosmosRequest("/dbs/offers", "", http.MethodPost, body, nil, obj)
	return err
}
