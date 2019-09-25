package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	apierrors "github.com/Bonial-International-GmbH/site24x7-go/api/errors"
	log "github.com/sirupsen/logrus"
)

const (
	BasePath = "https://www.site24x7.com/api"
)

type Request struct {
	client     HTTPClient
	resource   string
	resourceID string
	header     http.Header
	verb       string
	body       []byte
	respBody   []byte
	resp       *http.Response
	err        error
}

func NewRequest(client HTTPClient) *Request {
	return &Request{
		client: client,
	}
}

func (r *Request) Resource(resource string) *Request {
	r.resource = resource
	return r
}

func (r *Request) ResourceID(resourceID string) *Request {
	r.resourceID = resourceID
	return r
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

func (r *Request) Get() *Request {
	return r.Verb("GET")
}

func (r *Request) Post() *Request {
	return r.Verb("POST")
}

func (r *Request) Put() *Request {
	return r.Verb("PUT")
}

func (r *Request) Delete() *Request {
	return r.Verb("DELETE")
}

func (r *Request) AddHeader(key, value string) *Request {
	if r.header == nil {
		r.header = http.Header{}
	}

	r.header.Add(key, value)
	return r
}

func (r *Request) Body(v interface{}) *Request {
	r.body, r.err = json.Marshal(v)
	return r
}

func (r *Request) buildRawURL() string {
	rawURL := BasePath + "/" + r.resource
	if r.resourceID != "" {
		rawURL += "/" + r.resourceID
	}

	return rawURL
}

func (r *Request) Do() *Request {
	if r.err != nil {
		return r
	}

	req, err := r.buildRequest()
	if err != nil {
		r.err = err
		return r
	}

	r.respBody, r.err = r.doRequest(req)

	return r
}

func (r *Request) buildRequest() (*http.Request, error) {
	url, err := url.Parse(r.buildRawURL())
	if err != nil {
		return nil, err
	}

	req := &http.Request{
		Method: r.verb,
		Header: r.header,
		Body:   ioutil.NopCloser(bytes.NewReader(r.body)),
		URL:    url,
	}

	return req, nil
}

func (r *Request) doRequest(req *http.Request) ([]byte, error) {
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Debugf("received: %s", string(respBody))

	if resp.StatusCode > 0 && resp.StatusCode < 400 {
		return respBody, nil
	}

	err = parseErrorResponse(resp.StatusCode, respBody)

	return respBody, err
}

func (r *Request) Into(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	resp := &response{}

	r.err = json.Unmarshal(r.respBody, resp)
	if r.err != nil {
		return r.err
	}

	return json.Unmarshal([]byte(resp.Data), v)
}

func (r *Request) Err() error {
	return r.err
}

func parseErrorResponse(statusCode int, body []byte) error {
	resp := &errorResponse{}

	err := json.Unmarshal(body, resp)
	if err != nil {
		return err
	}

	return apierrors.NewStatusError(statusCode, resp.ErrorCode, resp.Message, resp.ErrorInfo)
}
