package rest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	apierrors "github.com/Bonial-International-GmbH/site24x7-go/api/errors"
	log "github.com/sirupsen/logrus"
)

const (
	// BasePath is the root url of the Site24x7 API.
	BasePath = "https://www.site24x7.com/api"
)

// Request is a wrapper for preparing and sending a *http.Request. It provides
// funtionality for encoding arbitrary types to the wire format and back.
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

// NewRequest creates a new *Request which uses client to send out the prepared
// *http.Request.
func NewRequest(client HTTPClient, verb string) *Request {
	return &Request{
		client: client,
		verb:   verb,
	}
}

// Resource sets the API resource which the request should be built for, e.g.
// 'monitors'. The resulting API resource path for this would be
// '/api/monitors'.
func (r *Request) Resource(resource string) *Request {
	r.resource = resource
	return r
}

// ResourceID sets the API resource ID which the request should be built for,
// e.g. '123'. Example: if the resource was set to 'monitors', the resulting
// API resource path will be '/api/monitors/123'.
func (r *Request) ResourceID(resourceID string) *Request {
	r.resourceID = resourceID
	return r
}

// AddHeader adds a HTTP header to the request.
func (r *Request) AddHeader(key, value string) *Request {
	if r.header == nil {
		r.header = http.Header{}
	}

	r.header.Add(key, value)
	return r
}

// Body marshals v into the request body.
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

// Do sends the request. This is a no-op if there were errors while building
// the request.
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

	err = createStatusError(resp.StatusCode, respBody)

	return respBody, err
}

// Into unmarshals the response body into v. The passed in value must be a
// pointer. It returns any error that occured during the request. This is a
// no-op if there were errors before.
func (r *Request) Into(v interface{}) error {
	if r.err != nil {
		return r.err
	}

	resp := &api.Response{}

	r.err = json.Unmarshal(r.respBody, resp)
	if r.err != nil {
		return r.err
	}

	return json.Unmarshal([]byte(resp.Data), v)
}

// Err returns the request error if there was one.
func (r *Request) Err() error {
	return r.err
}

func createStatusError(statusCode int, body []byte) error {
	resp := &api.ErrorResponse{}

	err := json.Unmarshal(body, resp)
	if err != nil {
		return err
	}

	return apierrors.NewExtendedStatusError(statusCode, resp.Message, resp.ErrorCode, resp.ErrorInfo)
}
