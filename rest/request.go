package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
)

// Request allows for building up a request to a server in a chained fashion.
// Any errors are stored until the end of your call, so you only have to
// check once.
type Request struct {
	c *RESTClient

	verb string

	subPath string

	params string

	// output
	err  error
	body io.Reader
}

func NewRequest(c *RESTClient) *Request {

	r := &Request{
		c: c,
	}
	return r
}

func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

// SubPath set subPath
// e.g. /api/v1/credential/init
func (r *Request) SubPath(subPath string) *Request {
	r.subPath = subPath
	return r
}

func (r *Request) Params(params string) *Request {
	r.params = params
	return r
}

// url get url for request
func (r *Request) url() string {
	if r.params == "" {
		r.params = fmt.Sprintf("?DelegationUUID=%s&LifeCycleUUID=%s", r.c.delegationUUID, r.c.lifeCycleUUID)
	}
	return fmt.Sprintf("%s://%s:%s", r.c.protocol, r.c.addr, r.c.port+r.subPath+r.params)
}

// Body makes the request use obj as the body. Optional.
// If obj is a string, try to read a file of that name.
// If obj is a []byte, send it directly.
// default marshal it
func (r *Request) Body(obj interface{}) *Request {
	if r.err != nil {
		return r
	}

	switch t := obj.(type) {
	case string:
		data, err := ioutil.ReadFile(t)
		if err != nil {
			r.err = err
			return r
		}
		r.body = bytes.NewReader(data)
	case []byte:
		r.body = bytes.NewReader(t)
	default:
		data, err := json.Marshal(obj)
		if err != nil {
			r.err = err
			return r
		}
		r.body = bytes.NewReader(data)
	}
	return r
}

// Result contains the result of calling Request.Do().
type Result struct {
	body       []byte
	err        error
	statusCode int
}

// Do format and executes the request. Returns a Result object for easy response
// processing.
//
// Error type:
//  * http.Client.Do errors are returned directly.
func (r *Request) Do(ctx context.Context) Result {
	request, err := http.NewRequestWithContext(ctx, "POST", r.url(), r.body)
	if err != nil {
		return Result{err: err}
	}

	request.Header = r.c.headers

	client := http.DefaultClient

	var rawResp *http.Response
	// if meet error, retry times that you set
	for k := 0; k < r.c.retryTimes; k++ {
		rawResp, err = doRequest(client, request)
		if err != nil {
			// sleep retry delay
			time.Sleep(r.c.retryDelay)
			continue
		}
		break
	}

	if rawResp == nil {
		return Result{err: err}
	}

	data, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return Result{err: err, statusCode: rawResp.StatusCode}
	}
	return Result{
		body:       data,
		err:        nil,
		statusCode: rawResp.StatusCode,
	}
}

func doRequest(client *http.Client, request *http.Request) (*http.Response, error) {
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, errors.New("response is nil")
	}
	defer res.Body.Close()
	return res, nil
}

// Into stores the result into obj, if possible. If obj is nil it is ignored.
func (r Result) Into(obj interface{}) error {
	if r.err != nil {
		return r.err
	}
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		return errors.New("object is not a ptr")
	}
	return json.Unmarshal(r.body, obj)
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
