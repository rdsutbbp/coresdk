package rest

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// Request allows for building up a request to a server in a chained fashion.
// Any errors are stored until the end of your call, so you only have to
// check once.
type Request struct {
	c *RESTClient

	timeout    time.Duration
	maxRetries int

	verb    string
	headers http.Header
	params  url.Values

	// output
	err  error
	body io.Reader
}

func NewRequest(c *RESTClient) *Request {
	var timeout time.Duration
	if c.Client != nil {
		timeout = c.Client.Timeout
	}
	r := &Request{
		c:          c,
		timeout:    timeout,
		maxRetries: 10,
	}
	return r
}

// Timeout makes the request use the given duration as an overall timeout for the
// request. Additionally, if set passes the value as "timeout" parameter in URL.
func (r *Request) Timeout(d time.Duration) *Request {
	if r.err != nil {
		return r
	}
	r.timeout = d
	return r
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}

func (r *Request) MaxRetries(maxRetries int) *Request {
	if maxRetries < 0 {
		maxRetries = 0
	}
	r.maxRetries = maxRetries
	return r
}

// Param creates a query parameter with the given string value.
func (r *Request) Param(paramName, s string) *Request {
	if r.err != nil {
		return r
	}
	return r.setParam(paramName, s)
}

func (r *Request) setParam(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

// Body makes the request use obj as the body. Optional.
// If obj is a string, try to read a file of that name.
// If obj is a []byte, send it directly.
// If obj is an io.Reader, use it directly.
// If obj is a runtime.Object, marshal it correctly, and set Content-Type header.
// If obj is a runtime.Object and nil, do nothing.
func (r *Request) Body(obj interface{}) {}

// DoRaw executes the request but does not process the response body.
func (r *Request) DoRaw(ctx context.Context) ([]byte, error) {
	return nil, nil
}

// Result contains the result of calling Request.Do().
type Result struct {
	body        []byte
	contentType string
	err         error
	statusCode  int
}

// Do format and executes the request. Returns a Result object for easy response
// processing.
//
// Error type:
//  * http.Client.Do errors are returned directly.
func (r *Request) Do() Result {
	fmt.Println(r.c.addr)
	return Result{}
}

// Raw returns the raw result.
func (r Result) Raw() ([]byte, error) {
	return r.body, r.err
}

// Into stores the result into obj, if possible. If obj is nil it is ignored.
func (r Result) Into() (interface{}, error) {
	return nil, nil
}

// StatusCode returns the HTTP status code of the request. (Only valid if no
// error was returned.)
func (r Result) StatusCode(statusCode *int) Result {
	*statusCode = r.statusCode
	return r
}

// Error returns the error executing the request, nil if no error occurred.
func (r Result) Error() error {
	return nil
}
