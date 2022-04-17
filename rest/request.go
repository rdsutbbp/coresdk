package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rdsutbbp/utilx/httpx"
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

func (r *Request) C() *RESTClient {
	return r.c
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
func (r *Request) url() (string, error) {
	if r.c.auth != nil {
		newLifecycleUUID, err := r.StoreAuth2LifeCycleUUID(r.c.auth)
		if err != nil {
			return "", err
		}
		r.c.lifeCycleUUID = newLifecycleUUID
	}
	if r.params == "" && r.c.delegationUUID != "" && r.c.lifeCycleUUID != "" {
		r.params = fmt.Sprintf("?DelegationUUID=%s&LifeCycleUUID=%s", r.c.delegationUUID, r.c.lifeCycleUUID)
	}
	return fmt.Sprintf("%s://%s:%s", r.c.protocol, r.c.addr, r.c.port+r.subPath+r.params), nil
}

func (r *Request) StoreAuth2LifeCycleUUID(user *XForwardedAuthUser) (string, error) {
	// generate a new lifeCycleUUID DELEGATION-EVENT-XXX
	lifeCycleUUID := "DELEGATION-EVENT-" + uuid.New().String()
	// store auth to lifeCycleUUID
	url := fmt.Sprintf("%s://%s:%s/gateway/delegation/api/v1/auth/store?LifeCycleUUID=%s", r.c.protocol, r.c.addr, r.c.port, lifeCycleUUID)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := httpx.DoPost(ctx, user, url, r.c.headers, r.c.retryTimes, r.c.retryDelay)
	if err != nil {
		return "", err
	}
	return lifeCycleUUID, nil
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
	url, err := r.url()
	if err != nil {
		return Result{err: err}
	}
	request, err := http.NewRequestWithContext(ctx, r.verb, url, r.body)
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

	defer rawResp.Body.Close()

	if rawResp == nil {
		return Result{err: err}
	}

	if rawResp.StatusCode != 200 {
		return Result{err: errors.Errorf("unhealthy status code [%d]", rawResp.StatusCode)}
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

	// parse response data
	// code message data
	j, err := simplejson.NewJson(r.body)
	if err != nil {
		return err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return err
	}
	if code != http.StatusOK {
		message, _ := j.Get("message").String()
		return fmt.Errorf(message)
	}
	marshalJSON, err := j.Get("data").MarshalJSON()
	if err != nil {
		return err
	}

	return json.Unmarshal(marshalJSON, obj)
}

// StatusCode returns the HTTP status code of the request. (Only valid if no
// error was returned.)
func (r Result) StatusCode() int {
	return r.statusCode
}

func (r Result) Body() []byte {
	return r.body
}

// Error returns the error executing the request, nil if no error occurred.
func (r Result) Error() error {
	return r.err
}
