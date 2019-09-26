package rest

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	apierrors "github.com/Bonial-International-GmbH/site24x7-go/api/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequest_buildRequest(t *testing.T) {
	c := &fakeHTTPClient{}

	body := map[string]string{
		"foo": "bar",
	}

	r := NewRequest(c, "DELETE").
		Resource("foos").
		ResourceID("123").
		Body(body).
		AddHeader("Foo", "Bar")

	req, err := r.buildRequest()

	require.NoError(t, err)

	buf, err := ioutil.ReadAll(req.Body)

	require.NoError(t, err)

	assert.Equal(t, "Bar", req.Header.Get("Foo"))
	assert.Equal(t, "DELETE", req.Method)
	assert.Equal(t, "/api/foos/123", req.URL.Path)
	assert.Equal(t, `{"foo":"bar"}`, string(buf))
}

func TestRequest_DoInto(t *testing.T) {
	c := newFakeHTTPClient().
		WithStatusCode(200).
		WithResponseBody([]byte(`{"data":{"foo":"bar"}}`))

	r := NewRequest(c, "POST")

	var result map[string]string

	err := r.
		Resource("foos").
		ResourceID("123").
		Do().
		Into(&result)

	require.NoError(t, err)

	expected := map[string]string{
		"foo": "bar",
	}

	assert.Equal(t, expected, result)
}

func TestRequest_DoIntoError(t *testing.T) {
	c := newFakeHTTPClient().
		WithStatusCode(404).
		WithResponseBody([]byte(`{"error_code":456,"message":"not found","error_info":{"foo":"bar"}}`))

	r := NewRequest(c, "PUT")

	var result map[string]string

	err := r.
		Resource("foos").
		ResourceID("123").
		Body(nil).
		Do().
		Into(&result)

	require.Error(t, err)
	assert.Nil(t, result)

	statusErr, ok := err.(apierrors.ExtendedStatusError)
	if !ok {
		t.Fatalf("expected ExtendedStatusError, got %T", err)
	}

	assert.Equal(t, 404, statusErr.StatusCode())
	assert.Equal(t, 456, statusErr.ErrorCode())
	assert.Equal(t, "not found", statusErr.Error())
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, statusErr.ErrorInfo())
}

type fakeHTTPClient struct {
	resp       *http.Response
	err        error
	called     int
	calledWith []*http.Request
}

func newFakeHTTPClient() *fakeHTTPClient {
	return &fakeHTTPClient{}
}

func (c *fakeHTTPClient) WithStatusCode(code int) *fakeHTTPClient {
	if c.resp == nil {
		c.resp = &http.Response{}
	}

	c.resp.StatusCode = code
	return c
}

func (c *fakeHTTPClient) WithResponseBody(buf []byte) *fakeHTTPClient {
	if c.resp == nil {
		c.resp = &http.Response{}
	}

	c.resp.Body = ioutil.NopCloser(bytes.NewReader(buf))
	c.resp.ContentLength = int64(len(buf))
	return c
}

func (c *fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	c.called++
	if c.calledWith == nil {
		c.calledWith = make([]*http.Request, 0, 1)
	}
	c.calledWith = append(c.calledWith, req)
	return c.resp, c.err
}
