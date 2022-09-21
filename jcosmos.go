package jcosmos

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	jcosmosVersion                  = "0.0.1"
	timeoutSeconds                  = 9
	tokenVersion             string = "1.0"
	userAgent                string = "Jcosmos/" + jcosmosVersion
	cosmosDbApiVersionString string = "2015-12-16" //"2018-12-31"
)

func LocalInit(host, key, db, collection string, metrics, crossPartition bool, logger *log.Logger) Jcosmos {
	return Jcosmos{
		url:                  fmt.Sprintf("https://%s/", host),
		loglevel:             LogLevelWarn,
		keytype:              "master",
		key:                  key,
		db:                   db,
		host:                 host,
		coll:                 collection,
		populatequerymetrics: metrics,
		enablecrosspartition: crossPartition,
		logger:               logger,
	}
}
func EasyInit(host, key, db, collection string) Jcosmos {
	return Jcosmos{
		url:                  fmt.Sprintf("https://%s.documents.azure.com:443/", host),
		loglevel:             LogLevelError,
		keytype:              "master",
		key:                  key,
		db:                   db,
		host:                 host,
		coll:                 collection,
		populatequerymetrics: false,
		enablecrosspartition: true,
		logger:               log.New(os.Stderr, "", log.LstdFlags), // default logger
	}
}

func Init(host, keytype, key, db, collection string, loglevel loglevel, metrics, crossPartition bool, logger *log.Logger) Jcosmos {
	return Jcosmos{
		url:                  fmt.Sprintf("https://%s.documents.azure.com:443/", host),
		loglevel:             loglevel,
		keytype:              keytype,
		key:                  key,
		db:                   db,
		host:                 host,
		coll:                 collection,
		populatequerymetrics: metrics,
		enablecrosspartition: crossPartition,
		logger:               logger,
	}
}

type Jcosmos struct {
	url                  string
	loglevel             loglevel
	keytype              string
	key                  string
	db                   string
	host                 string
	coll                 string
	populatequerymetrics bool
	enablecrosspartition bool
	logger               *log.Logger
}

func (c Jcosmos) UseDB(db string) Jcosmos {
	c.db = db
	return c
}
func (c Jcosmos) UseCol(coll string) Jcosmos {
	c.coll = coll
	return c
}
func (c Jcosmos) UseLogLevel(loglevel loglevel) Jcosmos {
	c.loglevel = loglevel
	return c
}
func (c Jcosmos) cosmosRequest(rl, pk, method string, body []byte, headers map[string]string, obj interface{}) (string, error) {
	c.logReq(rl, pk, method, body, headers)
	client := &http.Client{Timeout: timeoutSeconds * time.Second}
	req, err := http.NewRequest(strings.ToUpper(method), c.url+rl, strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	c.generateHeaders(req, body, pk, rl, headers)
	if err != nil {
		return "", err
	}
	byts, _ := httputil.DumpRequest(req, true)
	c.log(LogLevelTrace, string(byts))
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	return resp.Header.Get("x-ms-continuation"), c.processResponse(resp, obj)
}
func (c Jcosmos) processResponse(r *http.Response, obj interface{}) error {
	if r.StatusCode >= 400 && r.StatusCode < 500 {
		c.log(LogLevelError, bodyToStr(r.Body))
		return errors.New("client error ")
	}
	if r.StatusCode >= 500 {
		c.log(LogLevelError, bodyToStr(r.Body))
		return errors.New("server error")
	}
	if r.StatusCode == http.StatusNoContent {
		c.log(LogLevelTrace, "No Content")
		return nil
	}
	if r.StatusCode == http.StatusCreated || r.StatusCode == http.StatusOK {

		err := json.Unmarshal([]byte(bodyToStr(r.Body)), obj)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("UNKNOWN ERROR")
}
func (c Jcosmos) generateHeaders(r *http.Request, body []byte, pk, resourceLink string, headers map[string]string) {
	t := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	if _, ok := headers["Content-Type"]; !ok {
		r.Header.Add("Content-Type", "application/json")
	}

	r.Header.Add("x-ms-activity-id", uuid.New().String())
	if pk != "" {
		r.Header.Add("x-ms-documentdb-partitionkey", "[\""+pk+"\"]")
	}
	r.Header.Add("Authorization", c.generateAuthHeader(r.Method, t, resourceLink))
	r.Header.Add("Content-Length", strconv.Itoa(len(body)))
	r.Header.Add("User-Agent", userAgent)
	r.Header.Add("X-Ms-Date", t)
	r.Header.Add("X-Ms-Documentdb-Populatequerymetrics", strconv.FormatBool(c.populatequerymetrics))
	r.Header.Add("X-Ms-Version", cosmosDbApiVersionString)

	for k, v := range headers {
		r.Header.Add(k, v)
	}
	for k, v := range headers {
		c.log(LogLevelDebug, fmt.Sprintf("%s: %s", k, v))
	}
}

func (c Jcosmos) generateAuthHeader(m, t, rl string) string {
	authParts := []string{
		strings.ToLower(m),
		strings.ToLower(getResourceType(rl)),
		strings.TrimSuffix(rl, "/docs"),
		strings.ToLower(t),
		"",
	}
	authString := strings.Join(authParts, "\n") + "\n"
	c.log(LogLevelTrace, authString)
	b64 := base64.StdEncoding
	byteKey, _ := b64.DecodeString(c.key)

	mac := hmac.New(sha256.New, []byte(byteKey))
	mac.Write([]byte(authString))
	sig := b64.EncodeToString(mac.Sum(nil))
	return url.QueryEscape(fmt.Sprintf("type=%s&ver=%s&sig=%s", c.keytype, tokenVersion, sig))
}

func getResourceType(resourceLink string) string {
	rLink := strings.ToLower(resourceLink)
	switch {
	case strings.Contains(rLink, "docs"):
		return "docs"
	case strings.Contains(rLink, "colls"):
		return "colls"
	case strings.Contains(rLink, "dbs"):
		return "dbs"
	case strings.Contains(rLink, "users"):
		return "users"
	default:
		return ""
	}
}
