package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Bonial-International-GmbH/site24x7-go/api"
	apierrors "github.com/Bonial-International-GmbH/site24x7-go/api/errors"
	log "github.com/sirupsen/logrus"
)

// Request is a wrapper for preparing and sending a *http.Request. It provides
// funtionality for encoding arbitrary types to the wire format and back.
type Request struct {
	client     HTTPClient
	baseURL    string
	resource   string
	resourceID string
	header     http.Header
	verb       string
	body       []byte
	err        error
}

// NewRequest creates a new *Request which uses client to send out the prepared
// *http.Request.
func NewRequest(client HTTPClient, verb, baseURL string) *Request {
	r := &Request{
		client:  client,
		baseURL: baseURL,
		verb:    verb,
	}

	r.AddHeader("Accept", "application/json; version=2.0")

	return r
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

// AddHeader adds an HTTP header to the request.
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
	rawURL := r.baseURL + "/" + r.resource
	if r.resourceID != "" {
		rawURL += "/" + r.resourceID
	}

	return rawURL
}

// Do sends the request. This is a no-op if there were errors while building
// the request.
func (r *Request) Do() Response {
	if r.err != nil {
		return Response{err: r.err}
	}

	req, err := r.buildRequest()
	if err != nil {
		return Response{err: err}
	}

	return r.doRequest(req)
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

func (r *Request) doRequest(req *http.Request) Response {
	resp, err := r.client.Do(req)
	if err != nil {
		return Response{err: err}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{err: err}
	}

	log.Debugf("received: %s", string(body))

	if resp.StatusCode > 0 && resp.StatusCode < 400 {
		return Response{body: body}
	}

	return Response{err: createStatusError(resp.StatusCode, body)}
}

func createStatusError(statusCode int, body []byte) error {
	resp := &api.ErrorResponse{}

	err := json.Unmarshal(body, resp)
	if err != nil {
		log.Errorf("received bad error response body: %q, error was: %s", string(body), err)
		return apierrors.NewStatusError(statusCode, fmt.Sprintf("server replied with: %s", string(body)))
	}

	return apierrors.NewExtendedStatusError(statusCode, resp.Message, resp.ErrorCode, resp.ErrorInfo)
}
