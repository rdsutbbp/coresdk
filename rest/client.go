package rest

import (
	"net/http"
	"time"
)

type Interface interface {
	Verb(verb string) *Request
	Post() *Request
	Get() *Request
}

type Opt func(client *RESTClient) error

type RESTClient struct {
	protocol string
	addr     string
	port     string

	retryTimes int
	retryDelay time.Duration

	delegationUUID string
	lifeCycleUUID  string

	auth *XForwardedAuthUser

	headers http.Header

	// Set specific behavior of the client.  If not set http.DefaultClient will be used.
	client *http.Client
}

func (r *RESTClient) Verb(verb string) *Request {
	return NewRequest(r).Verb(verb)
}

func (r *RESTClient) Post() *Request {
	return r.Verb("POST")
}

func (r *RESTClient) Get() *Request {
	return r.Verb("GET")
}

func RESTClientFor(config *RESTClient) (*RESTClient, error) {
	httpClient := &http.Client{}

	rest := &RESTClient{
		protocol:       config.protocol,
		addr:           config.addr,
		port:           config.port,
		retryTimes:     config.retryTimes,
		retryDelay:     config.retryDelay,
		delegationUUID: config.delegationUUID,
		lifeCycleUUID:  config.lifeCycleUUID,
		auth:           config.auth,
		headers:        config.headers,
		client:         httpClient,
	}
	return rest, nil
}
